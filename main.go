package corm

import (
	"context"
	"github.com/flimzy/kivik"
)

// Config can be passed in corm.New(...)
type Config struct {
	Host       string // default "couch"
	DriverName string // default "http://localhost:5984/"
	DBName     string
}

type ClientConfig struct {
	Host       string // default "couch"
	DriverName string // default "http://localhost:5984/"
}

var defaultDriverName = "couch"
var defaultHostName = "http://localhost:5984/"

// New returns a db instance connection by passed DBName.
// When it fails it will return an error. Optionally "Host"
// and DriverName can be passed with Config struct.
//
// 	ctx := context.TODO()
// 	db, err := corm.New(ctx, Config{
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
		return nil, err
	}

	return NewOrm(db), err
}

// NewClient returns a client instance that can be used
// for e.g. CouchDb´ s build in authentication. Here are the full method list
// of the client instance: https://godoc.org/github.com/flimzy/kivik#Client
func NewClient(ctx context.Context, config ClientConfig) (*Client, error) {

	client, err := kivik.New(ctx, config.DriverName, config.Host)
	if err != nil {
		return nil, err
	}

	return client, err
}

// InitDefaults sets the default values when
// Conifg struct does not contain the "Host" and "DriverName" member
func initDefaults(config *Config) {

	if config.Host == "" {
		config.Host = defaultHostName
	}
	if config.DriverName == "" {
		config.DriverName = defaultDriverName
	}
}
