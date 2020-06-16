package data

import (
	"fmt"
	"time"

	"github.com/hashicorp/go-hclog"
)

type ProfileStore interface {
	AddProfile(p *Profile)
	DeleteProfile(p string)
	UpdateProfile(p *Profile)
	GetProfile(pi string) (*Profile, error)
}

type Profiles []*Profile

type ProfileDB struct {
	Logger hclog.Logger
	DB     Profiles
}

func NewProfileDB(l hclog.Logger) *ProfileDB {
	return &ProfileDB{
		Logger: l,
		DB:     Profiles{},
	}
}

type Profile struct {
	Email      string `json:"Email"`
	FirstName  string `json:"Firstname"`
	LastName   string `json:"Lastname"`
	CreatedAt  time.Time
	ModifiedAt time.Time
}

func (p *ProfileDB) GetAllProfiles() Profiles {
	return p.DB
}

func (p *ProfileDB) GetProfile(pi string) (*Profile, error) {
	for _, profile := range p.DB {
		if pi == profile.Email {
			return profile, nil
		}
	}
	return nil, fmt.Errorf("Product not found")
}

func (p *ProfileDB) AddProfile(np *Profile) {
	p.DB = append(p.DB, np)
}
