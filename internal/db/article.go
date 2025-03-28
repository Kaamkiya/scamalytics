package db

import (
	"bytes"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
)

var md = goldmark.New(
	goldmark.WithExtensions(extension.GFM),
	goldmark.WithParserOptions(parser.WithAutoHeadingID()),
)

type Article struct {
	Title    string `json:"title"`
	Markdown string `json:"markdown"`
}

func (a Article) RenderHTML() (string, error) {
	var buf bytes.Buffer
	err := md.Convert([]byte(a.Markdown), &buf)

	return buf.String(), err
}
