package main

import (
	"flag"
	"fmt"
	"github.com/madari/goskirt"
	"io/ioutil"
	"os"
)

var eflags = map[string]uint{
	"EXT_NO_INTRA_EMPHASIS": goskirt.EXT_NO_INTRA_EMPHASIS,
	"EXT_TABLES":            goskirt.EXT_TABLES,
	"EXT_FENCED_CODE":       goskirt.EXT_FENCED_CODE,
	"EXT_AUTOLINK":          goskirt.EXT_AUTOLINK,
	"EXT_STRIKETHROUGH":     goskirt.EXT_STRIKETHROUGH,
	"EXT_SPACE_HEADERS":     goskirt.EXT_SPACE_HEADERS,
	"EXT_SUPERSCRIPT":       goskirt.EXT_SUPERSCRIPT,
	"EXT_LAX_SPACING":       goskirt.EXT_LAX_SPACING,
}

var rmflags = map[string]uint{
	"HTML_SKIP_HTML":   goskirt.HTML_SKIP_HTML,
	"HTML_SKIP_STYLE":  goskirt.HTML_SKIP_STYLE,
	"HTML_SKIP_IMAGES": goskirt.HTML_SKIP_IMAGES,
	"HTML_SKIP_LINKS":  goskirt.HTML_SKIP_LINKS,
	"HTML_EXPAND_TABS": goskirt.HTML_EXPAND_TABS,
	"HTML_SAFELINK":    goskirt.HTML_SAFELINK,
	"HTML_TOC":         goskirt.HTML_TOC,
	"HTML_HARD_WRAP":   goskirt.HTML_HARD_WRAP,
	"HTML_USE_XHTML":   goskirt.HTML_USE_XHTML,
	"HTML_ESCAPE":      goskirt.HTML_ESCAPE,
	"HTML_SMARTYPANTS": goskirt.HTML_SMARTYPANTS,
}

func parseFlags() (extensions, renderModes, renderer uint) {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, `Usage: %s -render=toc|html [OPTION]...
Reads markdown-formatted data from STDIN and renders it to STDOUT.

Available options are:
`, os.Args[0])
		flag.PrintDefaults()
	}

	for k, _ := range eflags {
		flag.Bool(k, false, k)
	}
	for k, _ := range rmflags {
		flag.Bool(k, false, k)
	}

	r := flag.String("render", "", "What to render (either html or toc)")

	flag.Parse()
	flag.Visit(func(f *flag.Flag) {
		if v, ok := eflags[f.Name]; ok {
			extensions |= v
		} else if v, ok := rmflags[f.Name]; ok {
			renderModes |= v
		}
	})

	switch *r {
	case "html":
		renderer = goskirt.HTMLRenderer
	case "toc":
		renderer = goskirt.TOCRenderer
	default:
		flag.Usage()
		os.Exit(1)
	}

	return
}

func main() {
	extensions, renderModes, renderer := parseFlags()
	gs := goskirt.Goskirt{extensions, renderModes}

	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	switch renderer {
	case goskirt.HTMLRenderer:
		gs.WriteHTML(os.Stdout, data)
	default:
		gs.WriteTOC(os.Stdout, data)
	}
}
