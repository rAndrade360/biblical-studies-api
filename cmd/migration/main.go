package main

import (
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/rAndrade360/biblical-studies-api/internal/infra/database/sqlite"
)

func main() {

	log.Println(os.Getenv("MIGRATIONS_PATH"))

	if len(os.Args) <= 1 {
		log.Fatal("Error: Please, specify the action")
	}

	db, err := sqlite.New()
	if err != nil {
		log.Fatalf("Err to create db: %s", err.Error())
	}

	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{
		DatabaseName: "biblical-studies",
	})

	if err != nil {
		log.Fatalf("Err to create driver: %s", err.Error())
	}

	m, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file:///%s", os.Getenv("MIGRATIONS_PATH")),
		"biblical-studies", driver)

	if err != nil {
		log.Fatalf("Err to create migrate: %s", err.Error())
	}

	if os.Args[1] == "up" {
		err = m.Up() // or m.Step(2) if you want to explicitly set the number of migrations to run
		if err != nil {
			log.Fatalf("Err to migrate: %s", err.Error())
		}
	} else if os.Args[1] == "down" {
		err = m.Down() // or m.Step(2) if you want to explicitly set the number of migrations to run
		if err != nil {
			log.Fatalf("Err to migrate: %s", err.Error())
		}
	}
}
