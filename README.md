# GQT - Go(lang) SQL Templates

[![Go Report Card](https://goreportcard.com/badge/github.com/Davmuz/gqt)](https://goreportcard.com/report/github.com/Davmuz/gqt) [![GoDoc](https://godoc.org/github.com/Davmuz/gqt?status.svg)](https://godoc.org/github.com/Davmuz/gqt)

Package gqt is a template engine for SQL queries.

It helps to separate SQL code from Go code and permits to compose the queries
with a simple syntax.

The template engine is the standard package "text/template".

Why this package?
Read more about [ORM is the Vietnam of computer science](http://blog.codinghorror.com/object-relational-mapping-is-the-vietnam-of-computer-science/).

Install/update using go get (no dependencies required by gqt):

```bash
go get -u github.com/Davmuz/gqt
```

#### Benefits

- SQL is the best language to write SQL.
- Separation between Go and SQL source code (your DB administrator will thank
you).
- Simpler template syntax for composing queries than writing Go code.
- Simplified maintenance of the SQL code.

#### Compatibility

Go >= 1.6

## Usage


Create a template directory tree of .sql files. Here an example template with
the definition of three blocks:

```sql
-- File /path/to/sql/repository/dir/example.sql
{{define "allUsers"}}
SELECT *
FROM users
WHERE 1=1
{{end}}

{{define "getUser"}}
SELECT *
FROM users
WHERE id=?
{{end}}

{{define "allPosts"}}
SELECT *
FROM posts
WHERE date>=?
{{if ne .Order ""}}ORDER BY date {{.Order}}{{end}}
{{end}}
```

Then, with Go, add the directory to the default repository and execute the
queries:

```go
// Setup
gqt.Add("/path/to/sql/repository/dir", "*.sql")

// Simple query without parameters
db.Query(gqt.Get("allUsers"))
// Query with parameters
db.QueryRow(gqt.Get("getuser"), 1)
// Query with context and parameters
db.Query(gqt.Exec("allPosts", map[string]interface{}{
	"Order": "DESC",
}), date)
```

The templates are parsed immediately and recursively.

## Namespaces

The templates can be organized in namespaces and stored in multiple root
directories.

```
templates1/
|-- roles/
|	|-- queries.sql
|-- users/
|	|-- queries.sql
|	|-- commands.sql

templates2/
|-- posts/
|	|-- queries.sql
|	|-- commands.sql
|-- users/
|	|-- queries.sql
|-- queries.sql
```

The blocks inside the sql files are merged, the blocks with the same namespace
and name will be overridden following the alphabetical order.

The sub-directories are used as namespaces and accessed like:

```go
gqt.Add("../templates1", "*.sql")
gqt.Add("../templates2", "*.sql")

// Will search inside templates1/users/*.sql and templates2/users/*.sql
gqt.Get("users/allUsers")
```

## Multiple databases

When dealing with multiple databases at the same time, like PostgreSQL and
MySQL, just create two repositories:

```go
// Use a common directory
dir := "/path/to/sql/repository/dir"

// Create the PostgreSQL repository
pgsql := gqt.NewRepository()
pgsql.Add(dir, "*.pg.sql")

// Create a separated MySQL repository
mysql := gqt.NewRepository()
mysql.Add(dir, "*.my.sql")

// Then execute
pgsql.Get("queryName")
mysql.Get("queryName")
```

## License

Copyright © 2016 Davide Muzzarelli. All right reserved.

Use of this source code is governed by a BSD-style license that can be found in the [LICENSE](LICENSE) file.
