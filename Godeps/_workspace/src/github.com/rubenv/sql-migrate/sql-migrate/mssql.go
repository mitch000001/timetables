// +build go1.3

package main

import (
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/mitch000001/timetables/Godeps/_workspace/src/github.com/go-gorp/gorp"
)

func init() {
	dialects["mssql"] = gorp.SqlServerDialect{}
}
