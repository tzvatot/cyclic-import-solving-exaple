package pkg

import (
	"example.com/m/v2/pkg/authz"
	"example.com/m/v2/pkg/user"
)

func main() {
	authzService := &authz.Service{}
	userService := user.UserService{}
	userService.AuthzService = authzService
	authzService.UserProvider = &userService
}
