package hejto

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/wombatDaiquiri/wombat/lajko/scraper"
	"github.com/wombatDaiquiri/wombat/ww"
)

func Test_ArticleURIs(t *testing.T) {
	s := Source{
		httpDoer: &http.Client{},
	}
	_, referencedURIs, err := s.Scrape(context.Background(), scraper.ScrapedURIs()[0])
	if err != nil {
		panic(err)
	}
	ww.PJSON(referencedURIs)
	panic(fmt.Sprintf("%+#v", referencedURIs))
}
