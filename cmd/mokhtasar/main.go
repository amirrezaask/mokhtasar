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
	"github.com/spf13/cobra"
	"go.uber.org/zap"
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

var rootCmd = &cobra.Command{
	Use: "",
}
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: `Serves an HTTP server`,
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
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
		dc := zap.NewDevelopmentConfig()
		dc.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
		logger, err := dc.Build()
		if err != nil {
			panic(err)
		}
		sl := logger.Sugar()
		handler := &handlers.HTTPHandler{Mokhtasar: mokhtasar, Logger: sl}
		// mokhtasar.io/short?url=https://google.com
		http.HandleFunc("/short", handler.Shorten)
		//mokhtasar.io/long?key=harchi
		http.HandleFunc("/long", handler.Long)
		http.ListenAndServe(":8080", nil)

	},
}
var shortCmd = &cobra.Command{
	Use:   "short",
	Short: `Shorts the given url`,
	Long:  "Shorts given url",
	Run: func(cmd *cobra.Command, args []string) {
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
		if len(args) < 1 {
			panic("you need to pass a url")
		}
		key, err := mokhtasar.Shorten(args[0])
		if err != nil {
			panic(err)
		}
		fmt.Println(key)
	},
}
var longCmd = &cobra.Command{
	Use:   "long",
	Short: `expands given key`,
	Long:  "expands given key",
	Run: func(cmd *cobra.Command, args []string) {
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
		if len(args) < 1 {
			panic("you need to pass a key")
		}
		url, err := mokhtasar.GetOriginalURL(args[0])
		if err != nil {
			panic(err)
		}
		fmt.Println(url)
	},
}

func main() {
	rootCmd.AddCommand(shortCmd, longCmd, serveCmd)
	err := rootCmd.Execute()
	if err != nil {
		panic(err)
	}
}
