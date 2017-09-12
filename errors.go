package corm

import "errors"

var (
	errDocIdRequired       = errors.New("doc.Id are required")
	errDocIdAndRevRequired = errors.New("doc.Id and doc.rev are required")
	errDocIdAlreadyExists  = "doc.Id: %v for `%v` struct document already created"
)
