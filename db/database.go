// docker run --rm   --name pg-docker -e POSTGRES_PASSWORD=dev -d -p 5432:5432 -v $HOME/docker/volumes/postgres:/var/lib/postgresql/data  postgres

package db

import (
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func CreateDatabase() (*sqlx.DB, error) {

	connStr := "user=dev password=dev sslmode=disable"

	db, err := sqlx.Open("postgres", connStr)

	_, err = db.Exec("CREATE SCHEMA IF NOT EXISTS aegis")

	if err != nil {
		fmt.Printf(" CREATE SCHEMA failed with error: %s\n", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
		return nil, err
	}

	if err := migrateDatabase(db); err != nil {
		return db, err
	}

	return db, nil
}

func migrateDatabase(db *sqlx.DB) error {

	if err := db.Ping(); err != nil {
		log.Fatalf("could not ping DB... %v", err)
	}

	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	if err != nil {
		return err
	}

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	migration, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s/db/migrations", dir),
		"postgres",
		driver,
	)

	if err != nil {
		fmt.Println(err)
		return err
	}

	migration.Log = &MigrationLogger{}

	migration.Log.Printf("applying database migrations")

	// err = migration.Force(1)
	// err = migration.Down()

	if err != nil && err != migrate.ErrNoChange {
		fmt.Println(err)
		return err
	}
	err = migration.Up()

	if err != nil && err != migrate.ErrNoChange {
		fmt.Println(err)
		return err
	}

	version, _, err := migration.Version()
	if err != nil {
		return err
	}

	migration.Log.Printf("Active database version: %d", version)

	return nil
}
