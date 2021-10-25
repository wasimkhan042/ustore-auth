package mongo

import (
	"context"
	"fmt"

	wraperrors "github.com/pkg/errors"
	guuid "github.com/satori/go.uuid"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github/wasimkhan042/ustore-auth/config"
	"github/wasimkhan042/ustore-auth/db"
	"github/wasimkhan042/ustore-auth/models"
)

const (
	userCollection         = "user"
	subscriptionCollection = "subscription"
)

func init() {
	db.Register("mongo", NewMongoClient)
}

type client struct {
	conn *mongo.Client
}

// NewMongoClient initializes a mongo database connection.
func NewMongoClient(conf db.Option) (db.DataStore, error) {
	uri := fmt.Sprintf("mongodb://%s:%s", viper.GetString(config.DbHost), viper.GetString(config.DbPort))
	log().Infof("initializing mongodb: %s", uri)

	cli, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, wraperrors.Wrap(err, "failed to connect")
	}

	return &client{conn: cli}, nil
}

// SignUp stores new user details in 'user' collection.
func (m client) SignUp(userInfo *models.User) error {
	if _, err := m.GetProfile(userInfo.Email); err != nil {
		userInfo.ID = guuid.NewV4().String()
		collection := m.conn.Database(viper.GetString(config.DbName)).Collection(userCollection)

		if _, err := collection.InsertOne(context.TODO(), userInfo); err != nil {
			return wraperrors.Wrap(err, "failed to add user")
		}

		return nil
	}

	return wraperrors.New("User already exists.")
}

// SignIn return password from db on successful call else error.
func (m client) SignIn(email string) (string, error) {
	var user *models.User
	collection := m.conn.Database(viper.GetString(config.DbName)).Collection(userCollection)

	if err := collection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&user); err != nil {
		if mongo.ErrNoDocuments != nil {
			return "", wraperrors.Wrap(err, "failed to fetch user....not found")
		}

		return "", nil
	}

	return user.Password, nil
}

// GetProfile return user details from db on successful call else error.
func (m client) GetProfile(email string) (*models.User, error) {
	var user *models.User
	collection := m.conn.Database(viper.GetString(config.DbName)).Collection(userCollection)

	if err := collection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&user); err != nil {
		if mongo.ErrNoDocuments != nil {
			return nil, wraperrors.Wrap(err, "failed to fetch user....not found")
		}

		return user, nil
	}

	return user, nil
}

// SubscribeItem add user subscribed item details in the 'subscription' collection.
func (m client) SubscribeItem(email string, subscriptionInfo *models.Subscription) error {
	subscriptionInfo.ID = guuid.NewV4().String()

	user, err := m.GetProfile(email)
	if err != nil {
		return wraperrors.Wrap(err, "failed to retrieve user details")
	}

	subscriptionInfo.UserID = user.ID
	collection := m.conn.Database(viper.GetString(config.DbName)).Collection(subscriptionCollection)

	if _, err := collection.InsertOne(context.TODO(), subscriptionInfo); err != nil {
		return wraperrors.Wrap(err, "failed to add subscription details")
	}

	return nil
}

// ListSubscription return slice of given user subscriptions on successful call.
func (m client) ListSubscription(email string) ([]*models.Subscription, error) {
	var subscription []*models.Subscription

	user, err := m.GetProfile(email)
	if err != nil {
		return nil, wraperrors.Wrap(err, "failed to add subscription details")
	}

	collection := m.conn.Database(viper.GetString(config.DbName)).Collection(subscriptionCollection)

	cursor, err := collection.Find(context.TODO(), bson.M{"user_id": user.ID})
	if err != nil {
		return nil, wraperrors.Wrap(err, "failed to retrieve subscription documents")
	}

	if err = cursor.All(context.TODO(), &subscription); err != nil {
		return nil, wraperrors.Wrap(err, "failed to scan retrieved subscription documents")
	}

	return subscription, nil
}
