package handlers

import (
	"github.com/go-openapi/runtime/middleware"

	runtime "github.com/wasimkhan042/ustore-auth"
	genmodel "github.com/wasimkhan042/ustore-auth/gen/models"
	"github.com/wasimkhan042/ustore-auth/gen/restapi/operations/signup"
	"github.com/wasimkhan042/ustore-auth/models"
)

// NewSignUp handles request for saving student
func NewSignUp(rt *runtime.Runtime) signup.SignupHandler {
	return &signUp{rt: rt}
}

type signUp struct {
	rt *runtime.Runtime
}

func (u *signUp) Handle(params signup.SignupParams) middleware.Responder {
	user := models.User{
		FirstName:    *params.Signup.FirstName,
		MiddleName:   params.Signup.MiddleName,
		LastName:     params.Signup.LastName,
		Email:        *params.Signup.Email,
		Username:     *params.Signup.Username,
		Password:     *params.Signup.Password,
		ProfileImage: params.Signup.ProfileImage,
	}
	if err := u.rt.Service().Registration(&user); err != nil {
		log().Errorf("failed to add user: %s", err)
		return signup.NewSignupInternalServerError()
	}

	return signup.NewSignupOK().WithPayload(&genmodel.SignUpResponse{Success: true, Message: "You are successfully registered"})
}
