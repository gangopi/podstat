package main
hello world

import (
	"log"
	"net/http"
	"net/url"
	"os"
)

func redirectAndStore(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "http://devopsdeflope.ru/"+r.URL.Path[1:], http.StatusFound)

	values := make(url.Values)
	ga_cookie := "0"
	xff_ip := r.Header.Get("X-Forwarded-For")

	for i := range r.Cookies() {
		if r.Cookies()[i].Name == "_ga" {
			ga_cookie = r.Cookies()[i].Value
		}
	}

	values.Set("v", "1")
	values.Set("tid", "MO-41332661-1")
	values.Set("cid", ga_cookie)
	values.Set("t", "event")
	values.Set("ea", "Download")
	values.Set("el", r.URL.Path)
	values.Set("ec", "Podcast")
	if len(xff_ip) != 0 {
		values.Set("utmip", xff_ip)
	}

	http.PostForm("http://www.google-analytics.com/collect", url.Values(values))
}

func rdrToDeflope(w http.ResponseWriter, r *http.Request) {

	http.Redirect(w, r, "http://devopsdeflope.ru/", http.StatusFound)
}

func main() {
	http.HandleFunc("/mp3/", redirectAndStore)
	http.HandleFunc("/", rdrToDeflope)
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
	os.Exit(0)
}
