package models

import (
	"github.com/fatih/structs"
)

type User struct {
	ID           string `json:"id" structs:"id" bson:"_id" db:"id"`
	FirstName    string `json:"first_name" structs:"first_name"bson:"first_name" db:"first_name"`
	MiddleName   string `json:"middle_name" structs:"middle_name" bson:"middle_name" db:"middle_name"`
	LastName     string `json:"last_name" structs:"last_name" bson:"last_name" db:"last_name"`
	Email        string `json:"email" structs:"email" bson:"email" db:"email"`
	Username     string `json:"username" structs:"username" bson:"username" db:"username"`
	Password     string `json:"password" structs:"password" bson:"password" db:"password"`
	ProfileImage string `json:"profile_image" structs:"profile_image" bson:"profile_image" db:"profile_image"`
}

// Map converts structs to a map representation
func (u *User) Map() map[string]interface{} {
	return structs.Map(u)
}

// Names returns the field names of User model.
func (u *User) Names() []string {
	fields := structs.Fields(u)
	names := make([]string, len(fields))

	for i, field := range fields {
		name := field.Name()
		tagName := field.Tag(structs.DefaultTagName)
		if tagName != "" {
			name = tagName
		}
		names[i] = name
	}

	return names
}
