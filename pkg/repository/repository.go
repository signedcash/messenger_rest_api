package repository

import (
	"github.com/jmoiron/sqlx"
	textme "github.com/signedcash/messenger_rest_api"
)

type Authorization interface {
	CreateUser(user textme.User) (int, error)
}

type Message interface {
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
	}
}
