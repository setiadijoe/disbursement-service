package migration

import (
	"database/sql"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/file"
)

var db *sql.DB

// RunMigrations ...
func RunMigrations(connectionString string) error {
	var err error

	db, err = sql.Open("postgres", connectionString)
	if err != nil {
		return fmt.Errorf("error when open postgres connection: %v", err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("error when creating postgres instance: %v", err)
	}

	var m *migrate.Migrate

	fsrc, err := (&file.File{}).Open("file://migration/files/")
	if err != nil {
		return fmt.Errorf("error when open file: %v", err)
	}

	m, err = migrate.NewWithInstance(
		"file",
		fsrc,
		"postgres",
		driver,
	)

	if err != nil {
		return fmt.Errorf("error when creating database instance: %v", err)
	}

	err = m.Up()

	if err != nil && err != migrate.ErrNoChange {
		err := m.Down()
		if err != nil {
			fmt.Println(err)
			return fmt.Errorf("error when migrate down: %v", err)
		}
	}
	return nil
}
