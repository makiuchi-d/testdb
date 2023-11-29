package testdb

import (
	"database/sql"
	"strings"

	"github.com/dolthub/go-mysql-server/driver"
	"github.com/dolthub/go-mysql-server/memory"
	sqle "github.com/dolthub/go-mysql-server/sql"
)

type dbs []sqle.Database

var _ driver.Provider = dbs{}

func (d dbs) Resolve(name string, options *driver.Options) (string, sqle.DatabaseProvider, error) {
	return name, memory.NewDBProvider(d...), nil
}

func New(dbNames ...string) *sql.DB {
	var memdbs dbs
	for _, dbName := range dbNames {
		memdb := memory.NewDatabase(dbName)
		memdb.EnablePrimaryKeyIndexes()
		memdbs = append(memdbs, memdb)
	}
	drv := driver.New(memdbs, nil)
	conn := Must1(drv.OpenConnector(strings.Join(dbNames, ";")))
	db := sql.OpenDB(conn)
	if len(dbNames) > 0 {
		Must1(db.Exec("USE " + dbNames[0]))
	}
	return db
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
