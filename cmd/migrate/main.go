package main

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide a migration direction: 'up' or 'down'")
	}
	direction := os.Args[1]
	connstr := "host=localhost port=5432 user=postgres password=1234 dbname=rest-api sslmode=disable"
	db, err := sql.Open("postgres", connstr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	instance, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}
	fSrc, err := (&file.File{}).Open("migrations")
	if err != nil {
		log.Fatal(err)
	}
	m, err := migrate.NewWithInstance("file", fSrc, "postgres", instance)
	if err != nil {
		log.Fatal(err)
	}

	switch direction {
	case "up":
		if err := m.Up(); err != nil {
			if err == migrate.ErrNoChange {
				fmt.Println("No new migrations to apply.")
			} else {
				log.Fatal(err)
			}

		} else {
			fmt.Println("Migrations applied successfully!")
		}
	case "down":
		if err := m.Down(); err != nil {
			if err == migrate.ErrNoChange {
				log.Println("No migrations to rollback.")
			} else {
				log.Fatal(err)
			}
		} else {
			log.Println("Migrations rolled back successfully!")
		}
	default:
		log.Fatal("Invalid direction. Use 'up' or 'down'.")
	}
}
