package corm

import (
	"context"
	"fmt"
	"github.com/fatih/structs"
)

// NewOrm creates a new Orm instance by passed db instance.
// Note: not relevant for you! It will be used inside of corm.New(...) but if you
// really want to create your own orm, please see [Example].
func NewOrm(db *DB) *Orm {
	// returns Orm instance
	return &Orm{
		Db: db,
	}
}

// OrmOptions contains options as key value pair.
type OrmOptions map[string]interface{}

// Orm contains the Db instance that be used inside of corm.NewOrm(...)
type Orm struct {
	Db *DB
}

// Save saves new doc (struct) in the database. When it saved correctly
// it will return an id and revision. On fail it will return an error.
// When doc contain the Id property, the database will interpret it as
// predefined unique Id and when its omitted it will generate it automatically.
func (c *Orm) Save(ctx context.Context, doc interface{}) (newId string, rev string, err error) {

	structDoc := structs.New(doc)
	id := structDoc.Field("Id").Value().(string)
	Type := structDoc.Field("Type").Value().(string)

	// 1.) Das Feld muss so in der Form existieren
	// 1.) Darf kein Inhalt haben weil es intern genutzt wird
	fmt.Fprintln(Type)

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

// Read reads document from database by passed id and doc struct.
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

// Update updates doc (struct) as document in the database. The doc must
// contains the doc Id and Rev property. When its saved correctly it will return a
// new revision. Otherwise it will return an error.
func (c *Orm) Update(ctx context.Context, doc interface{}) (newRev string, err error) {

	// extract id and rev from empty interface
	structDoc := structs.New(doc)
	id := structDoc.Field("Id").Value().(string)
	rev := structDoc.Field("Rev").Value().(string)
	Type := structDoc.Field("Type").Value().(string)

	// 1.) Das Feld muss so in der Form existieren
	// 1.) Darf kein Inhalt haben weil es intern genutzt wird
	fmt.Fprintln(Type)

	if id == "" && rev == "" {
		return "", errDocIdAndRevRequired
	}
	return c.Db.Put(ctx, id, doc)
}

// Delete deletes the document by passed id and revision (rev).
// When id and revision (rev) are empty it will return an error.
// When everything works fine it will return the new deleted revision (newRev).
func (c *Orm) Delete(ctx context.Context, id, rev string) (newRev string, err error) {

	if id == "" && rev == "" {
		return "", errDocIdAndRevRequired
	}
	return c.Db.Delete(ctx, id, rev)
}
