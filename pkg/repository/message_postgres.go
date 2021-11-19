package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	textme "github.com/signedcash/messenger_rest_api"
)

type MessagePostgres struct {
	db *sqlx.DB
}

func NewMessagePostgres(db *sqlx.DB) *MessagePostgres {
	return &MessagePostgres{db: db}
}

func (r *MessagePostgres) Create(message textme.Message) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (chat_id, sender_id, content, created_at, state, type) values ($1, $2, $3, $4, $5, $6) RETURNING id", messagesTable)
	row := r.db.QueryRow(query, message.ChatId, message.SenderId, message.Content, message.CreatedAt, message.State, message.Type)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *MessagePostgres) GetAllByChatId(userId, chatId int) ([]textme.Message, error) {
	var messages []textme.Message

	query := fmt.Sprintf(`SELECT *
						  FROM %s
						  WHERE EXISTS (SELECT 1 
						   			    FROM %s 
										WHERE (user1_id = $1 OR user2_id = $1) 
										AND (id = $2))
						  AND chat_id = $2`,
		messagesTable, chatsTable)
	err := r.db.Select(&messages, query, userId, chatId)

	return messages, err
}
