package main

import (
	"database/sql"
	"errors"
	"flag"
	"fmt"
	"log"
	"odonto/internal/infra/config"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

func errControl() {
	if r := recover(); r != nil {
		e := r.(error)
		log.Printf("Application shutdown caused by error: '%s'\n", e.Error())
	}
}

func cmdName(order int) string {
	switch order {
	case 0:
		return "up"
	case 1:
		return "down"
	case 2:
		return "reverse"
	case 3:
		return "force"
	default:
		return "unknown"
	}
}

func main() {
	// capture panic error
	defer errControl()
	// vars
	var (
		migrateCmd     int
		migrateVersion int
		yamlFile       string
	)
	// flags
	flag.StringVar(&yamlFile, "c", "configs/context/app-dev.yaml", "Load application settings from path file name")
	flag.IntVar(&migrateCmd, "cmd", 0, "Comands: up=0, down=1, reverse=2, force=3")
	flag.IntVar(&migrateVersion, "force", 1, "For force version. Default 1")
	flag.Parse()
	//load configurations (panic if yamlFile not exists)
	log.Printf("Load configurations in %s", yamlFile)
	config.Init(yamlFile)
	//load configuration
	log.Printf("Loading configuration for postgresql")
	// configure database
	log.Printf("Connection database")
	db, err := sql.Open("sqlite3", "database/foo.db")
	if err != nil {
		panic(err)
	}
	// check connection
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	log.Printf("Database ping: OK")
	// configure pg drive
	driver, err := sqlite.WithInstance(db, &sqlite.Config{})

	if err != nil {
		panic(err)
	}

	// set directory for migrations
	path, _ := os.Getwd()
	m, err := migrate.NewWithDatabaseInstance(fmt.Sprintf("file:%s/migrations", path), "postgres", driver)
	if err != nil {
		panic(err)
	}
	//
	v, d, err := m.Version()
	if err != nil && !errors.Is(err, migrate.ErrNilVersion) {
		panic(err)
	}
	log.Printf("Current version: %d", v)
	log.Printf("Dirty: %t", d)
	if migrateCmd == 9 {
		return
	}
	// migration of commands
	log.Printf("Executing command: %s", cmdName(migrateCmd))
	err = nil
	if migrateCmd == 0 {
		log.Printf("Up to all")
		err = m.Up()
	} else if migrateCmd == 1 {
		log.Printf("Down to all")
		err = m.Down()
	} else if migrateCmd == 2 {
		log.Printf("Reversing to %d", v-1)
		err = m.Steps(-1)
	} else if migrateCmd == 3 {
		log.Printf("Forcing version: %d", migrateVersion)
		err = m.Force(migrateVersion)
	} else {
		log.Printf("Unknown command")
	}
	// fail
	if err != nil {
		panic(err)
	}
}
