package main

import (
	"fmt"
	"math/rand"
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
	shortenKey := shorten("http://google.com")
	fmt.Println(shortenKey)
	fmt.Println(getOriginalURL(shortenKey))
}
