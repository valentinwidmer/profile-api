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
	Email      string    `json:"Email" validation:"required,email" `
	FirstName  string    `json:"Firstname" validation:"required,name"`
	LastName   string    `json:"Lastname" validation:"required,name"`
	CreatedOn  time.Time `json:"-"`
	ModifiedOn time.Time `json:"-"`
	DeletedOn  time.Time `json:"-"`
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
