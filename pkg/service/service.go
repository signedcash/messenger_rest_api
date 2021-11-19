package service

import (
	textme "github.com/signedcash/messenger_rest_api"
	"github.com/signedcash/messenger_rest_api/pkg/repository"
)

type Authorization interface {
	CreateUser(user textme.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(accesToken string) (int, error)
}

type Message interface {
	Create(message textme.Message) (int, error)
	GetAllByChatId(userId, chatId int) ([]textme.Message, error)
}

type Chat interface {
}

type Profile interface {
}

type Service struct {
	Authorization
	Profile
	Chat
	Message
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Message:       NewMessageService(repos.Message),
	}
}
