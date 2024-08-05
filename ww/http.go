package ww

import (
	"io"
	"net/http"
)

// HTTPDoer is an interface for making HTTP requests suitable for mocking.
type HTTPDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

func DiscardHTTPResponse(resp *http.Response) {
	_, _ = io.Copy(io.Discard, resp.Body)
	_ = resp.Body.Close()
}
