package main

import (
	"flag"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/k-yomo/go_echo_api_boilerplate/config"
	"github.com/pkg/errors"
	"os"
	"time"
)

var migrationFilePath = "file://./migrations/"

func main() {
	flag.Parse()
	if flag.Arg(0) == "" {
		showUsage()
		os.Exit(1)
	}

	m := newMigrate()
	version, dirty, _ := m.Version()
	force := flag.Bool("f", false, "force execute fixed sql")
	if dirty && *force {
		fmt.Println("force=true: force execute current version sql")
		m.Force(int(version))
	}

	switch flag.Arg(0) {
	case "new":
		newMigration(flag.Arg(1))
	case "up":
		up(m)
	case "down":
		down(m)
	case "drop":
		drop(m)
	case "version":
		showVersionInfo(m.Version())
	default:
		fmt.Println("\nerror: invalid command '", flag.Arg(0), "'")
		showUsage()
		os.Exit(0)
	}
}

func newMigrate() *migrate.Migrate {
	var dbConfig *config.DBConfig
	var err error
	if os.Getenv("ENV") == "test" {
		dbConfig, err = config.NewTestDBConfig()
	} else {
		dbConfig, err = config.NewDBConfig()
	}
	fmt.Printf("Target database: %s\n", dbConfig.DBName)
	if err != nil {
		fmt.Println(errors.Wrap(err, "initialize DB config failed"))
	}
	m, err := migrate.New(migrationFilePath, config.NewDsn(dbConfig))
	if err != nil {
		fmt.Println(errors.Wrap(err, "initialize Migrate instance failed"))
		os.Exit(1)
	}
	return m
}

func showUsage() {
	fmt.Println(`
-------------------------------------
Usage:
  go run migrate.go <command>

Commands:
  new NAME	Create new up & down migration files
  up		Apply up migrations
  down		Apply down migrations
  drop		Drop everything
  version	Check current migrate version
-------------------------------------`)
}

func newMigration(name string) {
	if name == "" {
		fmt.Println("\nerror: migration file name must be supplied as an argument")
		os.Exit(1)
	}
	base := fmt.Sprintf("./migrations/%s_%s", time.Now().Format("20060102030405"), name)
	ext := ".sql"
	createFile(base + ".up" + ext)
	createFile(base + ".down" + ext)
}

func createFile(fname string) {
	if _, err := os.Create(fname); err != nil {
		panic(err)
	}
}

func up(m *migrate.Migrate) {
	fmt.Println("Before:")
	showVersionInfo(m.Version())
	err := m.Up()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("\nUpdated:")
		version, dirty, err := m.Version()
		showVersionInfo(version, dirty, err)
	}
}

func down(m *migrate.Migrate) {
	fmt.Println("Before:")
	showVersionInfo(m.Version())
	err := m.Steps(-1)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("\nUpdated:")
		showVersionInfo(m.Version())
	}
}

func drop(m *migrate.Migrate) {
	err := m.Drop()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Dropped all migrations")
		return
	}
}

func showVersionInfo(version uint, dirty bool, err error) {
	fmt.Println("-------------------")
	fmt.Println("version : ", version)
	fmt.Println("dirty   : ", dirty)
	fmt.Println("error   : ", err)
	fmt.Println("-------------------")
}
