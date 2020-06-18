package handlers

import (
	"net/http"

	"github.com/supercoast/profile-api/data"
)

func (p *Profile) CreateProfile(rw http.ResponseWriter, r *http.Request) {
	profile := r.Context().Value(KeyProduct{}).(*data.Profile)

	p.Logger.Info("Adding profile to database", "Email", profile.Email)
	p.ProfileDB.AddProfile(profile)
}
