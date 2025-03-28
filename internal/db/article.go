package db

import (
	"bytes"
	"os"

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

	data, err := os.ReadFile(a.Markdown)
	if err != nil {
		return "", err
	}

	err = md.Convert(data, &buf)

	return buf.String(), err
}
