package service

import (
	"github.com/wasimkhan042/ustore-auth/models"
	"golang.org/x/crypto/bcrypt"
)

// Registration to call signup func at db layer for user registration.
func (c *Service) Registration(userInfo *models.User) error {
	password := []byte(userInfo.Password)

	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	userInfo.Password = string(hashedPassword)
	return c.db.SignUp(userInfo)
}
