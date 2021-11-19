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
}

type Chat interface {
}

type Profile interface {
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
	}
}
