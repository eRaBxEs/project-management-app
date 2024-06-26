package main

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

type MySQLStorage struct {
	db *sql.DB
}

func NewMySQLStorage(cfg mysql.Config) *MySQLStorage {
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	// Ping to be sure connection succeeded
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to MySQL")
	return &MySQLStorage{db: db}
}

// This needs to be created so as to initialize tables in the database
func (s *MySQLStorage) Init() (*sql.DB, error) {

	// initialize the tables
	return s.db, nil
}
