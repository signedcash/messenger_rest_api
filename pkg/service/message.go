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

func (s *MessageService) GetLastByChatId(userId, chatId int) (textme.Message, error) {
	return s.repo.GetLastByChatId(userId, chatId)
}

func (s *MessageService) Delete(userId, messageId int) error {
	return s.repo.Delete(userId, messageId)
}

func (s *MessageService) Update(userId, messageId int, input textme.UpdateMessageInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(userId, messageId, input)
}
