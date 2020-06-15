package handlers

import (
	"net/http"

	"github.com/supercoast/profile-api/data"
)

func (p *Profile) GetProfile(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(http.StatusNotImplemented)

}

func (p *Profile) ListProfiles(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	products := data.GetAllProfiles()

	err := data.ToJSON(products, rw)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
	}
}
