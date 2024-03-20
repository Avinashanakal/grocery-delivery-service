package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type db interface {
	URIBuilder() string
}

// Connection create database connection
func New(config db) (*gorm.DB, error) {
	uri := config.URIBuilder()
	db, err := gorm.Open(postgres.Open(uri), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
