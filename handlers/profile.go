package handlers

import (
	"github.com/hashicorp/go-hclog"
	"github.com/supercoast/profile-api/data"
)

type Profile struct {
	Logger    hclog.Logger
	ProfileDB *data.ProfileDB
}

func NewProfile(l hclog.Logger, pd *data.ProfileDB) *Profile {
	return &Profile{
		Logger:    l,
		ProfileDB: pd,
	}
}
