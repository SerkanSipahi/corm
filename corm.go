package corm

import (
	"context"
	"fmt"
	"github.com/fatih/structs"
)

type Orm struct {
	Db *DB
}

func NewOrm(db *DB) *Orm {
	return &Orm{
		Db: db,
	}
}

func (c *Orm) Save(ctx context.Context, doc interface{}) (newId string, rev string, err error) {

	structDoc := structs.New(doc)
	id := structDoc.Field("Id").Value().(string)

	// when docId already exists
	_, err = c.Db.Get(ctx, id)
	if id != "" && err == nil {
		return "", "", fmt.Errorf(errDocIdAlreadyExists, id, structDoc.Name())
	}

	// create doc with predefined id
	if id != "" {
		rev, err = c.Db.Put(ctx, id, doc)
		return id, rev, err
	}

	// create doc with auto-generated id
	docId, rev, err := c.Db.CreateDoc(ctx, doc)
	return docId, rev, err
}

func (c *Orm) Read(ctx context.Context, id string, doc interface{}, options ...Options) (row *Row, err error) {

	if id == "" {
		return &Row{}, errDocIdRequired
	}
	row, err = c.Db.Get(ctx, id, options...)
	row.ScanDoc(&doc)
	return row, err
}

func (c *Orm) Update(ctx context.Context, doc interface{}) (newRev string, err error) {

	structDoc := structs.New(doc)
	id := structDoc.Field("Id").Value().(string)
	rev := structDoc.Field("Rev").Value().(string)

	if id == "" && rev == "" {
		return "", errDocIdAndRevRequired
	}
	return c.Db.Put(ctx, id, doc)
}

func (c *Orm) UpdateAll(ctx context.Context, id, rev string) {

}

func (c *Orm) Delete(ctx context.Context, id, rev string) (newRev string, err error) {

	if id == "" && rev == "" {
		return "", errDocIdAndRevRequired
	}
	return c.Db.Delete(ctx, id, rev)
}
