package testdb

import (
	"database/sql"
	"os"

	"github.com/dolthub/go-mysql-server/driver"
	"github.com/dolthub/go-mysql-server/memory"
	sqle "github.com/dolthub/go-mysql-server/sql"
	"github.com/dolthub/go-mysql-server/sql/information_schema"
)

type factory struct{}

func (factory) Resolve(name string, options *driver.Options) (string, sqle.DatabaseProvider, error) {
	memdb := memory.NewDatabase(name)
	memdb.EnablePrimaryKeyIndexes()

	provider := memory.NewDBProvider(
		memdb,
		information_schema.NewInformationSchemaDatabase(),
	)

	return name, provider, nil
}

func init() {
	sql.Register("sqle", driver.New(factory{}, nil))
}

func New(dbName string) *sql.DB {
	db := Must1(sql.Open("sqle", dbName))
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
