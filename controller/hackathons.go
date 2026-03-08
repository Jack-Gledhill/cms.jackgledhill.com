package controller

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Jack-Gledhill/crud.jackgledhill.com/dto"
	"github.com/Jack-Gledhill/crud.jackgledhill.com/middleware"
	"github.com/Jack-Gledhill/crud.jackgledhill.com/service"

	"github.com/gorilla/mux"
)

type Hackathon struct {
	Service service.Hackathon
}

func (c *Hackathon) RegisterRoutes(r *mux.Router) {
	pub := r.PathPrefix("/hackathons").Subrouter()
	pub.Use(middleware.EnsureContentType)
	pub.HandleFunc("/", c.GetAll).Methods("GET")
	pub.HandleFunc("/{id:[0-9]+}", c.GetByID).Methods("GET")

	priv := r.PathPrefix("/hackathons").Subrouter()
	priv.Use(middleware.EnsureContentType)
	priv.Use(middleware.EnsureAuthentication)
	priv.HandleFunc("/", c.Create).Methods("POST")
	priv.HandleFunc("/{id:[0-9]+}", c.Update).Methods("POST")
	priv.HandleFunc("/{id:[0-9]+}", c.Delete).Methods("DELETE")
}

func (c *Hackathon) Create(w http.ResponseWriter, r *http.Request) {
	var body dto.HackathonCreate
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

func (c *Hackathon) GetByID(w http.ResponseWriter, r *http.Request) {}

func (c *Hackathon) GetAll(w http.ResponseWriter, _ *http.Request) {
	es, err := c.Service.GetAll(context.Background())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		// TODO: log error
		return
	}

	body := &dto.HackathonList{}
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

func (c *Hackathon) Update(w http.ResponseWriter, r *http.Request) {}

func (c *Hackathon) Delete(w http.ResponseWriter, r *http.Request) {}
