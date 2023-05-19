package psql

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	snowflakelib "github.com/bwmarrin/snowflake"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/opencorelabs/fira/internal/config"
	"github.com/opencorelabs/fira/internal/persistence/snowflake"
	"github.com/stretchr/testify/suite"
	"strings"
)

type TestHelper struct {
	db            DBCloser
	migrationsDir string
	pgUrl         string
	schemaName    string
	gen           snowflake.Generator
	s             *suite.Suite
}

func NewTestHelper(s *suite.Suite) *TestHelper {
	cfg, cfgErr := config.Init()
	s.Require().NoError(cfgErr, "error initializing config")

	h := &TestHelper{s: s, migrationsDir: cfg.MigrationsDir}

	testSchemaName := h.randSchemaName()
	pgUrl := cfg.PostgresUrl
	if strings.Contains(pgUrl, "?") {
		pgUrl = pgUrl + "&search_path=" + testSchemaName
	} else {
		pgUrl = pgUrl + "?search_path=" + testSchemaName
	}

	h.schemaName = testSchemaName
	h.pgUrl = pgUrl

	poolCfg, poolCfgErr := pgxpool.ParseConfig(pgUrl)
	s.Require().NoError(poolCfgErr, "error parsing pgx pool config")

	var newErr error
	h.db, newErr = pgxpool.NewWithConfig(context.Background(), poolCfg)
	s.Require().NoError(newErr, "error initializing pgx pool")

	h.gen, newErr = snowflakelib.NewNode(cfg.NodeId)
	s.Require().NoError(newErr, "error initializing snowflake generator")

	return h
}

func (h *TestHelper) Migrate() {
	_, schemaErr := h.db.Exec(context.Background(), fmt.Sprintf(`DROP SCHEMA IF EXISTS %s; CREATE SCHEMA %s`, h.schemaName, h.schemaName))
	h.s.Require().NoErrorf(schemaErr, "error creating schema %s", h.schemaName)

	migrateErr := Migrate(h.migrationsDir, h.pgUrl)
	h.s.Require().NoErrorf(migrateErr, "error migrating database")
}

func (h *TestHelper) Close() {
	_, schemaErr := h.db.Exec(context.Background(), fmt.Sprintf("DROP SCHEMA IF EXISTS %s CASCADE", h.schemaName))
	h.s.Require().NoErrorf(schemaErr, "error dropping schema %s", h.schemaName)
	h.db.Close()
}

func (h *TestHelper) Reset() {
	cmds := []string{
		"TRUNCATE TABLE accounts",
	}
	for _, cmd := range cmds {
		_, err := h.db.Exec(context.Background(), cmd)
		h.s.Require().NoErrorf(err, "error executing command %s", cmd)
	}
}

func (h *TestHelper) DB() DBCloser {
	return h.db
}

func (h *TestHelper) Generator() snowflake.Generator {
	return h.gen
}

func (h *TestHelper) randSchemaName() string {
	buff := make([]byte, 16)
	_, readErr := rand.Read(buff)
	h.s.Require().NoErrorf(readErr, "error generating random schema name")
	return fmt.Sprintf("test_%s", hex.EncodeToString(buff))
}
