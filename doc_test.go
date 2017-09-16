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

// NewClient creates a new client instance that is very useful when you want to use the
// client api of kivik.Client see https://godoc.org/github.com/flimzy/kivik#Client .
// Here is an example for Authentication an user
func ExampleNewClient_client() {

	// create client instance
	ctx := context.TODO()
	client, err := corm.NewClient(ctx, corm.ClientConfig{
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

	fmt.Println(err) // nil when success and not nil when it fails

}

func ExampleOrm_Save_save() {

	// create document with auto generated id
	ctx := context.TODO()
	db, err := corm.New(ctx, corm.Config{
		// as you can see "Id" is not defined
		DBName: "mydbname",
	})

	if err != nil {
		log.Fatal(err)
	}

	docId, rev, err := db.Save(ctx, Product{
		Name: "Foo",
	})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(docId, rev) // 232434234234234, 1-s323sf34243

	// create document with predefined id
	docId, rev, err = db.Save(ctx, Product{
		Id:   "123456",
		Name: "Foo",
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(docId, rev) // 123456, 1-asdfw34sdf3
}
