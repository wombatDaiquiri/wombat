package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/wombatDaiquiri/wombat/lajko"
	"github.com/wombatDaiquiri/wombat/lajko/database/models"
)

type ContentStorage struct {
	db           *pgxpool.Pool
	actorStorage actorStorage
}

type actorStorage interface {
	UpsertActor(author models.Actor) (models.Actor, error)
}

func (s *ContentStorage) ContentAlreadyStored(ctx context.Context, uri string) (bool, error) {
	// source: https://donchev.is/post/working-with-postgresql-in-go-using-pgx/
	var exists bool
	err := s.db.QueryRow(ctx, "SELECT EXISTS(SELECT 1 FROM content WHERE uri = $1)", uri).Scan(&exists)
	if err != nil {
		return false, err

	}
	return exists, nil
}

func (s *ContentStorage) UpsertContent(ctx context.Context, content lajko.Content) (lajko.Content, error) {
	// source: https://donchev.is/post/working-with-postgresql-in-go-using-pgx/
	// upsert author -> upsert content -> (SEPARATE EXTENSION/PR/FEATURE: upsert liking people -> upsert likes)

	_, err := s.actorStorage.UpsertActor(models.Actor{
		URI:           content.Author.URI,
		DisplayedName: content.Author.DisplayedName,
		AvatarURI:     content.Author.AvatarURI,
	})
	if err != nil {
		return lajko.Content{}, fmt.Errorf("could not upsert actor: %w", err)
	}

	insertContentQuery := `INSERT INTO
		content (uri, author_uri, in_response_to_uri, parser_version, lajko_html, attachments)
		VALUES (@uri, @author_uri, @in_response_to_uri, @parser_version, @lajko_html, @attachments)
		ON CONFLICT (uri) DO UPDATE SET
			author_uri = @author_uri,
			in_response_to_uri = @in_response_to_uri,
			parser_version = @parser_version,
			lajko_html = @lajko_html,
			attachments = @attachments;`
	insertContentArgs := pgx.NamedArgs{
		"uri":                content.URI,
		"author_uri":         content.Author.URI,
		"in_response_to_uri": content.InResponseToURI,
		"parser_version":     content.ParserVersion,
		"lajko_html":         content.LajkoHTML,
		"attachments":        content.Attachments,
	}

	_, err = s.db.Exec(ctx, insertContentQuery, insertContentArgs)
	if err != nil {
		return lajko.Content{}, fmt.Errorf("could not upsert content: %w", err)
	}

	return content, nil
}
