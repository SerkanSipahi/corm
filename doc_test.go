package corm_test

import (
	"context"
	"fmt"
	"github.com/serkansipahi/corm"
	"log"
)

// You dont need this step when using corm. But if want to know how to init an custom
// orm, please follow the example step by step.
func ExampleNewOrm_orm() {

	// build config
	config := corm.Config{
		Host:       "http://localhost:5984/",
		DriverName: "couch",
	}

	// create client instance
	client, err := corm.NewClient(context.TODO(), corm.ClientConfig{
		Host:       config.Host,
		DriverName: config.DriverName,
	})
	if err != nil {
		log.Fatal(err)
	}

	// create db
	db, err := client.DB(context.TODO(), config.DBName)
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

	fmt.Println(id, rev, err)
	// Output: 889c9653a6b490cc24c85d78b10076c7, 1-68a533f5dc76a65b56b7329b9d4086ab, nil
}

// NewClient creates a new client instance that is very useful when you want to use the
// client api of kivik.Client see https://godoc.org/github.com/flimzy/kivik#Client .
// Here is an example for Authentication an user
func ExampleNewClient() {

	// create client instance
	client, err := corm.NewClient(context.TODO(), corm.ClientConfig{
		Host:       "http://localhost:5984/",
		DriverName: "couch",
	})
	if err != nil {
		log.Fatal(err)
	}

	type Credentials struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	}

	err = client.Authenticate(context.TODO(), Credentials{
		Name:     "myname",
		Password: "somepassword",
	})

	fmt.Println(err)
	// Output: nil

}

func ExampleOrm_Save_save1() {

	db, err := corm.New(context.TODO(), corm.Config{
		DBName: "mydbname",
	})

	if err != nil {
		log.Fatal(err)
	}

	docId, rev, err := db.Save(context.TODO(), Product{
		Name: "Foo",
	})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(docId, rev, nil)
	// Output: 889c9653a6b490cc24c85d78b10076c7, 1-68a533f5dc76a65b56b7329b9d4086ab, nil
}

func ExampleOrm_Save_save2() {

	db, err := corm.New(context.TODO(), corm.Config{
		DBName: "mydbname",
	})

	// create document with predefined id
	docId, rev, err := db.Save(context.TODO(), Product{
		Id:   "123456",
		Name: "Foo",
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(docId, rev)
	// Output: 123456, 1-68a533f5dc76a65b56b7329b9d4086ab, nil
}
