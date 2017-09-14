package corm

import (
	"context"
	"fmt"
	"github.com/fatih/structs"
)

// Orm contains the Db and Cache state
type Orm struct {
	Db    *DB
	Cache *map[string]interface{}
}

// OrmOptions contains key value pair as option
type OrmOptions map[string]interface{}

// NewOrm creates a new Orm instance by passed db instance
func NewOrm(db *DB) *Orm {
	// returns Orm instance
	return &Orm{
		Db: db,
	}
}

// Save save new doc (struct) in the database. When it saved correctly
// it will return an id and revision. On fail it will return an error.
// When doc contain the Id property, the database will interpret as
// predefined unquie Id and when its omitted it will generate it automatically.
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

// Read read document from database by passed id and doc struct.
// When it fails it will return an empty Row and error and when
// everything works fine, it will unmarshal the result (row) to
// passed doc struct.
func (c *Orm) Read(ctx context.Context, id string, doc interface{}, options ...Options) (row *Row, err error) {

	if id == "" {
		return &Row{}, errDocIdRequired
	}
	row, err = c.Db.Get(ctx, id, options...)
	// convert row to passed struct type
	row.ScanDoc(&doc)
	return row, err
}

// Update update doc (struct) as document in the database. The doc must
// contains the doc Id and Rev property. When its saved correctly it will return a
// new revision. Otherwise it will return an error.
func (c *Orm) Update(ctx context.Context, doc interface{}) (newRev string, err error) {

	// extract id and rev from empty interface
	structDoc := structs.New(doc)
	id := structDoc.Field("Id").Value().(string)
	rev := structDoc.Field("Rev").Value().(string)

	if id == "" && rev == "" {
		return "", errDocIdAndRevRequired
	}
	return c.Db.Put(ctx, id, doc)
}

// Delete deletes the document by passed id and revision (rev).
// When id and revision (rev) are empty it will return an error.
// When everything works fine will return the new deleted revision (newRev).
func (c *Orm) Delete(ctx context.Context, id, rev string) (newRev string, err error) {

	if id == "" && rev == "" {
		return "", errDocIdAndRevRequired
	}
	return c.Db.Delete(ctx, id, rev)
}
