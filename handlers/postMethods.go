package handlers

import (
	"net/http"

	"github.com/supercoast/profile-api/data"
)

func (p *Profile) CreateProfile(rw http.ResponseWriter, r *http.Request) {
	np, err := data.FromJSON(r.Body)
	defer r.Body.Close()
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
	}

	p.ProfileDB.AddProfile(np)

	rw.WriteHeader(http.StatusCreated)
}
