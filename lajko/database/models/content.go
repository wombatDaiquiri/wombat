package models

type Content struct {
	// URI is a unique identifier for this content from its source website.
	URI string `db:"uri"`

	AuthorURI       string `db:"author_uri"`
	InResponseToURI string `db:"in_response_to_uri"`

	ParserVersion string `db:"parser_version"`
	LajkoHTML     string `db:"lajko_html"`

	Attachments []Attachment
}

type Attachment struct {
	// URI is a unique identifier for this attachment from its source website.
	URI string `json:"uri"`
	// ContentType is the MIME type of the attachment that can be found under URI.
	ContentType string `json:"content_type"`
	// MetaData is a JSON with free-form metadata about the attachment.
	MetaData string `json:"meta_data"`
}

type Actor struct {
	// URI is a unique identifier for this author from its source website.
	URI string `db:"uri"`
	// DisplayedName is the name that should be displayed for this author.
	DisplayedName string `db:"displayed_name"`
	// AvatarURI is a URI to the avatar image of this author.
	AvatarURI string `db:"avatar_uri"`
}

type Like struct {
	// Who is a reference to the Actor liking a piece of content.
	WhoURI string `db:"who_uri"`
	Who    *Actor
	// What is a reference to the Content that is being liked.
	WhatURI string `db:"what_uri"`
	What    *Content
}
