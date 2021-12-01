package service

import (
	textme "github.com/signedcash/messenger_rest_api"
	"github.com/signedcash/messenger_rest_api/pkg/repository"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetById(id int) (textme.UserInfo, error) {
	return s.repo.GetById(id)
}

func (s *UserService) GetByName(name string) (textme.UserInfo, error) {
	return s.repo.GetByName(name)
}

func (s *UserService) Update(id int, input textme.UpdateUserInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(id, input)
}
