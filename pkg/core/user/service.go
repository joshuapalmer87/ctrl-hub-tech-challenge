package user

import (
	"ctrl-hub-technical-challenge/pkg/core/model"
	"errors"
)

// TODO - actually make this a real service either connecting to data or to another supplying service

type Service struct {
	userMap map[string]model.User
}

func NewService() *Service {
	service := &Service{
		userMap: make(map[string]model.User),
	}

	// Statically add some data for examples sake
	service.userMap["713be58e-0d79-4df2-a85c-9f44ca513a7d"] = model.User{
		ID:   "713be58e-0d79-4df2-a85c-9f44ca513a7d",
		Name: "Bobby Tables",
	}
	service.userMap["713be58e-0d79-4df2-a85c-9f44ca513a7e"] = model.User{
		ID:   "713be58e-0d79-4df2-a85c-9f44ca513a7e",
		Name: "Tobby Bables",
	}
	return service
}

func (s *Service) GetUserByID(id string) (model.User, error) {
	equipment, ok := s.userMap[id]
	if !ok {
		return model.User{}, errors.New("user not found")
	}
	return equipment, nil
}
