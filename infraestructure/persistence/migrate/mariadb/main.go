package main

import (
	"github.com/fabricioereche/mariadb_tst/infraestructure/persistence"
	_ "github.com/go-sql-driver/mysql"
	migrate "github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	db, err := persistence.GetDrive()

	if err != nil {
		println(err.Error())
		return
	}

	defer db.Close()
	driver, err := mysql.WithInstance(db, &mysql.Config{})

	if err != nil {
		println(err.Error())
		return
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"mysql",
		driver,
	)

	if err != nil {
		println(err.Error())
		return
	}

	m.Up()
}
