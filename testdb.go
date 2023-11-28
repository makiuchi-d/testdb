package testdb

import (
	"database/sql"
	"os"

	"github.com/dolthub/go-mysql-server/driver"
	"github.com/dolthub/go-mysql-server/memory"
	sqle "github.com/dolthub/go-mysql-server/sql"
)

type dbs []sqle.Database

func (d dbs) Resolve(name string, options *driver.Options) (string, sqle.DatabaseProvider, error) {
	return name, memory.NewDBProvider(d...), nil
}

func New(dbName string) *sql.DB {
	memdb := memory.NewDatabase(dbName)
	memdb.EnablePrimaryKeyIndexes()
	drv := driver.New(dbs{memdb}, nil)
	conn := Must1(drv.OpenConnector(dbName))
	db := sql.OpenDB(conn)
	Must1(db.Exec("USE " + dbName))
	return db
}

func ApplySQLFile(db *sql.DB, file string) {
	s := string(Must1(os.ReadFile(file)))
	Must1(db.Exec(s))
}

func Must(err error) {
	if err != nil {
		panic(err)
	}
}

func Must1[T any](t T, err error) T {
	Must(err)
	return t
}

func Must2[T, U any](t T, u U, err error) (T, U) {
	Must(err)
	return t, u
}
