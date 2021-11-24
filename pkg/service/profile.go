package service

import (
	textme "github.com/signedcash/messenger_rest_api"
	"github.com/signedcash/messenger_rest_api/pkg/repository"
)

type ProfileService struct {
	repo repository.Profile
}

func NewProfileService(repo repository.Profile) *ProfileService {
	return &ProfileService{repo: repo}
}

func (s *ProfileService) GetByUserId(user_id int) (textme.Profile, error) {
	return s.repo.GetByUserId(user_id)
}

func (s *ProfileService) Update(userId int, input textme.UpdateProfileInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(userId, input)
}
