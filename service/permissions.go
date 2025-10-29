package service

import (
	"net/http"

	"github.com/titpetric/platform/module/user"
)

type Permissions struct {
	Create bool
}

func NewPermissions(r *http.Request) Permissions {
	var result Permissions
	result.Create = user.IsLoggedIn(r)
	return result
}
