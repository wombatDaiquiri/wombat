package hejto

import (
	"fmt"
	"net/url"
	"time"

	"github.com/gomarkdown/markdown"
	"github.com/microcosm-cc/bluemonday"

	"github.com/wombatDaiquiri/wombat/lajko"
)

type postResponse struct {
	Title        string `json:"title"`
	Content      string `json:"content"`
	ContentPlain string `json:"content_plain"`
	Excerpt      string `json:"excerpt"`
	Images       []struct {
		Urls struct {
			X250 string `json:"250x250"`
			X500 string `json:"500x500"`
			X900 string `json:"1200x900"`
		} `json:"urls"`
		Uuid     string `json:"uuid"`
		Position int    `json:"position"`
	} `json:"images"`
	ContentLinks []interface{} `json:"content_links"`
	Comments     []interface{} `json:"comments"`
	Type         string        `json:"type"`
	Slug         string        `json:"slug"`
	Status       string        `json:"status"`
	Hot          bool          `json:"hot"`
	Community    struct {
		Name         string `json:"name"`
		Slug         string `json:"slug"`
		PrimaryColor string `json:"primary_color"`
	} `json:"community"`
	Author struct {
		Username string `json:"username"`
		Avatar   struct {
			Urls struct {
				X100 string `json:"100x100"`
				X250 string `json:"250x250"`
			} `json:"urls"`
			Uuid string `json:"uuid"`
		} `json:"avatar"`
		Status        string    `json:"status"`
		Roles         []string  `json:"roles"`
		Controversial bool      `json:"controversial"`
		CurrentRank   string    `json:"current_rank"`
		CurrentColor  string    `json:"current_color"`
		Verified      bool      `json:"verified"`
		Sponsor       bool      `json:"sponsor"`
		ExSponsor     bool      `json:"ex_sponsor"`
		CreatedAt     time.Time `json:"created_at"`
		Links         struct {
			Self struct {
				Href string `json:"href"`
			} `json:"self"`
			Follows struct {
				Href string `json:"href"`
			} `json:"follows"`
		} `json:"_links"`
	} `json:"author"`
	Tags []struct {
		Name       string `json:"name"`
		NumFollows int    `json:"num_follows"`
		NumPosts   int    `json:"num_posts"`
		Links      struct {
			Self struct {
				Href string `json:"href"`
			} `json:"self"`
			Follows struct {
				Href string `json:"href"`
			} `json:"follows"`
			Blocks struct {
				Href string `json:"href"`
			} `json:"blocks"`
		} `json:"_links"`
	} `json:"tags"`
	Nsfw             bool      `json:"nsfw"`
	Controversial    bool      `json:"controversial"`
	WarContent       bool      `json:"war_content"`
	Masked           bool      `json:"masked"`
	NumLikes         int       `json:"num_likes"`
	NumComments      int       `json:"num_comments"`
	NumFavorites     int       `json:"num_favorites"`
	IsLiked          bool      `json:"is_liked"`
	IsCommented      bool      `json:"is_commented"`
	IsFavorited      bool      `json:"is_favorited"`
	Uuid             string    `json:"uuid"`
	CommentsEnabled  bool      `json:"comments_enabled"`
	FavoritesEnabled bool      `json:"favorites_enabled"`
	LikesEnabled     bool      `json:"likes_enabled"`
	ReportsEnabled   bool      `json:"reports_enabled"`
	SharesEnabled    bool      `json:"shares_enabled"`
	CreatedAt        time.Time `json:"created_at"`
	Discr            string    `json:"discr"`
}

func (resp postResponse) LajkoHTML() string {
	maybeUnsafeHTML := markdown.ToHTML([]byte(resp.Content), nil, nil)
	html := bluemonday.UGCPolicy().SanitizeBytes(maybeUnsafeHTML)
	return string(html)
}

func (resp postResponse) LajkoAttachments() []lajko.Attachment {
	attachments := make([]lajko.Attachment, 0, len(resp.Images))
	for _, img := range resp.Images {
		imgURI, err := url.Parse(img.Urls.X900)
		if err != nil {
			fmt.Println("error parsing image URI:", err)
			continue
		}
		attachments = append(attachments, lajko.Attachment{
			URI:         imgURI,
			ContentType: "image/png",
		})
	}
	return attachments
}
