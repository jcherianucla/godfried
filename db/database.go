package database

import (
	"database/sql"
	"fmt"
	utils "github.com/jcherianucla/godfried/app/helpers"
	_ "github.com/lib/pq"
	"os"
)

const (
	DB_USER     = "godfried-root"
	DB_NAME     = "godfried"
	DB_PASSWORD = os.Getenv("GODFRIED_DB_PASS")
	DB_HOST     = "localhost"
)

type DB struct {
	store *sql.DB
	err   error
}

func (db *DB) StartDB() {
	DBInfo = os.Getenv("DB_INFO")
	if DBInfo == "" {
		DBInfo = fmt.Sprintf("user=%s password=%s dbname=%s host=%s sslmode=disable",
			DB_USER, DB_PASSWORD, DB_NAME, DB_HOST)
	}
	db.store, db.err = sql.Open("postgres", DBInfo)
	utils.CheckError(db.err)
}

func (db *DB) CreateTable(table string, record interface{}) {
	selString := fmt.Sprintf("SELECT * FROM %s", table)
	_, err = db.store.Exec(selString)
	if err != nil {
		inString := utils.Builder.BuildInsertion(record)
		_, err = db.store.Exec(inString)
		utils.CheckError(err)
	}
}
