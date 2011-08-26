package goskirt

/*
#cgo CFLAGS: -O3
#include "sundown/markdown.h"
#include "sundown/buffer.h"
#include "sundown/html.h"
*/
import "C"

import (
	"io"
	"os"
	"unsafe"
)

const (
	Version = C.UPSKIRT_VERSION
)

// Markdown extensions
const (
	EXT_NO_INTRA_EMPHASIS = C.MKDEXT_NO_INTRA_EMPHASIS // 1 << 0
	EXT_TABLES            = C.MKDEXT_TABLES
	EXT_FENCED_CODE       = C.MKDEXT_FENCED_CODE
	EXT_AUTOLINK          = C.MKDEXT_AUTOLINK
	EXT_STRIKETHROUGH     = C.MKDEXT_STRIKETHROUGH
	EXT_LAX_HTML_BLOCKS   = C.MKDEXT_LAX_HTML_BLOCKS
	EXT_SPACE_HEADERS     = C.MKDEXT_SPACE_HEADERS
	EXT_SUPERSCRIPT       = C.MKDEXT_SUPERSCRIPT // 1 << 7
)

// Render modes
const (
	HTML_SKIP_HTML        = C.HTML_SKIP_HTML // 1 << 0
	HTML_SKIP_STYLE       = C.HTML_SKIP_STYLE
	HTML_SKIP_IMAGES      = C.HTML_SKIP_IMAGES
	HTML_SKIP_LINKS       = C.HTML_SKIP_LINKS
	HTML_EXPAND_TABS      = C.HTML_EXPAND_TABS
	HTML_SAFELINK         = C.HTML_SAFELINK
	HTML_TOC              = C.HTML_TOC
	HTML_HARD_WRAP        = C.HTML_HARD_WRAP
	HTML_GITHUB_BLOCKCODE = C.HTML_GITHUB_BLOCKCODE
	HTML_USE_XHTML        = C.HTML_USE_XHTML    // 1 << 11
	HTML_SMARTYPANTS      = HTML_USE_XHTML << 1 // 1 << 12
)

// Renderers
const (
	HTMLRenderer = iota
	TOCRenderer
)

// An Goskirt represents a combination of rendering modes and enabled markdown
// extensions.
type Goskirt struct {
	Extensions, RenderModes uint
}

// WriteHTML renders HTML into w using the source markdown-data p.
// It returns the number of bytes written or and error if writing failed.
func (g Goskirt) WriteHTML(w io.Writer, p []byte) (n int, err os.Error) {
	return render(w, g.Extensions, g.RenderModes, HTMLRenderer, p)
}

// WriteTOC renders a list table of contents into w using the source markdown-data p
// It returns the number of bytes written or and error if writing failed.
func (g Goskirt) WriteTOC(w io.Writer, p []byte) (n int, err os.Error) {
	return render(w, g.Extensions, g.RenderModes, TOCRenderer, p)
}

func render(w io.Writer, extensions, renderModes, rndr uint, p []byte) (n int, err os.Error) {
	var ob *C.struct_buf
	var ib C.struct_buf
	var callbacks C.struct_sd_callbacks
	var options C.struct_html_renderopt

	ib.data = (*C.char)(unsafe.Pointer(&p[0]))
	ib.size = C.size_t(len(p))
	ib.asize = ib.size
	ib.unit = C.size_t(0)

	ob = C.bufnew(128)
	C.bufgrow(ob, C.size_t(float64(ib.size)*1.2))

	switch rndr {
	case HTMLRenderer:
		C.sdhtml_renderer(&callbacks, &options, C.uint(renderModes&^HTML_SMARTYPANTS))
	case TOCRenderer:
		C.sdhtml_toc_renderer(&callbacks, &options)
	default:
		panic("unknown renderer")
	}

	C.sd_markdown(ob, &ib, C.uint(extensions), &callbacks, unsafe.Pointer(&options))

	if renderModes&HTML_SMARTYPANTS > 0 {
		sb := C.bufnew(128)
		C.sdhtml_smartypants(sb, ob)
		n, err = w.Write((*[1 << 30]byte)(unsafe.Pointer(sb.data))[0:sb.size])
		C.bufrelease(sb)
	} else {
		n, err = w.Write((*[1 << 30]byte)(unsafe.Pointer(ob.data))[0:ob.size])
	}

	C.bufrelease(ob)
	return
}
