package repository

import "github.com/scylladb/gocqlx/v2"

type Dependency struct {
	DB gocqlx.Session
}
