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
	prefs := strings.TrimPrefix(strings.TrimSpace(r.URL.Query().Get("prfe")), "https://www.startpage.com/do/mypage.pl?prfe=")
	if query == "" {
		address := "https://startpage.com"
		if prefs != "" {
			address = address + "?prfe=" + prefs
		}
		http.Redirect(w, r, address, http.StatusSeeOther)
		return
	}

	matched, err := regexp.Match(`![^\s]{1,}`, []byte(query))
	if err != nil {
		fmt.Fprint(w, html.EscapeString(err.Error()))
		return
	}

	if !matched {
		address := "https://www.startpage.com/sp/search?query=" + url.QueryEscape(query)
		if prefs != "" {
			address = address + "&prfe=" + prefs
		}
		http.Redirect(w, r, address, http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, "https://duckduckgo.com/?q="+url.QueryEscape(query), http.StatusSeeOther)
}

func main() {
	http.HandleFunc("/", root)

	http.ListenAndServe(":3000", nil)
}
