package service

import (
	textme "github.com/signedcash/messenger_rest_api"
	"github.com/signedcash/messenger_rest_api/pkg/repository"
)

type MessageService struct {
	repo repository.Message
}

func NewMessageService(repo repository.Message) *MessageService {
	return &MessageService{repo: repo}
}

func (s *MessageService) Create(message textme.Message) (int, error) {
	return s.repo.Create(message)
}

func (s *MessageService) GetAllByChatId(userId, chatId int) ([]textme.Message, error) {
	return s.repo.GetAllByChatId(userId, chatId)
}
