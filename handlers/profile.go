package handlers

import (
	"github.com/go-playground/validator/v10"
	"github.com/hashicorp/go-hclog"
	"github.com/supercoast/profile-api/data"
)

type Profile struct {
	Logger    hclog.Logger
	Validator *data.Validation
	ProfileDB *data.ProfileDB
}

func NewProfile(l hclog.Logger, v *validator.Validate, pd *data.ProfileDB) *Profile {
	return &Profile{
		Logger:    l,
		Validator: v,
		ProfileDB: pd,
	}
}
