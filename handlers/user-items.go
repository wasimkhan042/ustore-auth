package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	runtime "github.com/wasimkhan042/ustore-auth"
	genmodel "github.com/wasimkhan042/ustore-auth/gen/models"
	"github.com/wasimkhan042/ustore-auth/gen/restapi/operations/user"
)

// NewSignUp handles request for saving student
func NewUserItemList(rt *runtime.Runtime) user.UseritmesHandler {
	return &userItemList{rt: rt}
}

type userItemList struct {
	rt *runtime.Runtime
}

func (u *userItemList) Handle(params user.UseritmesParams, principal interface{}) middleware.Responder {

	items, err := u.rt.Service().UserItemList(params.HTTPRequest.Header.Get("Authorization"))
	if err != nil {
		log().Errorf("failed to get userProfile data: %s", err)
		return user.NewProfileInternalServerError()
	}

	itemList := make([]*genmodel.UserProduct, 0)
	for _, item := range items {
		itemList = append(itemList, toItemGen(item))
	}
	return user.NewUseritmesOK().WithPayload(itemList)
}
