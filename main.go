package main

import (
	"log"
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/viniciusfeitosa/mongo/handlers"
	"gopkg.in/mgo.v2"
)

func main() {
	cmaHandler := handlers.NewCmaHandler(getMongoSession())
	mux := pat.New()
	mux.Get("/ping/cma/:identifier", http.HandlerFunc(cmaHandler.Get))
	mux.Get("/ping/cma/all/:identifier", http.HandlerFunc(cmaHandler.FindAll))
	mux.Post("/ping/cma/create", http.HandlerFunc(cmaHandler.Create))
	mux.Del("/ping/cma/remove/:identifier", http.HandlerFunc(cmaHandler.Delete))

	log.Println("App start on port: 5000")
	http.ListenAndServe(":5000", mux)
}

func getMongoSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		log.Println(err)
	}
	return s
}
