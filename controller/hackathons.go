package controller

import (
	"net/http"

	"github.com/Jack-Gledhill/cms.jackgledhill.com/middleware"
	"github.com/Jack-Gledhill/cms.jackgledhill.com/service"

	"github.com/gorilla/mux"
)

type HackathonController struct {
	Service service.HackathonService
}

func (c *HackathonController) RegisterRoutes(r *mux.Router) {
	pub := r.PathPrefix("/hackathons").Subrouter()
	pub.Use(middleware.EnsureContentType)
	pub.HandleFunc("/", c.GetAll).Methods("GET")

	priv := r.PathPrefix("/hackathons").Subrouter()
	priv.Use(middleware.EnsureAuthentication)
	priv.HandleFunc("/", c.Create).Methods("POST")
}

func (c *HackathonController) Create(w http.ResponseWriter, r *http.Request) {

}

func (c *HackathonController) GetAll(w http.ResponseWriter, r *http.Request) {

}
