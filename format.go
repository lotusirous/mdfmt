package mdfmt

import (
	"io"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/text"
)

// Format writes reformatted markdown source.
func Format(source []byte, w io.Writer, opts ...parser.ParseOption) error {
	md := goldmark.New(
		goldmark.WithExtensions(
			extension.GFM,
			extension.DefinitionList,
			extension.Footnote,
		),
		goldmark.WithParserOptions(
			parser.WithAttribute(),
		),
	)
	doc := md.Parser().Parse(
		text.NewReader(source), opts...)
	return Render(w, source, doc)
}

var Markdown renderer.Renderer = new(mdRender)

type mdRender struct{}

// AddOptions adds given option to this renderer.
func (md *mdRender) AddOptions(opts ...renderer.Option) {}

// Write render node as Markdown.
func (md *mdRender) Render(w io.Writer, source []byte, node ast.Node) (err error) {
	return Render(w, source, node)
}
