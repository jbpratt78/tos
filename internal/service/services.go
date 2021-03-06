// Package service ...
package service

import (
	"database/sql"
	"log"
)

// Services is the implementation of both the
// Menu and Order service for interacting with the database
// as well as the databse itself
type Services struct {
	Order OrderService
	Menu  MenuService
	db    *sql.DB
}

// Config is used for determing use of which services and db
type Config func(*Services) error

// NewServices creates a Service struct with all of the
// ServiceConfigs passed into it
func NewServices(cfgs ...Config) (*Services, error) {
	var s Services
	for _, cfg := range cfgs {
		if err := cfg(&s); err != nil {
			return nil, err
		}
	}
	return &s, nil
}

// WithSqlite takes in a path and opens the database
func WithSqlite(path string) Config {
	return func(s *Services) error {
		db, err := sql.Open("sqlite3", path)
		if err != nil {
			log.Fatalf("Failed to open DB: %v\n", err)
		}

		s.db = db
		return nil
	}
}

// WithMenu is used for calling NewMenuService with a specific db
func WithMenu() Config {
	return func(s *Services) error {
		menu, err := NewMenuService(s.db)
		if err != nil {
			return err
		}
		s.Menu = menu
		return nil
	}
}

// WithOrder is used for calling NewOrderService with a specific db
func WithOrder() Config {
	return func(s *Services) error {
		order, err := NewOrderService(s.db)
		if err != nil {
			return err
		}
		s.Order = order
		return nil
	}
}

// Close closes the current database
func (s *Services) Close() error {
	return s.db.Close()
}
