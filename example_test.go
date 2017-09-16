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

	// type Product struct {
	// 	Id          string `json:"_id,omitempty"`  // required in this style
	// 	Rev         string `json:"_rev,omitempty"` // required in this style
	// 	Type        string `json:"type"`           // required in this style
	// 	Name        string `json:"vendorId"`
	// }

	// create client
	client, err := corm.NewClient(context.TODO(), corm.ClientConfig{
		Host:       "http://localhost:5984/",
		DriverName: "couch",
	})
	if err != nil {
		log.Fatal(err)
	}

	// create db
	db, err := client.DB(context.TODO(), "mydbname")
	// create orm
	orm := corm.NewOrm(db)
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
	id, rev, err := orm.Save(context.TODO(), Person{
		Name:     "Serkan",
		Surename: "Sipahi",
	})

	fmt.Println(id, rev, err)
	// Output: 889c9653a6b490cc24c85d78b10076c7, 1-68a533f5dc76a65b56b7329b9d4086ab, nil
}

// Here is an example for Authentication an user
func ExampleNewClient() {

	// create client
	client, err := corm.NewClient(context.TODO(), corm.ClientConfig{
		Host:       "http://localhost:5984/",
		DriverName: "couch",
	})

	// log when it fails
	if err != nil {
		log.Fatal(err)
	}

	// define credentials
	type Credentials struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	}

	// authenticate user
	err = client.Authenticate(context.TODO(), Credentials{
		Name:     "myname",
		Password: "somepassword",
	})

	fmt.Println(err)
	// Output: nil

}

// Save a document with "auto generated" id by couchDB
func ExampleOrm_Save_save1() {

	// type Product struct {
	// 	Id          string `json:"_id,omitempty"`  // required in this style
	// 	Rev         string `json:"_rev,omitempty"` // required in this style
	// 	Type        string `json:"type"`           // required: tag this but don´t touch it
	// 	Name        string `json:"vendorId"`
	// }

	// create orm
	orm, err := corm.New(context.TODO(), corm.Config{
		DBName: "mydbname",
	})

	if err != nil {
		log.Fatal(err)
	}

	// save document
	docId, rev, err := orm.Save(context.TODO(), Product{
		Name: "Foo",
	})

	// log when it fails
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(docId, rev, nil)
	// Output: 889c9653a6b490cc24c85d78b10076c7, 1-68a533f5dc76a65b56b7329b9d4086ab, nil
}

// Save a document with "predefined" id
func ExampleOrm_Save_save2() {

	// type Product struct {
	// 	Id          string `json:"_id,omitempty"`  // required in this style
	// 	Rev         string `json:"_rev,omitempty"` // required in this style
	// 	Type        string `json:"type"`           // required: tag this but don´t touch it
	// 	Name        string `json:"name"`
	// }

	// create orm
	orm, err := corm.New(context.TODO(), corm.Config{
		DBName: "mydbname",
	})

	// create document with predefined id
	docId, rev, err := orm.Save(context.TODO(), Product{
		Id:   "123456",
		Name: "Foo",
	})

	// log when it fails
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(docId, rev, err)
	// Output: 123456, 1-68a533f5dc76a65b56b7329b9d4086ab, nil
}

// Update an document by given Id and Rev
func ExampleOrm_Update() {

	// type Product struct {
	// 	Id          string `json:"_id,omitempty"`  // required in this style
	// 	Rev         string `json:"_rev,omitempty"` // required in this style
	// 	Type        string `json:"type"`           // required: tag this but don´t touch it
	// 	Name        string `json:"name"`
	// }

	// create orm
	orm, err := corm.New(context.TODO(), corm.Config{
		DBName: "mydbname",
	})

	// log when it fails
	if err != nil {
		log.Fatal(err)
	}

	// update document
	rev, err := orm.Update(context.TODO(), Product{
		Id:   "889c9653a6b490cc24c85d78b10076c7",
		Rev:  "1-68a533f5dc76a65b56b7329b9d4086ab",
		Name: "Bar",
	})

	// log when it fails
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(rev, err)
	// Output: 2-68a533f5dc76a65b56b7329b9d4086ab, nil
}

// Read document
func ExampleOrm_Read() {

	// type Product struct {
	// 	Id          string `json:"_id,omitempty"`  // required in this style
	// 	Rev         string `json:"_rev,omitempty"` // required in this style
	// 	Type        string `json:"type"`           // required: tag this but don´t touch it
	// 	Name        string `json:"name"`
	// }

	// create orm
	orm, err := corm.New(context.TODO(), corm.Config{
		DBName: "mydbname",
	})

	// log when it fails
	if err != nil {
		log.Fatal(err)
	}

	var product Product
	_, err = orm.Read(context.TODO(), "889c9653a6b490cc24c85d78b10076c7", &product)

	fmt.Println(product)
	// Output: Product{Name: "Bar", Id: "889c9653a6b490cc24c85d78b10076c7", Rev: "2-68a533f5dc76a65b56b7329b9d4086ab"}
}

// Read document with options
func ExampleOrm_Read2() {

	// example coming soon
}

// Delete document
func ExampleOrm_Delete() {

	// create orm
	orm, err := corm.New(context.TODO(), corm.Config{
		DBName: "mydbname",
	})

	// log when it fails
	if err != nil {
		log.Fatal(err)
	}

	// delete doc
	_, err = orm.Delete(context.TODO(),
		"889c9653a6b490cc24c85d78b10076c7",
		"2-68a533f5dc76a65b56b7329b9d4086ab",
	)

	// log when it fails
	if err != nil {
		log.Fatal(err)
	}
}
