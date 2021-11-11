package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	runtime "github.com/wasimkhan042/ustore-auth"
	genmodel "github.com/wasimkhan042/ustore-auth/gen/models"
	"github.com/wasimkhan042/ustore-auth/gen/restapi/operations/login"
	"github.com/wasimkhan042/ustore-auth/models"
)

// NewSignUp handles request for saving student
func NewLogin(rt *runtime.Runtime) login.LoginHandler {
	return &signIn{rt: rt}
}

type signIn struct {
	rt *runtime.Runtime
}

func (u *signIn) Handle(params login.LoginParams) middleware.Responder {
	user := models.User{
		Email:    *params.Login.Email,
		Password: *params.Login.Password,
	}

	token, err := u.rt.Service().Login(&user)
	if err != nil {
		log().Errorf("failed to add user: %s", err)
		return login.NewLoginInternalServerError()
	}

	return login.NewLoginOK().WithPayload(&genmodel.LoginSuccess{Success: true, Token: token})
}
