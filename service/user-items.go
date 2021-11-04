package service

import (
	wraperrors "github.com/pkg/errors"
	"github.com/wasimkhan042/ustore-auth/authentication"
	"github.com/wasimkhan042/ustore-auth/models"
)

// UserItemList to generate bearer-token.
func (c *Service) UserItemList(bearerHeader string) ([]*models.Subscription, error) {
	email, err := authentication.ValidateHeader(bearerHeader)
	if err != nil {
		return nil, wraperrors.New("Error validating user authentication")
	}
	items, err := c.db.ListSubscription(email.(string)) //p.serviceInfoHandler.ViewProfile(p.dbClient, email.(string))
	if err != nil {
		return nil, wraperrors.New("Error fetching profile details")
	}

	return items, nil
}
