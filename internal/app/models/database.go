package models

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jackc/pgx/stdlib"
	"log"

	"gopkg.in/ini.v1"
)

type database struct {
	LocalDB *sql.DB
}

var db *database

func InitDatabase(cfg *ini.File) *database {
	internalDB := cfg.Section("database").Key("internal_db").String()

	localdb, err := NewDBConnection("pgx", internalDB)
	if err != nil {
		log.Fatal(err)
	}

	db = &database{
		LocalDB: localdb,
	}

	return db
}

func (d *database) Close() {
	d.LocalDB.Close()
}

func NewDBConnection(driver string, databaseURL string) (*sql.DB, error) {
	db, err := sql.Open(driver, databaseURL)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, err
}
