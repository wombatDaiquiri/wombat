package scraper

import (
	"context"
	"fmt"
	"net/url"
)

// SourceSelector selects an appropriate Source for scraping provided URI.
//
// It selects the right Source based on a hostname to Source mapping passed to it on initialization.
type SourceSelector struct {
	hostToSource map[string]Source
}

func NewSourceSelector(hostToSource map[string]Source) SourceSelector {
	return SourceSelector{hostToSource: hostToSource}
}

func (ss SourceSelector) SelectScrapingSource(ctx context.Context, uri *url.URL) (Source, error) {
	source, found := ss.hostToSource[uri.Host]
	if !found {
		return nil, fmt.Errorf("no source for host %q", uri.Host)
	}
	return source, nil
}
