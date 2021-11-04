package mongo

import (
	"fmt"
	"github.com/wasimkhan042/ustore-auth/db"
	"github.com/wasimkhan042/ustore-auth/models"
	"os"
	"reflect"
	"testing"
	"time"
)

// connection initiation with mongodb
func mongoConnection() (db.DataStore, error) {
	os.Setenv("DB_PORT", "27017")
	os.Setenv("DB_HOST", "ustore-auth-mongo-db")

	dbStore, err := NewMongoClient(db.Option{})
	return dbStore, err
}

// Testing SignUp for inserting user data
func Test_client_SignUp(t *testing.T) {

	type args struct {
		userInfo *models.User
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success - added user in db",
			args: args{userInfo: &models.User{
				FirstName:    "Wasim",
				MiddleName:   "",
				LastName:     "Khan",
				Email:        "wk@gmail.com",
				Username:     "wk",
				Password:     "password",
				ProfileImage: "image.png",
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, _ := mongoConnection()
			if err := c.DeleteUserByEmail(tt.args.userInfo.Email); err != nil {
				t.Error("Failed to delete user.")
			}

			err := c.SignUp(tt.args.userInfo)
			if (err != nil) != tt.wantErr {
				t.Errorf("SignUp() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// Testing SignIn for retrieving signed-up user data
func Test_client_SignIn(t *testing.T) {

	user := &models.User{
		Email: "wk@gmail.com",
	}

	type args struct {
		email string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "success-get user password by email",
			args:    args{email: user.Email},
			want:    "password",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, err := mongoConnection()
			if err != nil {
				fmt.Println("mongo database connection issue")
			}
			got, err := c.SignIn(tt.args.email)

			if (err != nil) != tt.wantErr {
				t.Errorf("SignIn() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != tt.want {
				t.Errorf("SignIn() got = %v, want %v", got, tt.want)
			}
		})
	}
}

// Testing GetProfile for retrieving signed-up user data
func Test_client_GetProfile(t *testing.T) {
	c, _ := mongoConnection()
	user := &models.User{
		Email:    "wk@gmail.com",
		Username: "wk",
	}

	type args struct {
		email string
	}
	tests := []struct {
		name    string
		args    args
		want    *models.User
		wantErr bool
	}{
		{
			name:    "success-get user by email",
			args:    args{email: user.Email},
			want:    user,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := c.GetProfile(tt.args.email)

			if (err != nil) != tt.wantErr {
				t.Errorf("GetProfile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got.Username, tt.want.Username) {
				t.Errorf("GetProfile() got = %v, want %v", got, tt.want)
			}
		})
	}
}

// Testing SubscribeItem for inserting subscription details
func Test_client_SubscribeItem(t *testing.T) {
	c, _ := mongoConnection()

	startTime := time.Now()
	endTime := time.Now().AddDate(0, 1, 0)
	user, _ := c.GetProfile("wk@gmail.com")

	type args struct {
		email            string
		subscriptionInfo *models.Subscription
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success - added subscribed item in db",
			args: args{subscriptionInfo: &models.Subscription{
				StartTime:         startTime,
				EndTime:           endTime,
				SubscriptionPrice: 1200,
				Status:            true,
				UserID:            user.ID,
				ItemName:          "TV",
			},
				email: "wk@gmail.com"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if err := c.SubscribeItem(tt.args.email, tt.args.subscriptionInfo); (err != nil) != tt.wantErr {
				t.Errorf("SubscribeItem() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// Testing ListSubscription for retrieving slice of given user subscription details
func Test_client_ListSubscription(t *testing.T) {
	c, _ := mongoConnection()
	user, _ := c.GetProfile("wk@gmail.com")

	subscription := []*models.Subscription{
		{
			UserID: user.ID,
		},
	}
	type args struct {
		email string
	}
	tests := []struct {
		name    string
		args    args
		want    []*models.Subscription
		wantErr bool
	}{
		{
			name:    "success-get subscription detail by email",
			args:    args{email: user.Email},
			want:    subscription,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := c.ListSubscription(tt.args.email)

			if (err != nil) != tt.wantErr {
				t.Errorf("ListSubscription() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got[0].UserID, tt.want[0].UserID) {
				t.Errorf("ListSubscription() got = %v, want %v", got, tt.want)
			}
		})
	}
}
