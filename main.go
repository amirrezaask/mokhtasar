package main

import (
	"math/rand"
	"net/http"
	"time"
)

var db = make(map[string]string)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randString(n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func getOriginalURL(key string) string {
	return db[key]
}

func shorten(url string) string {
	key := randString(5)
	db[key] = url
	return key
}

func main() {
	// mokhtasar.io/short?url=https://google.com
	http.HandleFunc("/short", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Only GET request supported"))
			return
		}
		url := r.URL.Query().Get("url")
		if url == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("need a url to short"))
			return
		}
		key := shorten(url)
		toClickURL := "localhost:8080/long?key=" + key
		w.Write([]byte(toClickURL))
	})
	//mokhtasar.io/long?key=harchi
	http.HandleFunc("/long", func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Query().Get("key")
		if key == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("need a url to short"))
			return
		}
		url := getOriginalURL(key)
		w.Write([]byte(url))
	})
	http.ListenAndServe(":8080", nil)
}
