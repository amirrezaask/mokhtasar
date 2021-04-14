package main

import (
	"database/sql"
	"math/rand"
	"net/http"
	"time"

	_ "github.com/lib/pq"
)

var DB *sql.DB

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randString(n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func getOriginalURL(key string) (string, error) {
	query := `SELECT url FROM urls WHERE key=$1`
	rows, err := DB.Query(query, key)
	if err != nil {
		return "", err
	}
	var url string
	for rows.Next() {
		err = rows.Scan(&url)
		if err != nil {
			return "", err
		}
	}
	return url, nil
}

func shorten(url string) (string, error) {
	key := randString(5)
	_, err := DB.Exec(`INSERT INTO urls (url, key) VALUES ($1, $2)`, url, key)
	if err != nil {
		return "", err
	}
	return key, nil
}

func main() {
	db, err := sql.Open("postgres", "host=127.0.0.1 user=postgres password=admin database=mokhtasar sslmode=disable")
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	DB = db
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
		key, err := shorten(url)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("error in adding new url"))
			return
		}
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
		url, err := getOriginalURL(key)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("error in finding key you gave us"))
			return
		}
		w.Write([]byte(url))
	})
	http.ListenAndServe(":8080", nil)
}
