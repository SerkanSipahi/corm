package corm

import "errors"

var (
	errDocIdRequired            = errors.New("doc.Id are required")
	errDocIdNotFound            = errors.New("doc.Id not found")
	errDocIdAndRevRequired      = errors.New("doc.Id and doc.rev are required")
	errDocAndTypeNotSameOrEmpty = errors.New("doc.Type is reserved for corm")
	errDocIdRevTypeDynRequired  = "doc.Id, doc.Rev and doc.Type: for `%v` struct document are required"
	errDocIdAlreadyExists       = "doc.Id: %v for `%v` struct document already created"
)
