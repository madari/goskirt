/*
Package goskirt provides Go-bindings for the excellent Sundown markdown parser
(F/K/A Upskirt).

To use goskirt, create a new Goskirt-value with the markdown extensions and
render modes you want. The extensions and render modes are encapsulated in the
struct as bitsets following the schematics of the Sundown library. The created
value contains two methods: WriteHTML and WriteTOC that both parse the data
given in a byte slice and writes the formatted results into the given io.Writer
using the setup encapsuled in the underlaying struct type.

Usage:

	package main

	import (
		"goskirt"
		"os"
	)

	func main() {
		data := []byte("Hello, sundown!\n===============\n")

		skirt := goskirt.Goskirt{
			goskirt.EXT_AUTOLINK | goskirt.EXT_STRIKETHROUGH,
			goskirt.HTML_SMARTYPANTS | goskirt.HTML_USE_XHTML,
		}

		// <h1>Hello, sundown!</h1>
		skirt.WriteHTML(os.Stdout, data)
	}

Render mode is a combination of:

	HTML_SKIP_HTML
	HTML_SKIP_STYLE
	HTML_SKIP_IMAGES
	HTML_SKIP_LINKS
	HTML_EXPAND_TABS
	HTML_SAFELINK
	HTML_TOC
	HTML_HARD_WRAP
	HTML_USE_XHTML
	HTML_ESCAPE
	HTML_SMARTYPANTS

... and the extensions respectively:

	EXT_NO_INTRA_EMPHASIS
	EXT_TABLES
	EXT_FENCED_CODE
	EXT_AUTOLINK
	EXT_STRIKETHROUGH
	EXT_SPACE_HEADERS
	EXT_SUPERSCRIPT
	EXT_LAX_SPACING

See LICENSE file for licensing details.
*/
package goskirt
