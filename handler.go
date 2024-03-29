package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

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
