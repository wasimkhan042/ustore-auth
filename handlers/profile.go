package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	runtime "github.com/wasimkhan042/ustore-auth"
	genmodel "github.com/wasimkhan042/ustore-auth/gen/models"
	"github.com/wasimkhan042/ustore-auth/gen/restapi/operations/user"
)

// NewSignUp handles request for saving student
func NewUserProfile(rt *runtime.Runtime) user.ProfileHandler {
	return &userProfile{rt: rt}
}

type userProfile struct {
	rt *runtime.Runtime
}

func (u *userProfile) Handle(params user.ProfileParams, principal interface{}) middleware.Responder {

	userInfo, err := u.rt.Service().UserProfile(params.HTTPRequest.Header.Get("Authorization"))
	if err != nil {
		log().Errorf("failed to get userProfile data: %s", err)
		return user.NewProfileInternalServerError()
	}

	return user.NewProfileOK().WithPayload(&genmodel.Profile{
		FirstName:    userInfo.FirstName,
		MiddleName:   userInfo.MiddleName,
		LastName:     userInfo.LastName,
		Username:     userInfo.Username,
		Email:        userInfo.Email,
		ProfileImage: userInfo.ProfileImage,
	})
}
