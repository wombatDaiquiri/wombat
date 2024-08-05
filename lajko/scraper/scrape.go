package scraper

import (
	"context"
	"net/url"
	"time"

	"github.com/wombatDaiquiri/wombat/lajko"
	"github.com/wombatDaiquiri/wombat/ww"
)

// SourceStub is a stub for Source interface, useful in testing.
type SourceStub struct {
	VersionFunc func() string
	ScrapeFunc  func(ctx context.Context, uri *url.URL) (Result, error)
}

func (ss SourceStub) Version() string {
	return ss.VersionFunc()
}

func (ss SourceStub) ScrapeNewContent(ctx context.Context, uri *url.URL) (Result, error) {
	return ss.ScrapeFunc(ctx, uri)
}

type Source interface {
	Version() string
	ScrapeNewContent(ctx context.Context, uri *url.URL) (Result, error)
}

type Storage interface {
	ContentAlreadyStored(ctx context.Context, uri *url.URL) (bool, error)
	UpsertContent(ctx context.Context, article lajko.Content) error
}

type Request struct {
	URI string
}

type Result struct {
	Content      lajko.Content
	ReferencedBy []*url.URL
	References   []*url.URL
}

// ScrapedURIs returns initial URIs that should be scraped and monitored for new content.
func ScrapedURIs() []*url.URL {
	return []*url.URL{
		ww.Must(url.Parse("https://api.hejto.pl/posts?page=1")),
	}
}

// ScrapeNewContent scrapes new content appearing on source uris and content they reference.
func ScrapeNewContent(ctx context.Context, storage Storage, selector SourceSelector) error {
	initialURIs := ScrapedURIs()
	uriChan := make(chan *url.URL, 5*len(initialURIs))
	for _, uri := range initialURIs {
		uriChan <- uri
	}

	homepageScrapeTicker := time.NewTicker(time.Minute)
	go func() {
		for {
			select {
			case <-homepageScrapeTicker.C:
				ww.P("scraping homepage uris")
				for _, uri := range ScrapedURIs() {
					// TODO: write w/ timeout?
					uriChan <- uri
				}
			case <-ctx.Done():
				ww.P("context done, stopping scraping")
				homepageScrapeTicker.Stop()
				close(uriChan)
			}
		}
	}()

	for uri := range uriChan {
		// check if content is already stored
		alreadyStored, err := storage.ContentAlreadyStored(ctx, uri)
		if err != nil {
			return err
		}
		if alreadyStored {
			continue
		}

		// scrape
		source, err := selector.SelectScrapingSource(ctx, uri)
		if err != nil {
			return err
		}
		res, err := source.ScrapeNewContent(ctx, uri)
		if err != nil {
			return err
		}
		// store
		err = storage.UpsertContent(ctx, res.Content)
		if err != nil {
			return err
		}

		// add content references to scrape queue
		// TODO: need to unblock a pool of goroutines here to not get stuck on the processing OR run this inside a goroutine
		for _, referencedByURI := range res.ReferencedBy {
			uriChan <- referencedByURI
		}
		for _, referencedURI := range res.References {
			uriChan <- referencedURI
		}
	}
	return nil
}
