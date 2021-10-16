package db

import (
	"log"

	"github/wasimkhan042/uStore-Auth/models"
)

// DataStore is an interface for query ops.
type DataStore interface {
	SignUp(userInfo *models.User) error
	SignIn(email string) (string, error)
	GetProfile( email string) (*models.User, error)
	SubscribeItem(email string, subscriptionInfo *models.Subscription) error
	ListSubscription(email string) ([]*models.Subscription, error)
}

// Option holds configuration for data store clients.
type Option struct {
	TestMode bool
}

// DataStoreFactory holds configuration for data store.
type DataStoreFactory func(conf Option) (DataStore, error)

var datastoreFactories = make(map[string]DataStoreFactory)

// Register saves data store into a data store factory.
func Register(name string, factory DataStoreFactory) {
	if factory == nil {
		log.Fatalf("Datastore factory %s does not exist.", name)
		return
	}
	if _, ok := datastoreFactories[name]; ok {
		log.Fatalf("Datastore factory %s already registered. Ignoring.", name)
		return
	}
}
