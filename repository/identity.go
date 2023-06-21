package repository

import (
	"context"

	"github.com/scylladb/gocqlx/v2/qb"
)

func (d Dependency) SaveUser(ctx context.Context, identity User) error {
	q := d.DB.Query(TableUsers.Insert()).BindStruct(identity)
	if err := q.Exec(); err != nil {
		return err
	}

	return nil
}

func (d Dependency) GetUserByEmail(ctx context.Context, email string) (User, error) {
	var identity User

	q := qb.Select(TableUsers.Name()).
		Columns(TableUsers.Metadata().Columns...).
		Where(qb.Eq("email")).
		Limit(1).
		Query(d.DB)

	err := q.Bind(email).GetRelease(&identity)
	if err != nil {
		return User{}, err
	}

	return identity, nil
}

func (d Dependency) GetUserByID(ctx context.Context, userID string) (User, error) {
	var identity User

	q := qb.Select(TableUsers.Name()).
		Columns(TableUsers.Metadata().Columns...).
		Where(qb.Eq("id")).
		Limit(1).
		Query(d.DB)

	err := q.Bind(userID).GetRelease(&identity)
	if err != nil {
		return User{}, err
	}

	return identity, nil
}
