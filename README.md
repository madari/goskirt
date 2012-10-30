Goskirt
=======

Package goskirt provides Go-bindings for the excellent
[Sundown](https://github.com/vmg/sundown) Markdown parser. (F/K/A Upskirt).

To use goskirt, create a new Goskirt-value with the markdown extensions and
render modes you want. The extensions and render modes are encapsulated in the
struct as bitsets following the schematics of the Sundown library. The created
value contains two methods: WriteHTML and WriteTOC that both parse the data
given in a byte slice and writes the formatted results into the given io.Writer
using the setup encapsuled in the underlaying struct type.

Example
-------

	package main
	
	import (
		"github.com/madari/goskirt"
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

Extensions and render modes
---------------------------

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

Install
-------

	go get github.com/madari/goskirt

License
-------

*For the Sundown license, see the bundled C files.*

Copyright (c) 2012 Jukka-Pekka Kekkonen <karatepekka@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining
a copy of this software and associated documentation files (the
"Software"), to deal in the Software without restriction, including
without limitation the rights to use, copy, modify, merge, publish,
distribute, sublicense, and/or sell copies of the Software, and to
permit persons to whom the Software is furnished to do so, subject to
the following conditions:

The above copyright notice and this permission notice shall be
included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
