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
			log.Printf("redirecting to \"%v\"\n", sh.targetUrl.String())
			http.Redirect(w, req, sh.targetUrl.String(), http.StatusTemporaryRedirect)
		} else {
			log.Printf("specified target not found, using default config")
			log.Printf("redirecting to \"%v\"\n", DefaultConfig.DefaultTarget)
			http.Redirect(w, req, DefaultConfig.DefaultTarget, http.StatusTemporaryRedirect)
		}
	} else {
		panic("no target requested!")
	}
}

func list(w http.ResponseWriter, req *http.Request) {
	log.Printf("listing all shorts")
	var htmlcontent string = "<html><head>"
	htmlcontent += "<style>" + css + "</style>"
	htmlcontent += "</head><body><table>"
	htmlcontent += "<thead><tr><th>Slug</th><th>Target</th><th>Description</th></tr></thead><tbody>"
	for _, short := range AllShorts.Shorts {
		htmlcontent += "<tr><td><a href=\"" + short.targetUrl.String() + "\">" + short.Slug + "</a></td>" + "<td>" + short.Target + "</td>" + "<td>" + short.Description + "</td></tr>"
	}
	htmlcontent += "</tbody></table></body></html>"
	fmt.Fprint(w, htmlcontent)
}
