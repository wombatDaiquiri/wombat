package hejto

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/wombatDaiquiri/wombat/lajko"
	"github.com/wombatDaiquiri/wombat/ww"
)

const parserVersion = "hejto:0.0.1"

// Source implements scraper.Source for hejto.pl.
type Source struct {
	httpDoer ww.HTTPDoer
}

func (Source) Version() string {
	return parserVersion
}

func (s Source) Scrape(ctx context.Context, uri *url.URL) (lajko.Content, []*url.URL, error) {
	switch {
	case s.isPostListURI(ctx, uri):
		return s.scrapePostList(ctx, uri)
	case s.isPostURI(ctx, uri):
		return s.scrapePost(ctx, uri)
	case s.isCommentURI(ctx, uri):
		return s.scrapeComment(ctx, uri)
	default:
		return lajko.Content{}, nil, fmt.Errorf("unrecognized hejto URI format: %v", uri.String())
	}
}

func (Source) isPostListURI(ctx context.Context, uri *url.URL) bool {
	return uri.Path == "/posts"
}

func (s Source) scrapePostList(ctx context.Context, uri *url.URL) (lajko.Content, []*url.URL, error) {
	req, err := http.NewRequest("GET", uri.String(), nil)
	if err != nil {
		return lajko.Content{}, nil, err
	}
	req = req.WithContext(ctx)

	var response postListResponse
	resp, err := s.httpDoer.Do(req)
	if err != nil {
		return lajko.Content{}, nil, err
	}
	defer ww.DiscardHTTPResponse(resp)

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return lajko.Content{}, nil, err
	}

	// post list references all listed posts
	referencedURIs := make([]*url.URL, 0, len(response.Embedded.Items))
	for _, item := range response.Embedded.Items {
		referencedURIs = append(referencedURIs, s.relativeURI(item.Links.Self.Href))
	}

	// post list also references next page
	nextPageURI := s.relativeURI("/posts")
	query := url.Values{}
	query.Set("page", strconv.Itoa(response.Page+1))
	nextPageURI.RawQuery = query.Encode()
	referencedURIs = append(referencedURIs, nextPageURI)

	return lajko.Content{}, referencedURIs, nil
}

func (Source) isPostURI(ctx context.Context, uri *url.URL) bool {
	splitPath := strings.Split(uri.Path, "/")
	if len(splitPath) == 3 && splitPath[0] == "" && splitPath[1] == "posts" {
		return true
	}
	return false
}

func (s Source) scrapePost(ctx context.Context, uri *url.URL) (lajko.Content, []*url.URL, error) {
	// request uri and parse response
	req, err := http.NewRequest("GET", uri.String(), nil)
	if err != nil {
		return lajko.Content{}, nil, err
	}
	req = req.WithContext(ctx)

	var response postResponse
	resp, err := s.httpDoer.Do(req)
	if err != nil {
		return lajko.Content{}, nil, err
	}
	defer ww.DiscardHTTPResponse(resp)

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return lajko.Content{}, nil, err
	}

	// post is referenced by all comments
	var commentURIs []*url.URL
	if response.NumComments > 0 {
		commentURIs, err = s.CommentsURIs(ctx, uri.JoinPath("comments"))
		if err != nil {
			// TODO: proper logging
			fmt.Printf("error getting comments URIs: %v", err)
		}
	}

	return lajko.Content{
		URI: uri,
		Author: lajko.Actor{
			URI:           s.relativeURI(response.Author.Links.Self.Href),
			DisplayedName: response.Author.Username,
			AvatarURI:     response.Author.Avatar.Urls.X250,
		},
		InResponseToURI: nil,
		ParserVersion:   s.Version(),
		LajkoHTML:       response.LajkoHTML(),
		Attachments:     response.LajkoAttachments(),
	}, commentURIs, nil
}

func (s Source) CommentsURIs(ctx context.Context, commentListURI *url.URL) ([]*url.URL, error) {
	req, err := http.NewRequest("GET", commentListURI.String(), nil)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)

	var response commentsListResponse
	resp, err := s.httpDoer.Do(req)
	if err != nil {
		return nil, err
	}
	defer ww.DiscardHTTPResponse(resp)

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	// comment list references all listed comments
	referencedURIs := make([]*url.URL, 0, len(response.Embedded.Items))
	for _, comment := range response.Embedded.Items {
		referencedURIs = append(referencedURIs, s.relativeURI(comment.Links.Self.Href))
		for _, reply := range comment.Replies {
			referencedURIs = append(referencedURIs, s.relativeURI(reply.Links.Self.Href))
		}
	}

	return referencedURIs, nil
}

func (Source) isCommentURI(ctx context.Context, uri *url.URL) bool {
	splitPath := strings.Split(uri.Path, "/")
	if len(splitPath) == 5 && splitPath[0] == "" && splitPath[1] == "posts" && splitPath[3] == "comments" {
		return true
	}
	return false
}

func (Source) scrapeComment(ctx context.Context, uri *url.URL) (lajko.Content, []*url.URL, error) {
	return lajko.Content{}, nil, nil
}

// relativeURI returns a URL with the same hostname as the Source, but with the provided relativePath.
func (Source) relativeURI(relativePath string) *url.URL {
	baseURL := ww.Must(url.Parse("https://api.hejto.pl"))
	baseURL.Path = relativePath
	return baseURL
}
