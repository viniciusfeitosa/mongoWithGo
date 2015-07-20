package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/viniciusfeitosa/mongo/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// CmaHandler is the struct of handler to cma
type CmaHandler struct {
	mongo *mgo.Session
}

// NewCmaHandler generate an instance of the CmaHandler
func NewCmaHandler(mongo *mgo.Session) *CmaHandler {
	return &CmaHandler{mongo: mongo}
}

// Create is a method handler of cma to create a new register in mongo
func (c *CmaHandler) Create(w http.ResponseWriter, r *http.Request) {
	cma := models.CmaJSON{}
	if err := json.NewDecoder(r.Body).Decode(&cma); err != nil {
		log.Println(err)
	}

	cma.ID = bson.NewObjectId()
	if err := c.mongo.DB("pingo").C("cma").Insert(cma); err != nil {
		log.Println(err)
	}

	cmaj, err := json.Marshal(cma)
	if err != nil {
		log.Println(err)
	}
	log.Println("Salvo com sucesso")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(cmaj)
}

// Get the rigister of cma in mongodb
func (c *CmaHandler) Get(w http.ResponseWriter, r *http.Request) {
	identifier := r.URL.Query().Get(":identifier")
	cma := models.CmaJSON{}
	if err := c.mongo.DB("pingo").C("cma").Find(bson.M{"identifier": identifier}).One(&cma); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	cmaj, err := json.Marshal(cma)
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(cmaj)
}

// FindAll all rigisters of cma in mongodb
func (c *CmaHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	identifier := r.URL.Query().Get(":identifier")
	cma := []models.CmaJSON{}
	if err := c.mongo.DB("pingo").C("cma").Find(bson.M{"identifier": identifier}).All(&cma); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	cmaj, err := json.Marshal(cma)
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(cmaj)
}

// Delete remove a register of the cma in mongodb
func (c *CmaHandler) Delete(w http.ResponseWriter, r *http.Request) {
	identifier := r.URL.Query().Get(":identifier")
	if err := c.mongo.DB("pingo").C("cma").Remove(bson.M{"identifier": identifier}); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
