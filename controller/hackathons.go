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

func (c *Hackathon) Create(w http.ResponseWriter, r *http.Request) {}

func (c *Hackathon) GetByID(w http.ResponseWriter, r *http.Request) {}

func (c *Hackathon) GetAll(w http.ResponseWriter, _ *http.Request) {
	hs, err := c.Service.GetAll(context.Background())
	if err != nil {
		// TODO: handle error
	}

	body := &dto.HackathonList{}
	body.FromEntities(hs)
	res, err := json.Marshal(body)
	if err != nil {
		// TODO: handle error
	}

	_, err = w.Write(res)
	if err != nil {
		// TODO: handle error
	}
}

func (c *Hackathon) Update(w http.ResponseWriter, r *http.Request) {}

func (c *Hackathon) Delete(w http.ResponseWriter, r *http.Request) {}
