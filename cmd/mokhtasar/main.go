package main

import (
	"database/sql"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/amirrezaask/mokhtasar/config"
	"github.com/amirrezaask/mokhtasar/handlers"
	"github.com/amirrezaask/mokhtasar/pkg"
	_ "github.com/lib/pq"
)

func randString(n int) string {
	letterRunes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func main() {
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s user=%s password=%s database=%s sslmode=%s",
		config.DatabaseHost,
		config.DatabaseUser,
		config.DatabasePass,
		config.DatabaseName,
		config.DatabaseSSLMode,
	))
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	mokhtasar := &pkg.Mokhtasar{
		DB:              db,
		RandomGenerator: randString,
	}
	handler := &handlers.HTTPHandler{Mokhtasar: mokhtasar}
	// mokhtasar.io/short?url=https://google.com
	http.HandleFunc("/short", handler.Shorten)
	//mokhtasar.io/long?key=harchi
	http.HandleFunc("/long", handler.Long)
	http.ListenAndServe(":8080", nil)
}
