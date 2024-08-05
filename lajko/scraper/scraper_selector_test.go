package scraper

import (
	"context"
	"net/url"
	"testing"

	"github.com/wombatDaiquiri/wombat/wt"
	"github.com/wombatDaiquiri/wombat/ww"
)

func Test_SourceSelector_SelectScrapingSource_NoMatchingSource(t *testing.T) {
	t.Parallel()
	// given
	ctx := context.Background()
	mapping := map[string]Source{}
	requestURI := ww.Must(url.Parse("https://api.hejto.pl/posts"))
	sourceSelector := NewSourceSelector(mapping)
	// when
	source, err := sourceSelector.SelectScrapingSource(ctx, requestURI)
	// then
	wt.ExpectError(t, err, "error selecting source")
	wt.ExpectNil(t, source, "source should be nil")
}

func Test_SourceSelector_SelectScrapingSource_MatchingSourceExits(t *testing.T) {
	t.Parallel()
	// given
	ctx := context.Background()
	mapping := map[string]Source{
		"api.hejto.pl": SourceStub{
			VersionFunc: func() string { return "matching host" },
		},
	}
	requestURI := ww.Must(url.Parse("https://api.hejto.pl/posts"))
	sourceSelector := NewSourceSelector(mapping)
	// when
	source, err := sourceSelector.SelectScrapingSource(ctx, requestURI)
	// then
	wt.ExpectNoError(t, err, "error selecting source")
	wt.ExpectEqual(t, "matching host", source.Version(), "invalid source returned")
}
