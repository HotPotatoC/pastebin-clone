package repository

import (
	"github.com/redis/go-redis/v9"
	"github.com/scylladb/gocqlx/v2"
)

type Dependency struct {
	DB    gocqlx.Session
	Redis *redis.Client
}
