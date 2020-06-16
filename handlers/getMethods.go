package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/supercoast/profile-api/data"
)

func (p *Profile) GetProfile(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-Type", "application/json")

	vars := mux.Vars(r)
	val := vars["email"]

	profile, err := p.ProfileDB.GetProfile(val)
	if err != nil {
		p.Logger.Error("Couldn't fetch product", "Error:", err)
	}

	err = data.ToJSON(profile, rw)
	if err != nil {
		p.Logger.Error("Couldn't write data to customer", "Error:", err)
	}
}

func (p *Profile) ListProfiles(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	products := p.ProfileDB.GetAllProfiles()

	err := data.ToJSON(products, rw)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
	}
}
