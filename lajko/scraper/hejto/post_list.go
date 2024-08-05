package hejto

import "time"

// postListResponse is the structure of response from /posts endpoint.
//
// This was parsed automatically by Goland from copy-pasting the response from hejto API.
type postListResponse struct {
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
		Next struct {
			Href string `json:"href"`
		} `json:"next"`
		Previous struct {
			Href string `json:"href"`
		} `json:"previous"`
	} `json:"_links"`
	Embedded struct {
		Items []struct {
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
			ContentLinks []struct {
				Url    string        `json:"url"`
				Site   string        `json:"site"`
				Type   string        `json:"type"`
				Title  string        `json:"title"`
				Audios []interface{} `json:"audios"`
				Images []interface{} `json:"images"`
				Videos []interface{} `json:"videos"`
				// Favicon is an empty array if empty (?)
				// Favicon struct {
				// 	Url  string `json:"url"`
				// 	Safe string `json:"safe"`
				// } `json:"favicon"`
			} `json:"content_links"`
			Comments []struct {
				Content      string `json:"content"`
				ContentPlain string `json:"content_plain"`
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
				PostSlug     string        `json:"post_slug"`
				Status       string        `json:"status"`
				Author       struct {
					Username string `json:"username"`
					Sex      string `json:"sex,omitempty"`
					Avatar   struct {
						Urls struct {
							X100 string `json:"100x100"`
							X250 string `json:"250x250"`
						} `json:"urls"`
						Uuid string `json:"uuid"`
					} `json:"avatar,omitempty"`
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
					Author       struct {
						Username string `json:"username"`
						Sex      string `json:"sex,omitempty"`
						Avatar   struct {
							Urls struct {
								X100 string `json:"100x100"`
								X250 string `json:"250x250"`
							} `json:"urls"`
							Uuid string `json:"uuid"`
						} `json:"avatar,omitempty"`
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
				UpdatedAt time.Time `json:"updated_at,omitempty"`
			} `json:"comments"`
			Type      string `json:"type"`
			Slug      string `json:"slug"`
			Status    string `json:"status"`
			Hot       bool   `json:"hot"`
			Community struct {
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
				Sex string `json:"sex,omitempty"`
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
			Links            struct {
				Self struct {
					Href string `json:"href"`
				} `json:"self"`
				Comments struct {
					Href string `json:"href"`
				} `json:"comments"`
				Likes struct {
					Href string `json:"href"`
				} `json:"likes"`
				Favorites struct {
					Href string `json:"href"`
				} `json:"favorites"`
			} `json:"_links"`
			UpdatedAt time.Time `json:"updated_at,omitempty"`
		} `json:"items"`
	} `json:"_embedded"`
}
