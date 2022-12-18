package main

import (
	"fmt"
	"html"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

var ssp_prefs string = "53308301f70432e8195ac16843b38ea5c0ab0e03080d61edd06d725a311fc3c0597636e9064b5df843962b8fd795febd252657d9b0cd79587c886ba0e5e953e3254d1068bc2e010e1ce4d2293e"

func root(w http.ResponseWriter, r *http.Request) {
	query := strings.TrimSpace(r.URL.Query().Get("q"))
	if query == "" {
		http.Redirect(w, r, "https://startpage.com?prfe="+ssp_prefs, http.StatusSeeOther)
		return
	}

	matched, err := regexp.Match(`![^\s]{1,}`, []byte(query))
	if err != nil {
		fmt.Fprint(w, html.EscapeString(err.Error()))
		return
	}

	if !matched {
		// Open search in new tab disabled, safe search disabled and server region set to EU servers
		http.Redirect(w, r, "https://www.startpage.com/sp/search?query="+url.QueryEscape(query)+"&prfe="+ssp_prefs, http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, "https://duckduckgo.com/?q="+url.QueryEscape(query), http.StatusSeeOther)
}

func main() {
	http.HandleFunc("/", root)

	http.ListenAndServe(":3000", nil)
}
