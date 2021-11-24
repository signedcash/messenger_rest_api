package repository

import (
	"github.com/jmoiron/sqlx"
	textme "github.com/signedcash/messenger_rest_api"
)

type Authorization interface {
	CreateUser(user textme.User) (int, error)
	GetUser(username, password string) (textme.User, error)
}

type Message interface {
	Create(message textme.Message) (int, error)
	GetAllByChatId(userId, chatId int) ([]textme.Message, error)
	Delete(userId, messageId int) error
	Update(userId, messageId int, input textme.UpdateMessageInput) error
}

type Chat interface {
	Create(chat textme.Chat) (int, error)
	GetAllByUserId(userId int) ([]textme.Chat, error)
	Update(userId, chatId int, input textme.UpdateChatInput) error
}

type Profile interface {
	GetByUserId(userId int) (textme.Profile, error)
	Update(userId int, input textme.UpdateProfileInput) error
}

type Repository struct {
	Authorization
	Message
	Chat
	Profile
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Message:       NewMessagePostgres(db),
		Chat:          NewChatPostgres(db),
		Profile:       NewProfilePostgres(db),
	}
}
