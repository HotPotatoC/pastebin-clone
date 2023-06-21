package repository

import (
	"time"

	"github.com/scylladb/gocqlx/v2/table"
)

type Paste struct {
	Id        string
	UserId    string
	Paste     []byte
	ShortLink string
	Hash      string
	CreatedAt time.Time
}
type User struct {
	Id        string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}

// Table models.
var (
	TablePastes = table.New(table.Metadata{
		Name: "pastes",
		Columns: []string{
			"id",
			"user_id",
			"paste",
			"short_link",
			"hash",
			"created_at",
		},
		PartKey: []string{
			"id",
		},
		SortKey: []string{
			"short_link",
			"hash",
			"created_at",
		},
	})

	TableUsers = table.New(table.Metadata{
		Name: "users",
		Columns: []string{
			"id",
			"name",
			"email",
			"password",
			"created_at",
		},
		PartKey: []string{
			"id",
		},
		SortKey: []string{
			"email",
			"created_at",
		},
	})
)
