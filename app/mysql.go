package app

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	// database technology name
	database = "mysql"

	// database connection string properties
	dbUser            = "root"
	dbPass            = "windows!"
	dbConnectProtocol = "tcp"
	dbHost            = "127.0.0.1"
	dbPort            = "3306"
	dbName            = "school"

	TblClass = "class"
)

type (
	// IDb is interface of project's database
	IDb interface {
		Query(query string, args ...interface{}) (IRows, error)
	}

	IRows interface{}

	db struct {
		db *sql.DB
	}

	rows struct {
		rows *sql.Rows
	}
)

func initMySQL() (IDb, error) {
	dbSession, err := sql.Open(database, fmt.Sprintf("%s:%s@%s(%s:%s)/%s",
		dbUser, dbPass, dbConnectProtocol, dbHost, dbPort, dbName))
	if err != nil {
		return nil, err
	}

	return &db{db: dbSession}, nil
}

func (d *db) Query(query string, args ...interface{}) (IRows, error) {
	retRows, err := d.db.Query(query, args...)
	return &rows{rows: retRows}, err
}
