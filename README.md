# testdb
Utility to make [go-mysql-server](https://github.com/dolthub/go-mysql-server) easier to use in unit tests.

go-mysql-server is a MySQL-compatible in-memory database.
This database is ideal for unit testing as it operates within the same process.

## Example

`testdb.New()` prepares a `*sql.DB` that can be used in the normal way.

```go
package main

import (
	"fmt"

	"github.com/makiuchi-d/testdb"
)

const (
	sqlCreate = `CREATE TABLE people (
		id    integer NOT NULL AUTO_INCREMENT,
		name  varchar(128) NOT NULL,
		PRIMARY KEY (id)
	) Engine=InnoDB`
	sqlInsert = `INSERT INTO people (id, name) VALUES (?, ?)`
	sqlSelect = `SELECT id, name FROM people WHERE id <= ?`
)

func main() {
	db := testdb.New("my_test_db")

	db.Exec(sqlCreate)
	db.Exec(sqlInsert, 1, "Alice")
	db.Exec(sqlInsert, 2, "Bob")
	db.Exec(sqlInsert, 3, "Carol")

	var (
		id   int
		name string
	)
	rows, _ := db.Query(sqlSelect, 2)
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&id, &name)
		fmt.Println(id, name)
	}
}
```

### Examples of use with ORM

- [ent](https://github.com/makiuchi-d/testdb/blob/main/example-ent/main_test.go)
- [GORM](https://github.com/makiuchi-d/testdb/blob/main/example-gorm/main_test.go)
- [sqlc](https://github.com/makiuchi-d/testdb/blob/main/example-sqlc/main_test.go)
- [sqlx](https://github.com/makiuchi-d/testdb/blob/main/example-sqlx/main_test.go)

