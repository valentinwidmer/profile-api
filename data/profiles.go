package data

import "time"

type Profiles []*Profile

var ProfilesStore = Profiles{
	&Profile{
		Email:     "valentin.wali@gmail.com",
		FirstName: "Valentin",
		LastName:  "Widmer",
		CreatedAt: time.Now(),
	},
	&Profile{
		Email:     "francesca.ferago@gmail.com",
		FirstName: "Francesca",
		LastName:  "Ferago",
		CreatedAt: time.Now(),
	},
}

type Profile struct {
	Email      string `json:"Email"`
	FirstName  string `json:"Firstname"`
	LastName   string `json:"Lastname"`
	CreatedAt  time.Time
	ModifiedAt time.Time
}

func GetAllProfiles() Profiles {
	return ProfilesStore
}
