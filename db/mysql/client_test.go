package mysql

import (
	"fmt"
	"os"
	"reflect"
	"time"

	//	"reflect"
	"testing"
	//	"time"

	"github.com/wasimkhan042/ustore-auth/db"
	"github.com/wasimkhan042/ustore-auth/models"
)

// The tests are written in a sequence, testing all functions of mysql/client file.
// These test functions are dependent on each other because the functions
// are testing single entry.

// connection initiation with mysql
func mysqlConnection() (db.DataStore, error) {
	os.Setenv("DB_PORT", "3306")
	os.Setenv("DB_HOST", "ustore-auth-mysql-db")
	os.Setenv("DB_USER", "root")

	dbStore, err := NewMysqlClient(db.Option{})
	return dbStore, err
}

// Testing register function with the given data
func Test_client_SignUp(t *testing.T) {
	type args struct {
		user *models.User
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success - added user in db",
			args: args{user: &models.User{
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
			c, err := mysqlConnection()
			if err != nil {
				panic(fmt.Sprintf("errrrorrr %v", err))
			}
			if tt.args.user != nil {
				fmt.Println("hello")
				fmt.Println(tt.args.user)
				err = c.SignUp(tt.args.user)
				fmt.Println("hello")
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("SignUp() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// Testing on the data inserted in Test_client_Register function.
func Test_client_SignIn(t *testing.T) {
	c, _ := mysqlConnection()

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
			got, err := c.SignIn(tt.args.email)

			if (err != nil) != tt.wantErr {
				t.Errorf("SignIn() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SignIn() got = %v, want %v", got, tt.want)
			}
		})
	}
}

// Testing on the data inserted in Test_client_Register function.
func Test_client_GetProfile(t *testing.T) {
	c, _ := mysqlConnection()
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

// Testing SubscribeItem for dummy data
func Test_client_SubscribeItem(t *testing.T) {
	c, _ := mysqlConnection()

	start_time := time.Now()
	end_time := time.Now().AddDate(0, 1, 0)
	user, _ := c.GetProfile("wk@gmail.com")
	type args struct {
		email        string
		subscription *models.Subscription
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success - added subscribed item in db",
			args: args{subscription: &models.Subscription{
				StartTime:         start_time,
				EndTime:           end_time,
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
			if err := c.SubscribeItem(tt.args.email, tt.args.subscription); (err != nil) != tt.wantErr {
				t.Errorf("SubscribeItem() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// Testing ListSubscription for inserted subscribedItem data
func Test_client_ListSubscription(t *testing.T) {
	c, _ := mysqlConnection()
	user, _ := c.GetProfile("wk@gmail.com")

	subscription := []*models.Subscription{
		{ItemName: "TV",
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
