package lajko

import (
	"net/url"
)

type Content struct {
	// URI is a unique identifier for this content from its source website.
	URI             *url.URL
	Author          Actor
	InResponseToURI *url.URL

	// ParserVersion is version of parser used to render LajkoHTML from original article.
	ParserVersion string
	// LajkoHTML is HTML representation of the content, rendered by using
	LajkoHTML string

	// Attachments is a list of non-text content that are a part of this content.
	Attachments []Attachment
}

type Attachment struct {
	// URI is a unique identifier for this attachment from its source website.
	URI *url.URL
	// ContentType is the MIME type of the attachment that can be found under URI.
	ContentType string
	// MetaData is a JSON with free-form metadata about the attachment.
	MetaData string
}

type Actor struct {
	// URI is a unique identifier for this author from its source website.
	URI *url.URL
	// DisplayedName is the name that should be displayed for this author.
	DisplayedName string
	// AvatarURI is a URI to the avatar image of this author.
	AvatarURI string
}

type Like struct {
	WhoURI  *url.URL
	WhatURI *url.URL
}
