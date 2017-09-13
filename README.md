# CORM

The awesome CouchDB ORM library for Golang, aims to be developer friendly.

[![GoDoc](https://godoc.org/github.com/SerkanSipahi/corm?status.svg)](https://godoc.org/github.com/SerkanSipahi/corm)

## Overview

* Our goal is to adapt/implement the Domain-Classes Methods from [grails](http://docs.grails.org/latest/ref/Domain%20Classes/save.html) and some of the CouchDB specific [Api](https://godoc.org/github.com/flimzy/kivik) for CouchDb
* Be careful when using. The api can changed. We are in early state.
* If you have some suggestions or concerns please contact us or make a issue ticket.

### ORM methods (our goal for a Nosql-Database)

#### db methods

```golang
ctx := context.Background()
db, err := corm.New(ctx, Config{
    DBName: "myDatabase",
})
```

- [ ] BelongsTo
- [ ] Count
- [ ] CountBy
- [x] Delete
- [ ] DeleteMany
- [ ] ExecuteQuery
- [x] Update
- [ ] UpdateAll
- [ ] Exists
- [ ] Find
- [ ] FindAll
- [ ] FindAllBy
- [ ] FindAllWhere
- [ ] FindBy
- [ ] FindWhere
- [ ] First
- [ ] Get
- [ ] GetAll
- [ ] HasMany
- [ ] HasOne
- [ ] Last
- [ ] List
- [ ] ListOrderBy
- [x] Read
- [ ] Refresh
- [x] Save
- [ ] SaveAll
- [ ] SaveJson
- [ ] Sync
- [ ] Validate
- [ ] Where
- [ ] WhereAny

## Getting Started

* coming soon

### Basic usage
```golang

// init DB
ctx := context.Background()
db, err := corm.New(ctx, Config{
    DBName: "myDatabase",
})

// save document with custom Id
docId, rev, err := db.Save(ctx, Product{
    Id:   "111-222-333",
    Name: "Foo",
})

// save document with auto Id
docId, rev, err = db.Save(ctx, Product{
    Name: "Bar",
})

// update document
rev, err = db.Update(ctx, Product{
    Id:   docId,
    Rev:  rev,
    Name: "Baz",
})

// read document
var product Product
_, err = db.Read(ctx, docId, &product)
fmt.Println(product) // product{ Id: "asdfj334234f34asdfq34", Rev: "1-alsj34lkjij3lksife" ...

// delete document
rev, err = db.Delete(ctx, docId, rev)
```

## Author

**SerkanSipahi**

* <http://github.com/SerkanSipahi>
* <serkan.sipahi@yahoo.de>
* <https://twitter.com/Bitcollage>

## Contributors

https://github.com/SerkanSipahi/corm/graphs/contributors

## License

This software is released under the terms of the Apache 2.0 license. See LICENCE.md, or read the [full license](http://www.apache.org/licenses/LICENSE-2.0).