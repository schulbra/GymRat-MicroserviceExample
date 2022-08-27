package main

import (
	_ "encoding/json"
	"fmt"

	_ "github.com/lib/pq"

	_ "github.com/go-chi/chi/v5"
	_ "github.com/go-chi/chi/v5/middleware"
	_ "github.com/go-pg/pg/v10"
	"log"
	"net/http"
	_ "net/url"
	_ "strconv"

	"cs361new/pkg/api"
	"cs361new/pkg/db"

	"os"
)

func main() {
	pgdb, err := db.NewDB()
	if err != nil {
		panic(err)
	}

	router := api.NewAPI(pgdb)
	//router := api.NewAPI()
	log.Print("we're up and running!")
	//log.Print("server is AOK")
	port := os.Getenv("PORT")
	err = http.ListenAndServe(fmt.Sprintf(":%s", port), router)
	if err != nil {
		log.Printf("error due to router: %v\n", err)
	}
	//router setup below:
	/*	r := chi.NewRouter()
		r.Use(middleware.Logger)
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("hello world"))
		})
		http.ListenAndServe(":3000", r)
		log.Print("we're up and running!")
	*/

	//pgdb, err := db.NewDB()

	//err = http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")))

	//err = http.ListenAndServe(fmt.Sprintf("PORT", router)
	//	if err != nil {
	//	log.Printf("error from router: %v\n", err, router)
}
