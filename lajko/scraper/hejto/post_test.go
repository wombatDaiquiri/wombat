package hejto

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/wombatDaiquiri/wombat/wt"
)

func Test_postResponse_LajkoContent(t *testing.T) {

}

func Test_postResponse_LajkoHTML(t *testing.T) {
	t.Parallel()
	// given
	// content from https://api.hejto.pl/posts/no-w-sumie-parsowanie-to-nauka-html-css-i-js-tylko-odwrotnie-proba-mikrofonu-gdy
	fd, err := os.Open("fixtures/post.json")
	wt.ExpectNoError(t, err, "error opening fixture")
	defer fd.Close()

	var resp postResponse
	err = json.NewDecoder(fd).Decode(&resp)
	wt.ExpectNoError(t, err, "error decoding fixture into response")

	// when
	lajkoHTML := resp.LajkoHTML()

	// then
	expectedLajkoHTMLBytes, err := os.ReadFile("fixtures/posts_lajko_html.html")
	wt.ExpectNoError(t, err, "error reading expected lajkoHTML")
	wt.ExpectEqual(t, string(expectedLajkoHTMLBytes), lajkoHTML)
}

func Test_postResponse_Attachments(t *testing.T) {
	t.Parallel()
	// given
	// content from https://api.hejto.pl/posts/no-w-sumie-parsowanie-to-nauka-html-css-i-js-tylko-odwrotnie-proba-mikrofonu-gdy
	fd, err := os.Open("fixtures/post.json")
	wt.ExpectNoError(t, err, "error opening fixture")
	defer fd.Close()

	var resp postResponse
	err = json.NewDecoder(fd).Decode(&resp)
	wt.ExpectNoError(t, err, "error decoding fixture into response")

	// when
	lajkoHTML := resp.LajkoHTML()

	// then
	expectedLajkoHTMLBytes, err := os.ReadFile("fixtures/posts_lajko_html.html")
	wt.ExpectNoError(t, err, "error reading expected lajkoHTML")
	wt.ExpectEqual(t, string(expectedLajkoHTMLBytes), lajkoHTML)
}
