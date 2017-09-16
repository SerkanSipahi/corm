package corm_test

import (
	"context"
	"fmt"
	"github.com/serkansipahi/corm"
	"log"
)

// How to init an custom orm explained step by step
func ExampleNewOrm() {

	// build config
	config := corm.Config{
		Host:       "http://localhost:5984/",
		DriverName: "couch",
	}

	// create client instance
	ctx := context.TODO()
	client, err := corm.NewClient(ctx, corm.ClientConfig{
		Host:       config.Host,
		DriverName: config.DriverName,
	})
	if err != nil {
		log.Fatal(err)
	}

	// create db
	db, err := client.DB(ctx, config.DBName)
	// create orm model
	model := corm.NewOrm(db)
	if err != nil {
		log.Fatal(err)
	}

	// define any struct
	type Person struct {
		Name     string
		Surename string
		Age      int
	}

	// save person
	id, rev, err := model.Save(context.TODO(), Person{
		Name:     "Serkan",
		Surename: "Sipahi",
	})

	fmt.Println(id, rev, err) // log 2233434323, 1-33434565, nil
}

func ExampleMethodenName_inklammer2() {
	fmt.Println("Hello")
}
