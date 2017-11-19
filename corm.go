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
func (c *Orm) Save(ctx context.Context, doc interface{}) (docId string, docRev string, err error) {

	structName, docId, docRev, docType, err := CheckDocument(doc)
	if err != nil {
		return "", "", err
	}

	// docType must not be empty or it should be equal structName
	if docType != "" && structName != docType {
		return "", "", errDocAndTypeNotSameOrEmpty
	}

	// when docId already exists
	_, err = c.Db.Get(ctx, docId)
	if docId != "" && err == nil {
		return "", "", fmt.Errorf(errDocIdAlreadyExists, docId, structName)
	}

	// create doc with predefined id
	if docId != "" {
		docRev, err = c.Db.Put(ctx, docId, doc)
		return docId, docRev, err
	}

	// create doc with auto-generated id
	docId, docRev, err = c.Db.CreateDoc(ctx, doc)
	return docId, docRev, err
}

// Read reads document from database by passed id and doc struct.
// When it fails it will return an empty Row and error and when
// everything works fine, it will unmarshal the result (row) to
// passed doc struct.
func (c *Orm) Read(ctx context.Context, docId string, doc interface{}, options ...Options) (row *Row, err error) {

	if docId == "" {
		return nil, errDocIdRequired
	}
	row, err = c.Db.Get(ctx, docId, options...)
	if err != nil {
		return nil, errDocIdNotFound
	}

	// convert row to passed struct type
	row.ScanDoc(&doc)
	return row, err
}

// Update updates doc (struct) as document in the database. The doc must
// contains the doc Id and Rev property. When its saved correctly it will return a
// new revision. Otherwise it will return an error.
func (c *Orm) Update(ctx context.Context, doc interface{}) (docId, newRev string, err error) {

	structName, docId, docRev, docType, err := CheckDocument(doc)
	if err != nil {
		return "", "", err
	}

	// document type must not be empty or it should be equal structName
	if docType != "" && structName != docType {
		return "", "", errDocAndTypeNotSameOrEmpty
	}

	// document id and document revision are required
	if docId == "" && docRev == "" {
		return "", "", errDocIdAndRevRequired
	}

	newRev, err = c.Db.Put(ctx, docId, doc)
	return docId, newRev, err
}

// Delete deletes the document by passed id and revision (rev).
// When id and revision (rev) are empty it will return an error.
// When everything works fine it will return the new deleted revision (newRev).
func (c *Orm) Delete(ctx context.Context, docId string, docRev ...string) (newRev string, err error) {

	type IdRev struct {
		Id  string `json:"_id,omitempty"`
		Rev string `json:"_rev,omitempty"`
	}

	_docRev := ""
	docRevLen := len(docRev)

	// keep passed docRev but when no docRev passed
	// determine it by
	if docRevLen >= 1 {
		_docRev = docRev[0]
	} else {
		idRev := IdRev{}
		_, err = c.Read(ctx, docId, &idRev)
		if err != nil {
			return "", err
		}
		_docRev = idRev.Rev
	}

	if docId == "" {
		return "", errDocIdRequired
	}

	return c.Db.Delete(ctx, docId, _docRev)
}

// CheckDocument checks if the required fields Id, Rev and Type are exists.
// When everything gone right, it return the extracted fields Id, Rev and Type.
// When it fails it return an error.
func CheckDocument(doc interface{}) (structName, docId, docRev, docType string, err error) {

	structDoc := structs.New(doc)

	// Check require field props exists
	IdField, IdOk := structDoc.FieldOk("Id")
	RevField, RevOk := structDoc.FieldOk("Rev")
	TypeField, TypeOk := structDoc.FieldOk("Type")

	if !IdOk || !RevOk || !TypeOk {
		return "", "", "", "", fmt.Errorf(errDocIdRevTypeDynRequired, structDoc.Name())
	}

	docId = IdField.Value().(string)
	docRev = RevField.Value().(string)
	docType = TypeField.Value().(string)

	return structDoc.Name(), docId, docRev, docType, err
}
