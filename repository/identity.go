package repository

import (
	"context"
)

func (d Dependency) SaveUser(ctx context.Context, identity UsersStruct) error {
	q := d.DB.Query(Users.Insert()).BindStruct(identity)
	if err := q.Exec(); err != nil {
		return err
	}

	return nil
}

func (d Dependency) GetUserByEmail(ctx context.Context, email string) (UsersStruct, error) {
	var identity UsersStruct
	err := d.DB.Query("SELECT * FROM users WHERE email = ? LIMIT 1", nil).Bind(email).Get(&identity)
	if err != nil {
		return UsersStruct{}, err
	}

	return identity, nil
}
