package main

import (
	"fmt"
	"html"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

func root(w http.ResponseWriter, r *http.Request) {
	query := strings.TrimSpace(r.URL.Query().Get("query"))
	var suffix string
	if r.URL.RawQuery != "" {
		suffix = "?" + r.URL.RawQuery
	}

	if query == "" {
		http.Redirect(w, r, "https://startpage.com"+suffix, http.StatusSeeOther)
		return
	}

	matched, err := regexp.Match(`![^\s]{1,}`, []byte(query))
	if err != nil {
		fmt.Fprint(w, html.EscapeString(err.Error()))
		return
	}

	if !matched {
		http.Redirect(w, r, "https://www.startpage.com/sp/search"+suffix, http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, "https://duckduckgo.com/?q="+url.QueryEscape(query), http.StatusSeeOther)
}

func main() {
	http.HandleFunc("/", root)

	http.ListenAndServe(":3000", nil)
}
