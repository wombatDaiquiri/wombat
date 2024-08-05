package database

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type BlogpostStorage struct {
	db *pgxpool.Pool
}

func (bs *BlogpostStorage) GetBlogpost() {

}

func (bs *BlogpostStorage) GetBlogposts() {

}

func (bs *BlogpostStorage) CreateBlogpost(ctx context.Context, blogpost Blogpost) (Blogpost, error) {
	return Blogpost{}, nil
}

type Blogpost struct {
	ID        int       `db:"id"`
	CreatedAt time.Time `db:"created_at"`

	ULID            string     `db:"ulid"`
	Title           string     `db:"title"`
	Slug            string     `db:"slug"`
	Lead            string     `db:"lead"`
	PublicationDate *time.Time `db:"publication_date"`

	ContentMarkdown string `db:"content_markdown"`
}
