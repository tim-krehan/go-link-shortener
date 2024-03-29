package main

import (
	_ "embed"
	"fmt"
	"log"
	"net/http"
	"strings"
)

//go:embed table.css
var css string

func redirect(w http.ResponseWriter, req *http.Request) {
	if len(strings.Split(req.URL.Path, "/")) >= 2 {
		target := strings.Split(req.URL.Path, "/")[2]
		log.Printf("requested target is \"%v\"\n", target)
		sh := AllShorts.GetShort(target)
		if sh != nil {
			fmt.Fprintf(w, "%v\n", sh.targetUrl.String())
		} else {
			log.Printf("specified target not found, using default config")
			fmt.Fprintf(w, "%v\n", DefaultConfig.DefaultTarget)
		}
	} else {
		panic("no target requested!")
	}
}

func list(w http.ResponseWriter, req *http.Request) {
	var htmlcontent string = "<html><head>"
	htmlcontent += "<style>" + css + "</style>"
	htmlcontent += "</head><body><table>"
	htmlcontent += "<thead><tr><th>Slug</th><th>Target</th></tr></thead><tbody>"
	for _, short := range AllShorts.Shorts {
		htmlcontent += "<tr><td><a href=\"" + short.targetUrl.String() + "\">" + short.Slug + "</a></td>" + "<td>" + short.Target + "</td></tr>"
	}
	htmlcontent += "</tbody></table></body></html>"
	fmt.Fprint(w, htmlcontent)
}
