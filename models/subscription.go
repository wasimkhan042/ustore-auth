package models

import (
	"time"

	"github.com/fatih/structs"
)

// Subscription model.
type Subscription struct {
	ID                string    `json:"id" structs:"id" bson:"_id" db:"id"`
	StartTime         time.Time `json:"startTime" structs:"start_time" bson:"start_time" db:"start_time"`
	EndTime           time.Time `json:"endTime" structs:"end_time" bson:"end_time" db:"end_time"`
	SubscriptionPrice float64   `json:"subsPrice" structs:"subs_price" bson:"subs_price" db:"subs_price"`
	Status            bool      `json:"status" structs:"status" bson:"status" db:"status"`
	UserID            string    `json:"userId" structs:"user_id" bson:"user_id" db:"user_id"`
	ItemName          string    `json:"itemName" structs:"item_name" bson:"item_name" db:"item_name"`
}

// Map converts structs to a map representation.
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
