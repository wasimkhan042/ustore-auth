package service

import (
	wraperrors "github.com/pkg/errors"
	"github.com/wasimkhan042/ustore-auth/authentication"
	"github.com/wasimkhan042/ustore-auth/models"
)

// UserProfile to generate bearer-token.
func (c *Service) UserProfile(bearerHeader string) (*models.User, error) {
	email, err := authentication.ValidateHeader(bearerHeader)
	if err != nil {
		return nil, wraperrors.New("Error validating user authentication")
	}
	userInfo, err := c.db.GetProfile(email.(string)) //p.serviceInfoHandler.ViewProfile(p.dbClient, email.(string))
	if err != nil {
		return nil, wraperrors.New("Error fetching profile details")
	}

	return userInfo, nil
}
