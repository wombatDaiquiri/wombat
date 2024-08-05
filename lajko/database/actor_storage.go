package database

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/wombatDaiquiri/wombat/lajko/database/models"
)

type ActorStorage struct {
	db *pgxpool.Pool
}

func (as *ActorStorage) UpsertActor(ctx context.Context, actor models.Actor) (models.Actor, error) {
	insertAuthorQuery := `INSERT INTO
    		actors (uri, displayed_name, avatar_uri)
    		VALUES (@uri, @displayed_name, @avatar_uri)
    		ON CONFLICT (uri) DO UPDATE SET
				displayed_name = @displayed_name,
				avatar_uri = @avatar_uri;`
	insertAuthorArgs := pgx.NamedArgs{
		"uri":            actor.URI,
		"displayed_name": actor.DisplayedName,
		"avatar_uri":     actor.AvatarURI,
	}

	_, err := as.db.Exec(ctx, insertAuthorQuery, insertAuthorArgs)
	if err != nil {
		return models.Actor{}, err
	}
	return actor, nil
}
