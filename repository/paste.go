package repository

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
	"github.com/scylladb/gocqlx/v2/qb"
)

func (d Dependency) SavePaste(ctx context.Context, paste Paste) error {
	q := qb.Insert(TablePastes.Name()).
		Columns(TablePastes.Metadata().Columns...).
		Query(d.DB)

	if err := q.BindStruct(paste).ExecRelease(); err != nil {
		return err
	}

	cacheKey := fmt.Sprintf("%s:%s", TablePastes.Name(), paste.ShortLink)
	d.Redis.Pipelined(ctx, func(rdb redis.Pipeliner) error {
		rdb.HSet(ctx, cacheKey, "id", paste.Id)
		rdb.HSet(ctx, cacheKey, "short_link", paste.ShortLink)
		rdb.HSet(ctx, cacheKey, "hash", paste.Hash)
		rdb.HSet(ctx, cacheKey, "created_at", paste.CreatedAt)
		rdb.HSet(ctx, cacheKey, "paste", paste.Paste)
		rdb.HSet(ctx, cacheKey, "user_id", paste.UserId)
		return nil
	})

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

	cacheKey := fmt.Sprintf("%s:%s", TablePastes.Name(), shortLink)
	isCached, _ := d.Redis.Exists(ctx, cacheKey).Result()
	if isCached == 1 {
		err := d.Redis.HGetAll(ctx, cacheKey).Scan(&paste)
		if err != nil {
			log.Error().Err(err).Msgf("failed to retrieve paste from cache for %s", shortLink)
		} else {
			log.Debug().Any("paste", paste).Msgf("cache hit for %s", shortLink)
			return paste, nil
		}
	}

	log.Debug().Msgf("cache miss for %s (result: %d); retrieving from database", shortLink, isCached)

	q := qb.Select(TablePastes.Name()).
		Columns(TablePastes.Metadata().Columns...).
		Where(qb.Eq("short_link")).
		Query(d.DB)

	if err := q.Bind(shortLink).GetRelease(&paste); err != nil {
		return Paste{}, err
	}

	return paste, nil
}
