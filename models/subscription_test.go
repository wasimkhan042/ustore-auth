package models

import (
	"reflect"
	"testing"
	"time"
)

// SubscriptionTime is used as a fixed time
var (
	subscriptionTime = time.Now()
)

func TestSubscription_Map(t *testing.T) {
	type fields struct {
		ID                string
		StartTime         time.Time
		EndTime           time.Time
		SubscriptionPrice float64
		Status            bool
		UserID            string
		ItemName          string
	}
	tests := []struct {
		name   string
		fields fields
		want   map[string]interface{}
	}{
		{
			name: "subscription",
			fields: fields{
				ID:                "123",
				StartTime:         subscriptionTime,
				EndTime:           subscriptionTime,
				SubscriptionPrice: 1200.00,
				Status:            true,
				UserID:            "1234",
				ItemName:          "TV",
			},
			want: map[string]interface{}{
				"id":         "123",
				"start_time": subscriptionTime,
				"end_time":   subscriptionTime,
				"subs_price": 1200.00,
				"status":     true,
				"user_id":    "1234",
				"item_name":  "TV",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Subscription{
				ID:                tt.fields.ID,
				StartTime:         tt.fields.StartTime,
				EndTime:           tt.fields.EndTime,
				SubscriptionPrice: tt.fields.SubscriptionPrice,
				Status:            tt.fields.Status,
				UserID:            tt.fields.UserID,
				ItemName:          tt.fields.ItemName,
			}
			if got := s.Map(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubscription_Names(t *testing.T) {
	type fields struct {
		ID                string
		StartTime         time.Time
		EndTime           time.Time
		SubscriptionPrice float64
		Status            bool
		UserID            string
		ItemName          string
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name: "subscription",
			fields: fields{
				ID:                "123",
				StartTime:         subscriptionTime,
				EndTime:           subscriptionTime,
				SubscriptionPrice: 1200.00,
				Status:            true,
				UserID:            "1234",
				ItemName:          "TV",
			},
			want: []string{
				"id",
				"start_time",
				"end_time",
				"subs_price",
				"status",
				"user_id",
				"item_name",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Subscription{
				ID:                tt.fields.ID,
				StartTime:         tt.fields.StartTime,
				EndTime:           tt.fields.EndTime,
				SubscriptionPrice: tt.fields.SubscriptionPrice,
				Status:            tt.fields.Status,
				UserID:            tt.fields.UserID,
				ItemName:          tt.fields.ItemName,
			}
			if got := s.Names(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Names() = %v, want %v", got, tt.want)
			}
		})
	}
}
