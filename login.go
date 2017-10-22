package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"

	"golang.org/x/net/html"

	"golang.org/x/net/publicsuffix"
)

var client http.Client

// Login effettua il login sul sito e ne salva la SID
func Login(conf *Configuration) {
	site := conf.LoginURL
	urlparse, _ := url.Parse(site)
	println("Chiamo la pagina per farmi assegnare un SID")
	option := cookiejar.Options{PublicSuffixList: publicsuffix.List}
	jar, _ := cookiejar.New(&option)
	client = http.Client{Jar: jar}
	client.Get(site)
	sid, _ := findSID(jar.Cookies(urlparse))
	fmt.Println(sid)

	resp, _ := client.PostForm(site, url.Values{
		"username": {conf.Username},
		"password": {conf.Password},
		//"redirect": {"./ucp.php?mode=login"},
		"sid": {sid},
		//"redirect": {"index.php"},
		"login": {"Login"},
	})
	println(resp.Status)

}

// GetEd2k retituisce tutti i link Ed2k nella pagina
func GetEd2k(link string) {
	resp, _ := client.Get(link)

	println(printTitle(resp.Body))

}

func findSID(cookie []*http.Cookie) (sid string, err error) {
	for _, c := range cookie {
		if c.Name == "phpbb3_ddu4final_sid" {
			return c.Value, nil
		}
	}
	err = errors.New("nessuno SID trovato")
	return
}

func printTitle(r io.Reader) string {
	z := html.NewTokenizer(r)

	for {
		tt := z.Next()
		switch tt {
		case html.StartTagToken, html.EndTagToken:
			tag, _ := z.TagName()
			if string(tag) == "title" {
				z.Next()
				return string(z.Text())
			}
		}
	}
}
