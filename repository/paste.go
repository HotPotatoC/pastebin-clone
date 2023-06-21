package repository

import (
	"context"

	"github.com/scylladb/gocqlx/v2/qb"
)

func (d Dependency) SavePaste(ctx context.Context, paste Paste) error {
	q := qb.Insert(TablePastes.Name()).
		Columns(TablePastes.Metadata().Columns...).
		Query(d.DB)

	if err := q.BindStruct(paste).ExecRelease(); err != nil {
		return err
	}

	return nil
}

// GetPasteByHash returns a paste by its hash and short link.
// shortLink is used to ensure that the hash is unique.
func (d Dependency) GetPasteByHash(ctx context.Context, hash, shortLink string) (Paste, error) {
	var paste Paste
	q := qb.Select(TablePastes.Name()).
		Columns(TablePastes.Metadata().Columns...).
		Where(qb.Eq("short_link"), qb.Eq("hash")).
		Query(d.DB)

	if err := q.Bind(shortLink, hash).GetRelease(&paste); err != nil {
		return Paste{}, err
	}

	return paste, nil
}

func (d Dependency) GetPasteByShortLink(ctx context.Context, shortLink string) (Paste, error) {
	var paste Paste
	q := qb.Select(TablePastes.Name()).
		Columns(TablePastes.Metadata().Columns...).
		Where(qb.Eq("short_link")).
		Query(d.DB)

	if err := q.Bind(shortLink).GetRelease(&paste); err != nil {
		return Paste{}, err
	}

	return paste, nil
}
