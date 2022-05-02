package persistence

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var (
	dbname = os.Getenv("DB_NAME")
	dbhost = os.Getenv("DB_HOST")
	dbuser = os.Getenv("DB_USER")
	dbpass = os.Getenv("DB_PASS")
	dbport = os.Getenv("DB_PORT")
)

func GetDrive() (*sql.DB, error) {

	// Create the database handle, confirm driver is present
	strconn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbuser, dbpass, dbhost, dbport, dbname)
	db, err := sql.Open("mysql", strconn)

	if err != nil {
		fmt.Println("could not connect to mariaDB: ", err.Error())
		return nil, err
	}

	return db, nil

}
