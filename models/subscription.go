package models

import (
	"time"

	"github.com/fatih/structs"
)

type Subscription struct {
	ID                int       `json:"id" structs:"id" bson:"_id" db:"id"`
	StartTime         time.Time `json:"start_time" structs:"start_time" db:"start_time"`
	EndTime           time.Time `json:"end_time" structs:"end_time" db:"end_time"`
	SubscriptionPrice float64   `json:"subs_price" structs:"subs_price" db:"subs_price"`
	Status            bool      `json:"status" structs:"status" db:"status"`
	UserId            int       `json:"user_id" structs:"user_id" db:"user_id"`
	ItemName          string    `json:"item_name" structs:"item_name" db:"item_name"`
}

// Map converts structs to a map representation
func (s *Subscription) Map() map[string]interface{} {
	return structs.Map(s)
}

// Names returns the field names of Subscription model.
func (s *Subscription) Names() []string {
	fields := structs.Fields(s)
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
