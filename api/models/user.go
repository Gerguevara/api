package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        string    `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"-"`
	FirstName string    `gorm:"not null" json:"-" `
	LastName  string    `gorm:"not null" json:"last_name" `
	Email     string    `gorm:"not null;uniqueIndex" json:"email" `
	Tasks     []Task    `json:"tasks"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"auto_preupdate" json:"updated_at,omitempty"`
}

// Hooks
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {

	if u.FirstName == "Gerardo" {
		return errors.New("You Name is too cool")
	}
	return
}

// Marshal
// func (u *User) MarshalJSON() ([]byte, error) {

// 	type Alias User
// 	return json.Marshal(&struct {
// 		*Alias
// 	}{
// 		Alias: (*Alias)(u),
// 	})
// }
