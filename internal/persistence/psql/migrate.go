package psql

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/opencorelabs/fira/internal/logging"
	"go.uber.org/zap"
	"path/filepath"
)

func Migrate(logger logging.Provider, migrationsDir, psqlUrl string) error {
	log := logger.Logger().Named("pg-migrator")
	log.Info("migrating database...")
	absMigrationsDir := migrationsDir
	if !filepath.IsAbs(migrationsDir) {
		var pathErr error
		absMigrationsDir, pathErr = filepath.Abs(migrationsDir)
		if pathErr != nil {
			return fmt.Errorf("unable to get absolute path for migrations dir: %w", pathErr)
		}
	}

	source := fmt.Sprintf("file://%s", migrationsDir)

	log.Info("migrations config",
		zap.String("dir", migrationsDir),
		zap.String("absDir", absMigrationsDir),
		zap.String("source", source),
	)

	m, err := migrate.New(source, psqlUrl)
	if err != nil {
		return fmt.Errorf("unable to create migrator: %w", err)
	}
	if m != nil {
		defer func() {
			_, closeErr := m.Close()
			if closeErr != nil {
				log.Warn("migrator closed with error:", zap.Error(closeErr))
			}
		}()
	}

	m.Log = &mgLog{verbose: true, log: log}

	merr := m.Up()
	if merr != nil && merr != migrate.ErrNoChange {
		return fmt.Errorf("unable to migrate: %w", merr)
	} else if merr == migrate.ErrNoChange {
		log.Info("database already up to date")
	} else {
		log.Info("migration complete")
	}

	return nil
}

// mgLog is a dummy logger for migrate
type mgLog struct {
	log     *zap.Logger
	verbose bool
}

func (l *mgLog) Printf(format string, v ...interface{}) {
	l.log.Info(fmt.Sprintf(format, v...))
}

func (l *mgLog) Println(args ...interface{}) {
	l.log.Info(fmt.Sprint(args...))
}

func (l *mgLog) Verbose() bool {
	return l.verbose
}

func (l *mgLog) fatal(args ...interface{}) {
	l.log.Fatal(fmt.Sprint(args...))
}

func (l *mgLog) fatalErr(err error) {
	l.fatal("error:", err)
}
