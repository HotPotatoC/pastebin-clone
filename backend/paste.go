package backend

import (
	"context"
	"errors"
	"time"

	"github.com/HotPotatoC/pastebin-clone/logic"
	"github.com/HotPotatoC/pastebin-clone/repository"
	"github.com/gocql/gocql"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

type SavePasteParams struct {
	Text          []byte
	UserID        string
	UserIPAddress string
}

func (d Dependency) SavePaste(ctx context.Context, params SavePasteParams) (string, error) {
	compressedText, err := logic.Compress(params.Text)
	if err != nil {
		return "", err
	}

	b := make([]byte, 0)
	b = append(b, params.UserID[:]...)
	b = append(b, []byte(params.UserIPAddress)...)
	b = append(b, params.Text...)

	hash, err := logic.Hash(b)
	if err != nil {
		return "", err
	}

	// Generate short link from hash
	// shortLink = base62(sha256(userID + userIPAddress + text))[0:7]
	shortLink := logic.EncodeBase62([]byte(hash))[:7]

	// Check if paste already exists with the same hash and short link
	existingPaste, err := d.Repository.GetPasteByHash(ctx, hash, shortLink)
	if !errors.Is(err, gocql.ErrNotFound) {
		log.Debug().Msgf("Paste with hash %s already exists", hash)
		return existingPaste.ShortLink, nil
	}

	err = d.Repository.SavePaste(ctx, repository.Paste{
		Id:        uuid.New().String(),
		UserId:    params.UserID,
		Paste:     compressedText,
		ShortLink: shortLink,
		Hash:      hash,
		CreatedAt: time.Now(),
	})
	if err != nil {
		return "", err
	}

	return shortLink, nil
}

type GetPasteOutput struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	Username  string `json:"username"`
	Paste     string `json:"paste"`
	ShortLink string `json:"short_link"`
	CreatedAt string `json:"created_at"`
	Metadata  struct {
		Size              int    `json:"size"`
		SizeHumanReadable string `json:"size_human_readable"`
	}
}

func (d Dependency) GetPaste(ctx context.Context, shortLink string) (GetPasteOutput, error) {
	var output GetPasteOutput

	paste, err := d.Repository.GetPasteByShortLink(ctx, shortLink)
	if err != nil {
		return GetPasteOutput{}, err
	}

	user, err := d.Repository.GetUserByID(ctx, paste.UserId)
	if err != nil {
		return GetPasteOutput{}, err
	}

	decompressed, err := logic.Decompress([]byte(paste.Paste))
	if err != nil {
		return GetPasteOutput{}, err
	}

	output = GetPasteOutput{
		ID:        paste.Id,
		UserID:    paste.UserId,
		Username:  user.Name,
		Paste:     string(decompressed),
		ShortLink: paste.ShortLink,
		CreatedAt: paste.CreatedAt.Format("2006-01-02 15:04:05"),
	}

	output.Metadata.Size = len(decompressed)
	output.Metadata.SizeHumanReadable = logic.ByteCount(uint64(output.Metadata.Size))

	return output, nil
}
