package service

import (
	textme "github.com/signedcash/messenger_rest_api"
	"github.com/signedcash/messenger_rest_api/pkg/repository"
)

type ChatService struct {
	repo repository.Chat
}

func NewChatService(repo repository.Chat) *ChatService {
	return &ChatService{repo: repo}
}

func (s *ChatService) Create(chat textme.Chat) (int, error) {
	return s.repo.Create(chat)
}

func (s *ChatService) GetAllByUserId(userId int) ([]textme.Chat, error) {
	return s.repo.GetAllByUserId(userId)
}

func (s *ChatService) GetByUserId(user1Id, user2Id int) (textme.Chat, error) {
	return s.repo.GetByUserId(user1Id, user2Id)
}

func (s *ChatService) Update(userId, chatId int, input textme.UpdateChatInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(userId, chatId, input)
}
