package main

import (
	"database/sql"
)

// Store interface can be easily injected into our services
type Store interface {
	// User
	CreateUser() error
}

type Storage struct {
	db *sql.DB
}

// create a constructor as well for the struct above
func NewStore(db *sql.DB) *Storage {
	return &Storage{
		db: db,
	}
}

// create a method for the storage so that it implements the interface
func (s *Storage) CreateUser() error {
	return nil
}
