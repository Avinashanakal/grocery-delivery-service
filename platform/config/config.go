package config

import (
	"fmt"
	"os"
)

type DBConfig struct {
	Host       string
	Port       string
	Database   string
	Username   string
	Password   string
	SSLMode    string
	EnableLogs bool
}

// Connection string
func (d DBConfig) URIBuilder() string {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=5432",
		d.Host, d.Username, d.Password, d.Database,
	)
	return dsn
}

type config struct {
	ENV string
}

type Config interface {
	Database() (DBConfig, error)
}

// Initialise the config
func New() (Config, error) {
	return &config{
		ENV: os.Getenv("ENV"),
	}, nil
}

// Add more environments
func (c *config) Database() (DBConfig, error) {
	switch c.ENV {
	case "dev":
		return DBConfig{
			Host:       "localhost",
			Port:       "5432",
			Database:   "postgres",
			Username:   "postgres",
			Password:   "postgres",
			SSLMode:    "disable",
			EnableLogs: true,
		}, nil

	}
	panic("ENV not supported")
}
