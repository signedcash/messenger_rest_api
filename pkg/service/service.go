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
	Delete(userId, messageId int) error
	Update(userId, messageId int, input textme.UpdateMessageInput) error
}

type Chat interface {
	Create(message textme.Chat) (int, error)
	GetAllByUserId(userId int) ([]textme.Chat, error)
	Update(userId, chatId int, input textme.UpdateChatInput) error
}

type Profile interface {
	GetByUserId(id int) (textme.Profile, error)
	Update(userId int, input textme.UpdateProfileInput) error
}

type User interface {
	GetById(id int) (textme.UserInfo, error)
	Update(id int, input textme.UpdateUserInput) error
}

type Service struct {
	Authorization
	Profile
	Chat
	Message
	User
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Message:       NewMessageService(repos.Message),
		Chat:          NewChatService(repos.Chat),
		Profile:       NewProfileService(repos.Profile),
		User:          NewUserService(repos.User),
	}
}
