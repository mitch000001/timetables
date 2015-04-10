// +build heroku

package main

import "os"

func init() {
	postgresDbUrl = os.Getenv("HEROKU_POSTGRESQL_ROSE_URL")
}
