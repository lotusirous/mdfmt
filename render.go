package mdfmt

import (
	"io"

	"github.com/yuin/goldmark/ast"
)

// Render write node as Markdown o writer.
func Render(w io.Writer, source []byte, node ast.Node) (err error) {
	return nil
}
