package psql

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"os"
	"path/filepath"
)

func Migrate(migrationsDir, psqlUrl string) error {
	fmt.Println("Migrating database...")
	fmt.Println("Migrations dir:", migrationsDir)
	if !filepath.IsAbs(migrationsDir) {
		var pathErr error
		migrationsDir, pathErr = filepath.Abs(migrationsDir)
		if pathErr != nil {
			return fmt.Errorf("unable to get absolute path for migrations dir: %w", pathErr)
		}
		fmt.Println("Absolute migrations dir:", migrationsDir)
	}

	source := fmt.Sprintf("file://%s", migrationsDir)

	fmt.Println("Migrations source: ", source)

	fmt.Println("Migrations dsn: ", psqlUrl)

	m, err := migrate.New(source, psqlUrl)
	if err != nil {
		return fmt.Errorf("unable to create migrator: %w", err)
	}

	m.Log = &mgLog{verbose: true}

	merr := m.Up()
	if merr != nil && merr != migrate.ErrNoChange {
		return fmt.Errorf("unable to migrate: %w", merr)
	}

	return nil
}

// mgLog is a dummy logger for migrate
type mgLog struct {
	verbose bool
}

func (l *mgLog) Printf(format string, v ...interface{}) {
	_, _ = fmt.Fprintf(os.Stderr, format, v...)
}

func (l *mgLog) Println(args ...interface{}) {
	_, _ = fmt.Fprintln(os.Stderr, args...)
}

func (l *mgLog) Verbose() bool {
	return l.verbose
}

func (l *mgLog) fatal(args ...interface{}) {
	l.Println(args...)
	os.Exit(1)
}

func (l *mgLog) fatalErr(err error) {
	l.fatal("error:", err)
}
