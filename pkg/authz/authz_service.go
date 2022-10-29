package authz

import (
	"fmt"

	"example.com/m/v2/pkg/user"
)

type AuthorizationService interface {
	IsAllowed(userID string, action Action, resource Resource) (bool, error)
}

type Service struct {
	UserService user.UserService
}

func (s *Service) IsAllowed(userID string, action Action, resource Resource) (bool, error) {
	user, err := s.UserService.Get(userID)
	if err != nil {
		return false, err
	}
	// TODO implement
	fmt.Println("user: ", user)
	return true, nil
}
