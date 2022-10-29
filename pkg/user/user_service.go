package user

import (
	"errors"

	"example.com/m/v2/pkg/authz"
	"example.com/m/v2/pkg/types"
)

type UserService struct {
	AuthzService authz.AuthorizationService
}

func (s *UserService) Create(user types.User) error {
	// TODO implement
	return nil
}

func (s *UserService) Get(userID string) (types.User, error) {
	// TODO implement
	return types.User{}, nil
}

func (s *UserService) Update(user types.User) error {
	// TODO implement
	return nil
}

func (s *UserService) Delete(userID string) error {
	// TODO implement
	allowed, err := s.AuthzService.IsAllowed(userID, authz.DeleteAction, authz.UserResource)
	if err != nil {
		return err
	}
	if !allowed {
		return errors.New("no allowed")
	}
	return nil
}
