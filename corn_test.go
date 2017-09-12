package corm

import (
	"context"
	"fmt"
	"log"
	"testing"
)

type Product struct {
	Id          string `json:"_id,omitempty"`
	Rev         string `json:"_rev,omitempty"`
	Type        string `json:"$_type,omitempty"`
	HasMany     string `json:"$_hasMany"`
	HasOny      string `json:"$_hasOne"`
	VendorId    int    `json:"vendorId"`
	VendorType  string `json:"vendorType"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CategoryId  string `json:"categoryId"`
	UrlId       string `json:"urlId"`
}

func TestCorn_Foo(t *testing.T) {

	ctx := context.Background()

	db, err := New(ctx, Config{
		DBName: "honeyglass",
	})

	if err != nil {
		log.Fatal(err)
	}

	docId, rev, err := db.Save(ctx, Product{
		Id:   "111-222-333",
		Name: "Liya",
	})

	db.SaveJson(ctx, `{
		"_id": 1234,
		"_type": "product",
		"_rev": "2233343-3342",
	}`)

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
	var product Product
	options := map[string]interface{}{"a": "b"}
	_, err = db.Read(ctx, docId, &product, options)

	db.First(ctx, &product)
	db.Last(ctx, &product)

	products := []Product{}
	db.FindAll(ctx, &products)

	var book Product
	db.FindBy(ctx, &book, "title", "Back to the Future")

	// delete doc
	rev, err = db.Delete(ctx, docId, rev)

	fmt.Println("Id", product.Id)
	fmt.Println("doc", rev)
	fmt.Println("err", err)

}
