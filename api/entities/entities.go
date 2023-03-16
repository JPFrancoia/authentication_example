package entities

import (
	"time"

	"github.com/google/uuid"
	"github.com/markbates/goth"
)

type User struct {
	UserId       uuid.UUID `db:"user_id" json:"user_id"`
	CreationTime time.Time `db:"creation_time" json:"creation_time"`
	Provider     string    `db:"provider" json:"provider"`
	Email        string    `db:"email" json:"email"`
}

// Alias the non-local goth.User type so that I can implement the ToUser()
// method for it
type GothUser goth.User

func (gothUser GothUser) ToUser() User {
	return User{
		UserId:       uuid.New(),
		CreationTime: time.Now(),
		Provider:     gothUser.Provider,
		Email:        gothUser.Email,
	}
}
