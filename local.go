// +build: localhost
package main

import "os"

func init() {
	postgresDbUrl = os.Getenv("POSTGRESQL_URL")
}
