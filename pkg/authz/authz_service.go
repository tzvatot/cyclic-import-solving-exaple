package authz

import (
	"fmt"

	"example.com/m/v2/pkg/types"
)

type AuthorizationService interface {
	IsAllowed(userID string, action Action, resource Resource) (bool, error)
}

type UserProvider interface {
	Get(userID string) (types.User, error)
}

type Service struct {
	UserProvider UserProvider
}

func (s *Service) IsAllowed(userID string, action Action, resource Resource) (bool, error) {
	user, err := s.UserProvider.Get(userID)
	if err != nil {
		return false, err
	}
	// TODO implement
	fmt.Println("user: ", user)
	return true, nil
}
