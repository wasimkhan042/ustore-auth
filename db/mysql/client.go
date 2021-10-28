package mysql

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	wraperrors "github.com/pkg/errors"
	guuid "github.com/satori/go.uuid"
	"github.com/spf13/viper"

	"github.com/wasimkhan042/ustore-auth/config"
	"github.com/wasimkhan042/ustore-auth/db"
	"github.com/wasimkhan042/ustore-auth/models"
)

type client struct {
	db *sqlx.DB
}

func init() {
	db.Register("mysql", NewMysqlClient)
}

func formatDSN() string {
	cfg := mysql.NewConfig()
	cfg.Addr = fmt.Sprintf("%s:%s", viper.GetString(config.DbHost), viper.GetString(config.DbPort))
	cfg.DBName = viper.GetString(config.DbName)
	cfg.ParseTime = true
	cfg.User = viper.GetString(config.DbUser)
	cfg.Passwd = viper.GetString(config.DbPass)
	return cfg.FormatDSN()
}

// NewMysqlClient initializes a mysql database connection.
func NewMysqlClient(conf db.Option) (db.DataStore, error) {
	log().Info("initializing mysql connection")

	cli, err := sqlx.Connect("mysql", formatDSN())
	if err != nil {
		return nil, wraperrors.Wrap(err, "failed to connect")
	}
	return &client{db: cli}, nil
}

// SignUp stores new user details in user table.
func (c *client) SignUp(user *models.User) error {
	user.ID = guuid.NewV4().String()
	names := user.Names()

	if _, err := c.db.NamedExec(fmt.Sprintf(`INSERT INTO user (%s) VALUES(%s)`, strings.Join(names, ","), strings.Join(mkPlaceHolder(names, ":", func(name, prefix string) string {
		return prefix + name
	}), ",")), user); err != nil {
		return wraperrors.Wrap(err, "failed to add user")
	}

	return nil
}

// SignIn return password on successful call else error.
func (c *client) SignIn(email string) (string, error) {
	var user models.User

	if err := c.db.Get(&user.Password, fmt.Sprintf(`SELECT password FROM user WHERE email='%s'`, email)); err != nil {
		if sql.ErrNoRows != nil {
			return "", wraperrors.Wrap(err, "failed to fetch user....not found")
		}

		return "", nil
	}

	return user.Password, nil
}

// GetProfile return user details from db on successful call else error.
func (c *client) GetProfile(email string) (*models.User, error) {
	var user models.User
	if err := c.db.Get(&user, fmt.Sprintf(`SELECT * FROM user WHERE email='%s'`, email)); err != nil {
		if sql.ErrNoRows != nil {
			return nil, wraperrors.Wrap(err, "failed to fetch user....not found")
		}

		return &user, nil
	}

	return &user, nil
}

// SubscribeItem add user subscribed item details on successful call else return error.
func (c *client) SubscribeItem(email string, subscription *models.Subscription) error {
	subscription.ID = guuid.NewV4().String()

	user, err := c.GetProfile(email)
	if err != nil {
		return wraperrors.Wrap(err, "failed to retrieve user profile details")
	}

	subscription.UserID = user.ID
	names := subscription.Names()

	if _, err := c.db.NamedExec(fmt.Sprintf(`INSERT INTO subscription (%s) VALUES(%s) `, strings.Join(names, ","), strings.Join(mkPlaceHolder(names, ":", func(name, prefix string) string {
		return prefix + name
	}), ",")), subscription); err != nil {
		return wraperrors.Wrap(err, "failed to add user")
	}

	return nil
}

// ListSubscription return slice of given user subscriptions on successful call.
func (c *client) ListSubscription(email string) ([]*models.Subscription, error) {
	var subscription []*models.Subscription

	user, err := c.GetProfile(email)
	if err != nil {
		return subscription, wraperrors.Wrap(err, "failed to retrieve user profile details")
	}

	if err := c.db.Select(&subscription, fmt.Sprintf(`SELECT * FROM subscription WHERE user_id='%s'`, user.ID)); err != nil {
		if sql.ErrNoRows != nil {
			return subscription, wraperrors.Wrap(err, "failed to fetch user....not found")
		}

		return subscription, nil
	}

	return subscription, nil
}

func mkPlaceHolder(names []string, prefix string, formatName func(name, prefix string) string) []string {
	ph := make([]string, len(names))

	for i, name := range names {
		ph[i] = formatName(name, prefix)
	}

	return ph
}
