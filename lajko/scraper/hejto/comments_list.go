package hejto

import "time"

type commentsListResponse struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
	Pages int `json:"pages"`
	Total int `json:"total"`
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
		First struct {
			Href string `json:"href"`
		} `json:"first"`
		Last struct {
			Href string `json:"href"`
		} `json:"last"`
	} `json:"_links"`
	Embedded struct {
		Items []struct {
			Content      string        `json:"content"`
			ContentPlain string        `json:"content_plain"`
			Images       []interface{} `json:"images"`
			ContentLinks []interface{} `json:"content_links"`
			PostSlug     string        `json:"post_slug"`
			Status       string        `json:"status"`
			Post         struct {
				Title        string `json:"title"`
				Content      string `json:"content"`
				ContentPlain string `json:"content_plain"`
				Excerpt      string `json:"excerpt"`
				Type         string `json:"type"`
				Slug         string `json:"slug"`
				Uuid         string `json:"uuid"`
				Discr        string `json:"discr"`
			} `json:"post"`
			Author struct {
				Username string `json:"username"`
				Avatar   struct {
					Urls struct {
						X100 string `json:"100x100"`
						X250 string `json:"250x250"`
					} `json:"urls"`
					Uuid string `json:"uuid"`
				} `json:"avatar"`
				Background struct {
					Urls struct {
						X300 string `json:"400x300"`
						X900 string `json:"1200x900"`
					} `json:"urls"`
					Uuid string `json:"uuid"`
				} `json:"background"`
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
			NumLikes   int  `json:"num_likes"`
			NumReports int  `json:"num_reports"`
			IsLiked    bool `json:"is_liked"`
			IsReported bool `json:"is_reported"`
			Replies    []struct {
				Content      string        `json:"content"`
				ContentPlain string        `json:"content_plain"`
				Images       []interface{} `json:"images"`
				ContentLinks []interface{} `json:"content_links"`
				PostSlug     string        `json:"post_slug"`
				Status       string        `json:"status"`
				Post         struct {
					Title        string `json:"title"`
					Content      string `json:"content"`
					ContentPlain string `json:"content_plain"`
					Excerpt      string `json:"excerpt"`
					Type         string `json:"type"`
					Slug         string `json:"slug"`
					Uuid         string `json:"uuid"`
					Discr        string `json:"discr"`
				} `json:"post"`
				Author struct {
					Username string `json:"username"`
					Avatar   struct {
						Urls struct {
							X100 string `json:"100x100"`
							X250 string `json:"250x250"`
						} `json:"urls"`
						Uuid string `json:"uuid"`
					} `json:"avatar"`
					Background struct {
						Urls struct {
							X300 string `json:"400x300"`
							X900 string `json:"1200x900"`
						} `json:"urls"`
						Uuid string `json:"uuid"`
					} `json:"background"`
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
					Sex string `json:"sex,omitempty"`
				} `json:"author"`
				NumLikes     int           `json:"num_likes"`
				NumReports   int           `json:"num_reports"`
				IsLiked      bool          `json:"is_liked"`
				IsReported   bool          `json:"is_reported"`
				Replies      []interface{} `json:"replies"`
				LikesEnabled bool          `json:"likes_enabled"`
				CreatedAt    time.Time     `json:"created_at"`
				Uuid         string        `json:"uuid"`
				Links        struct {
					Self struct {
						Href string `json:"href"`
					} `json:"self"`
					Likes struct {
						Href string `json:"href"`
					} `json:"likes"`
				} `json:"_links"`
			} `json:"replies"`
			LikesEnabled bool      `json:"likes_enabled"`
			CreatedAt    time.Time `json:"created_at"`
			Uuid         string    `json:"uuid"`
			Links        struct {
				Self struct {
					Href string `json:"href"`
				} `json:"self"`
				Likes struct {
					Href string `json:"href"`
				} `json:"likes"`
			} `json:"_links"`
		} `json:"items"`
	} `json:"_embedded"`
}
