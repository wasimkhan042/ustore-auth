package handlers

import (
	"github.com/go-openapi/loads"
	runtime "github.com/wasimkhan042/ustore-auth"
	"github.com/wasimkhan042/ustore-auth/gen/restapi/operations"
)

// Handler replaces swagger handler
type Handler *operations.UStoreAuthAPI

// NewHandler overrides swagger api handlers
func NewHandler(rt *runtime.Runtime, spec *loads.Document) Handler {
	handler := operations.NewUStoreAuthAPI(spec)

	// registration handler
	handler.SignupSignupHandler = NewSignUp(rt)
	handler.LoginLoginHandler = NewLogin(rt)

	return handler
}
