package handlers

import "github.com/hashicorp/go-hclog"

type Profile struct {
	Logger *hclog.Logger
}

func NewProfile(l *hclog.Logger) *Profile {
	return &Profile{
		Logger: l,
	}
}
