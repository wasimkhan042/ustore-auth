package service

import (
	"fmt"
	wraperrors "github.com/pkg/errors"
	"github.com/wasimkhan042/ustore-auth/authentication"
	"github.com/wasimkhan042/ustore-auth/models"
	"golang.org/x/crypto/bcrypt"
)

// Login to generate bearer-token.
func (c *Service) Login(userInfo *models.User) (string, error) {

	password, err := c.db.SignIn(userInfo.Email)
	if err != nil {
		return "", wraperrors.Wrap(err, "User not found")
	}

	// decrypt the hashed-password and compare
	err = bcrypt.CompareHashAndPassword([]byte(userInfo.Password), []byte(password))
	if err != nil {
		return "", wraperrors.Wrap(err, "wrong password")
	}
	token, err := authentication.GenerateJWT(userInfo.Email)
	if err != nil {
		fmt.Println("error defining token")
		return "", wraperrors.Wrap(err, "failed defining token")
	}
	return token, nil
}
