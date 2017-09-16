package corm_test

import (
	"context"
	"fmt"
	"github.com/serkansipahi/corm"
	"log"
	"testing"
)

type Product struct {
	Id          string `json:"_id,omitempty"`
	Rev         string `json:"_rev,omitempty"`
	VendorId    int    `json:"vendorId"`
	VendorType  string `json:"vendorType"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CategoryId  string `json:"categoryId"`
	UrlId       string `json:"urlId"`
}

func TestCorn_Foo(t *testing.T) {

	ctx := context.TODO()

	db, err := corm.New(ctx, corm.Config{
		DBName: "honeyglass",
	})

	if err != nil {
		log.Fatal(err)
	}

	docId, rev, err := db.Save(ctx, Product{
		Id:   "111-222-333",
		Name: "Liya",
	})

	if err != nil {
		log.Fatal(err)
	}

	// create doc
	docId, rev, err = db.Save(ctx, Product{
		Name: "Serkan",
	})
	if err != nil {
		log.Fatal(err)
	}

	// update doc
	rev, err = db.Update(ctx, Product{
		Id:   docId,
		Rev:  rev,
		Name: "Xing",
	})
	if err != nil {
		log.Fatal(err)
	}

	// read doc with optional options
	options := map[string]interface{}{"a": "b"}
	var product Product
	_, err = db.Read(ctx, docId, &product, options)

	// delete doc
	rev, err = db.Delete(ctx, docId, rev)

	fmt.Println("Id", product.Id)
	fmt.Println("doc", rev)
	fmt.Println("err", err)

}
