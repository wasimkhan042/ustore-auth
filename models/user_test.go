package models

import (
	"reflect"
	"testing"
)

func TestUser_Map(t *testing.T) {
	type fields struct {
		ID           string
		FirstName    string
		MiddleName   string
		LastName     string
		Email        string
		Username     string
		Password     string
		ProfileImage string
	}
	tests := []struct {
		name   string
		fields fields
		want   map[string]interface{}
	}{
		{
			name: "success-TestUserMap",
			fields: fields{
				ID:           "1234",
				FirstName:    "wasim",
				MiddleName:   "",
				LastName:     "khan",
				Email:        "simsim@gmail.com",
				Password:     "password",
				Username:     "simsim",
				ProfileImage: "default.png",
			},
			want: map[string]interface{}{
				"id":            "1234",
				"first_name":    "wasim",
				"middle_name":   "",
				"last_name":     "khan",
				"email":         "simsim@gmail.com",
				"password":      "password",
				"username":      "simsim",
				"profile_image": "default.png",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				ID:           tt.fields.ID,
				FirstName:    tt.fields.FirstName,
				MiddleName:   tt.fields.MiddleName,
				LastName:     tt.fields.LastName,
				Email:        tt.fields.Email,
				Username:     tt.fields.Username,
				Password:     tt.fields.Password,
				ProfileImage: tt.fields.ProfileImage,
			}
			if got := u.Map(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_Names(t *testing.T) {
	type fields struct {
		ID           string
		FirstName    string
		MiddleName   string
		LastName     string
		Email        string
		Username     string
		Password     string
		ProfileImage string
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name: "success-TestUserName",
			fields: fields{
				ID:           "1234",
				FirstName:    "wasim",
				MiddleName:   "",
				LastName:     "khan",
				Email:        "simsim@gmail.com",
				Password:     "password",
				Username:     "simsim",
				ProfileImage: "default.png",
			},
			want: []string{"id", "first_name", "middle_name", "last_name", "email", "username", "password", "profile_image"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				ID:           tt.fields.ID,
				FirstName:    tt.fields.FirstName,
				MiddleName:   tt.fields.MiddleName,
				LastName:     tt.fields.LastName,
				Email:        tt.fields.Email,
				Username:     tt.fields.Username,
				Password:     tt.fields.Password,
				ProfileImage: tt.fields.ProfileImage,
			}
			if got := u.Names(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Names() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Failure testCase
func TestUser_Map1(t *testing.T) {
	type fields struct {
		ID           string
		FirstName    string
		MiddleName   string
		LastName     string
		Email        string
		Username     string
		Password     string
		ProfileImage string
	}
	tests := []struct {
		name   string
		fields fields
		want   map[string]interface{}
	}{
		{
			name: "failed-TestUserMap",
			fields: fields{
				ID:           "12364",
				FirstName:    "wasim",
				MiddleName:   "",
				LastName:     "khan",
				Email:        "simsim@gmail.com",
				Password:     "password",
				Username:     "simsim",
				ProfileImage: "default.png",
			},
			want: map[string]interface{}{
				"id":            "1234",
				"first_name":    "wasim",
				"middle_name":   "",
				"last_name":     "khan",
				"email":         "simsim@gmail.com",
				"password":      "password",
				"username":      "simsim",
				"profile_image": "default.png",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				ID:           tt.fields.ID,
				FirstName:    tt.fields.FirstName,
				MiddleName:   tt.fields.MiddleName,
				LastName:     tt.fields.LastName,
				Email:        tt.fields.Email,
				Username:     tt.fields.Username,
				Password:     tt.fields.Password,
				ProfileImage: tt.fields.ProfileImage,
			}
			if got := u.Map(); reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Failure testCase
func TestUser_Names1(t *testing.T) {
	type fields struct {
		ID           string
		FirstName    string
		MiddleName   string
		LastName     string
		Email        string
		Username     string
		Password     string
		ProfileImage string
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name: "failed-TestUserNames",
			fields: fields{
				ID:           "1234",
				FirstName:    "wasim",
				MiddleName:   "",
				LastName:     "khan",
				Email:        "simsim@gmail.com",
				Password:     "password",
				Username:     "simsim",
				ProfileImage: "default.png",
			},
			want: []string{"first_name", "middle_name", "last_name", "email", "username", "password", "profile_image"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				ID:           tt.fields.ID,
				FirstName:    tt.fields.FirstName,
				MiddleName:   tt.fields.MiddleName,
				LastName:     tt.fields.LastName,
				Email:        tt.fields.Email,
				Username:     tt.fields.Username,
				Password:     tt.fields.Password,
				ProfileImage: tt.fields.ProfileImage,
			}
			if got := u.Names(); reflect.DeepEqual(got, tt.want) {
				t.Errorf("Names() = %v, want %v", got, tt.want)
			}
		})
	}
}
