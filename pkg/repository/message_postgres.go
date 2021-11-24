package repository

import (
	"fmt"
	"strings"

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

func (r *MessagePostgres) Delete(userId, messageId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE (id = $2) AND (sender_id = $1)", messagesTable)
	_, err := r.db.Exec(query, userId, messageId)

	return err
}

func (r *MessagePostgres) Update(userId, messageId int, input textme.UpdateMessageInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Content != nil {
		setValues = append(setValues, fmt.Sprintf("content=$%d", argId))
		args = append(args, *input.Content)
		argId++
	}

	if input.CreatedAt != nil {
		setValues = append(setValues, fmt.Sprintf("created_at=$%d", argId))
		args = append(args, *input.CreatedAt)
		argId++
	}

	if input.State != nil {
		setValues = append(setValues, fmt.Sprintf("state=$%d", argId))
		args = append(args, *input.State)
		argId++
	}

	if input.Type != nil {
		setValues = append(setValues, fmt.Sprintf("type=$%d", argId))
		args = append(args, *input.Type)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE sender_id = $%d AND id = $%d ",
		messagesTable, setQuery, argId, argId+1)

	args = append(args, userId, messageId)
	_, err := r.db.Exec(query, args...)

	return err
}
