Goskirt
=======

Package goskirt provides Go-bindings for the excellent
[Upskirt](https://github.com/tanoku/upskirt)
Markdown parser ["that doesn't suck"](https://github.com/tanoku/upskirt).

To use goskirt, create a new Goskirt-value with the markdown extensions and
render modes you want. The extensions and render modes are encapsulated in the
struct as bitsets following the schematics of the Upskirt library. The created
value contains two methods: WriteHTML and WriteTOC that both parse the data
given in a byte slice and writes the formatted results into the given io.Writer
using the setup encapsuled in the underlaying struct type.

Example
-------

	import (
		"goskirt"
		"os"
	)

	func main() {
		data := []byte("Hello, upskirt!\n===============\n")

		skirt := goskirt.Goskirt{goskirt.EXT_AUTOLINK, 
			goskirt.HTML_SMARTYPANTS, 

		// <h1>Hello, upskirt!</h1>
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
	HTML_GITHUB_BLOCKCODE
	HTML_USE_XHTML
	HTML_SMARTYPANTS

... and the extensions respectively:

	EXT_NO_INTRA_EMPHASIS
	EXT_TABLES
	EXT_FENCED_CODE
	EXT_AUTOLINK
	EXT_STRIKETHROUGH
	EXT_LAX_HTML_BLOCKS

Install
-------

Goskirt bundles the Upskirt library, and hence has zero depedencies.
The most convenient way to install Goskirt is to use goinstall:

	goinstall http://github.com/madari/goskirt

License
-------

*For the Upskirt license, see the files in bundled in upskirt/.*

Copyright (c) 2011 Jukka-Pekka Kekkonen <karatepekka@gmail.com>

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
