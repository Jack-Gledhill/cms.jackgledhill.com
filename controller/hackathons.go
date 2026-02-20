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

	priv := r.PathPrefix("/hackathons").Subrouter()
	priv.Use(middleware.EnsureAuthentication)
	priv.HandleFunc("/", c.Create).Methods("POST")
}

func (c *Hackathon) Create(w http.ResponseWriter, r *http.Request) {
	// TODO
}

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
