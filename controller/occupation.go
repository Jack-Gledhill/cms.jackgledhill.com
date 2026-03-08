package controller

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Jack-Gledhill/cms.jackgledhill.com/dto"
	"github.com/Jack-Gledhill/cms.jackgledhill.com/middleware"
	"github.com/Jack-Gledhill/cms.jackgledhill.com/service"

	"github.com/gorilla/mux"
)

type Occupation struct {
	Service service.Occupation
}

func (c *Occupation) RegisterRoutes(r *mux.Router) {
	pub := r.PathPrefix("/occupation").Subrouter()
	pub.Use(middleware.EnsureContentType)
	pub.HandleFunc("/", c.GetAll).Methods("GET")
	pub.HandleFunc("/{id:[0-9]+}", c.GetByID).Methods("GET")

	priv := r.PathPrefix("/occupation").Subrouter()
	priv.Use(middleware.EnsureContentType)
	priv.Use(middleware.EnsureAuthentication)
	priv.HandleFunc("/", c.Create).Methods("POST")
	priv.HandleFunc("/{id:[0-9]+}", c.Update).Methods("POST")
	priv.HandleFunc("/{id:[0-9]+}", c.Delete).Methods("DELETE")
}

func (c *Occupation) Create(w http.ResponseWriter, r *http.Request) {
	var body dto.OccupationCreate
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// TODO: validate entity
	_, err = c.Service.Create(context.Background(), body.ToEntity())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		// TODO: log error
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (c *Occupation) GetByID(w http.ResponseWriter, r *http.Request) {}

func (c *Occupation) GetAll(w http.ResponseWriter, _ *http.Request) {
	es, err := c.Service.GetAll(context.Background())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		// TODO: log error
		return
	}

	body := &dto.OccupationList{}
	body.FromEntities(es)
	res, err := json.Marshal(body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		// TODO: log error
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(res)
	if err != nil {
		// TODO: log error
	}
}

func (c *Occupation) Update(w http.ResponseWriter, r *http.Request) {}

func (c *Occupation) Delete(w http.ResponseWriter, r *http.Request) {}
