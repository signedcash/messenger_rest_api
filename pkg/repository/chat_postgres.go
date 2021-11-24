package repository

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	textme "github.com/signedcash/messenger_rest_api"
)

type ChatPostgres struct {
	db *sqlx.DB
}

func NewChatPostgres(db *sqlx.DB) *ChatPostgres {
	return &ChatPostgres{db: db}
}

func (r *ChatPostgres) Create(chat textme.Chat) (int, error) {
	var id int
	query := fmt.Sprintf(`INSERT INTO %s (user1_id, user2_id, vision_1, vision_2) 
						  values ($1, $2, $3, $4) 
						  RETURNING id`, chatsTable)
	row := r.db.QueryRow(query, chat.User1Id, chat.User2Id, chat.Vision1, chat.Vision2)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *ChatPostgres) GetAllByUserId(userId int) ([]textme.Chat, error) {
	var chats []textme.Chat

	query := fmt.Sprintf(`SELECT *
						  FROM %s
						  WHERE (user1_id = $1 AND vision_1 = 1) OR 
						  	    (user2_id = $1 AND vision_2 = 1)`,
		chatsTable)
	err := r.db.Select(&chats, query, userId)

	return chats, err
}

func (r *ChatPostgres) Update(userId, chatId int, input textme.UpdateChatInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Vision1 != nil {
		setValues = append(setValues, fmt.Sprintf("vision_1=$%d", argId))
		args = append(args, *input.Vision1)
		argId++
	}

	if input.Vision2 != nil {
		setValues = append(setValues, fmt.Sprintf("vision_2=$%d", argId))
		args = append(args, *input.Vision2)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE (user1_id = $%d OR user2_id = $%d) AND id = $%d ",
		chatsTable, setQuery, argId, argId, argId+1)

	args = append(args, userId, chatId)
	_, err := r.db.Exec(query, args...)

	return err
}
