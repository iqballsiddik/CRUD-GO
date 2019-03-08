package model

import "time"

// Profile struct
type Profile struct {
	ID        string
	FristName string
	LastName  string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Profile type Profiles List
type Profiles []Profile

// NewProfile Profile Constructor
func NewProfile() *Profile {
	return &Profile{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
