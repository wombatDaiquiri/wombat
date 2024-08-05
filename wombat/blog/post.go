package blog

import (
	"time"

	"github.com/wombatDaiquiri/wombat/wombat/database"
	"github.com/wombatDaiquiri/wombat/ww"
)

func NewPosts(blogposts []database.Blogpost) []Post {
	return ww.Map(blogposts, NewPost)
}

func NewPost(blogpost database.Blogpost) Post {
	return Post{
		ULID:            blogpost.ULID,
		Title:           blogpost.Title,
		Slug:            blogpost.Slug,
		Lead:            blogpost.Lead,
		PublicationDate: blogpost.PublicationDate,

		ContentMarkdown: blogpost.ContentMarkdown,
	}
}

type Post struct {
	ULID            string     `json:"ulid"`
	Title           string     `json:"title"`
	Slug            string     `json:"slug"`
	Lead            string     `json:"lead"`
	PublicationDate *time.Time `json:"publicationDate"`

	ContentMarkdown string `json:"contentMarkdown"`
}

// TODO (idea): sanitize output based on something in context?
