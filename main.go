package corm

import (
	"context"
	"github.com/flimzy/kivik"
)

type Config struct {
	Host       string
	DriverName string
	DBName     string
}

var defaultDriverName = "couch"
var defaultHostName = "http://localhost:5984/"

func New(ctx context.Context, config Config) (*Orm, error) {

	initDefaults(&config)

	client, err := kivik.New(ctx, config.DriverName, config.Host)
	if err != nil {
		return &Orm{}, err
	}

	db, err := client.DB(ctx, config.DBName)
	if err != nil {
		return &Orm{}, err
	}

	return NewOrm(db), err
}

func initDefaults(config *Config) {
	if config.Host == "" {
		config.Host = defaultHostName
	}
	if config.DriverName == "" {
		config.DriverName = defaultDriverName
	}
}
