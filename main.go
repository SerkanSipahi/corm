package corm

import (
	"context"
	"github.com/flimzy/kivik"
)

type Config struct {
	// default "couch"
	Host string
	// default "http://localhost:5984/"
	DriverName string
	DBName     string
}

type ClientConfig struct {
	// default "couch"
	Host string
	// default "http://localhost:5984/"
	DriverName string
}

var defaultDriverName = "couch"
var defaultHostName = "http://localhost:5984/"

// New returns a db instance connection by passed DBName.
// When it fails it will return an error. Optionally "Host"
// and DriverName can be passed with Config struct.
//
// 	ctx := context.TODO()
//
// 	db, err := New(ctx, Config{
// 		DBName: "honeyglass",
// 	})
func New(ctx context.Context, config Config) (*Orm, error) {

	initDefaults(&config)

	client, err := NewClient(ctx, ClientConfig{
		Host:       config.Host,
		DriverName: config.DriverName,
	})

	db, err := client.DB(ctx, config.DBName)
	if err != nil {
		return &Orm{Db: db}, err
	}

	return NewOrm(db), err
}

func NewClient(ctx context.Context, config ClientConfig) (*Client, error) {

	client, err := kivik.New(ctx, config.DriverName, config.Host)
	if err != nil {
		return &Client{}, err
	}

	return client, err
}

func initDefaults(config *Config) {

	if config.Host == "" {
		config.Host = defaultHostName
	}
	if config.DriverName == "" {
		config.DriverName = defaultDriverName
	}
}
