package database

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/chrislim1914/xrp-transaction/api/models"
	"github.com/chrislim1914/xrp-transaction/bootstrap"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func Migration(config *bootstrap.Config) error {
	dsn := connectionStr(config)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		fmt.Println(err)
		return err
	}
	row, err := db.Query("select * from schema_migrations")
	if err != nil {
		// this means that this is a freash migration
		driver, err := postgres.WithInstance(db, &postgres.Config{})
		m, err := migrate.NewWithDatabaseInstance(
			"file:./database/migration",
			"postgres", driver)
		if err != nil {
			log.Fatal(err)
		}
		if err := m.Up(); err != nil {
			log.Fatal(err)
		}
		return nil
	}
	defer row.Close()
	var schema models.SchemaMigrations
	for row.Next() {
		if err := row.Scan(&schema.Version, &schema.Dirty); err != nil {
			return err
		}
	}

	// check migration file
	lastfile := getRecentSchemaNumber()
	if lastfile <= schema.Version {
		return nil
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file:./database/migration",
		"postgres", driver)
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil {
		log.Fatal(err)
	}
	return nil
}

func getRecentSchemaNumber() int {
	fname := "./database/migration"
	files, err := ioutil.ReadDir(fname)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	var modTime time.Time
	var names []string
	for _, fi := range files {
		if fi.Mode().IsRegular() {
			if !fi.ModTime().Before(modTime) {
				if fi.ModTime().After(modTime) {
					modTime = fi.ModTime()
					names = names[:0]
				}
				names = append(names, fi.Name())
			}
		}
	}
	if len(names) > 0 {
		schemasplit := strings.Split(names[0], "_")
		number, _ := strconv.Atoi(schemasplit[0])
		return number
	}
	return 0
}
