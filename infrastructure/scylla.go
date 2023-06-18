package infrastructure

import (
	"context"

	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
)

func NewScylla(ctx context.Context, keyspace string, hosts []string) (gocqlx.Session, error) {
	cluster := gocql.NewCluster(hosts...)
	cluster.Keyspace = keyspace

	session, err := gocqlx.WrapSession(cluster.CreateSession())
	if err != nil {
		return gocqlx.Session{}, err
	}

	return session, nil
}
