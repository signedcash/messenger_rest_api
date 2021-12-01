package repository

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	textme "github.com/signedcash/messenger_rest_api"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (r *UserPostgres) GetById(id int) (textme.UserInfo, error) {
	var user textme.UserInfo

	query := fmt.Sprintf(`SELECT id, username, name, img_url
						  FROM %s
						  WHERE id = $1`,
		usersTable)
	err := r.db.Get(&user, query, id)

	return user, err
}

func (r *UserPostgres) GetByName(name string) (textme.UserInfo, error) {
	var user textme.UserInfo

	query := fmt.Sprintf(`SELECT id, username, name, img_url
						  FROM %s
						  WHERE name = $1 OR username = $1`,
		usersTable)
	err := r.db.Get(&user, query, name)

	return user, err
}

func (r *UserPostgres) Update(id int, input textme.UpdateUserInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.ImgUrl != nil {
		setValues = append(setValues, fmt.Sprintf("img_url=$%d", argId))
		args = append(args, *input.ImgUrl)
		argId++
	}

	if input.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, *input.Name)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = $%d",
		usersTable, setQuery, argId)

	args = append(args, id)
	_, err := r.db.Exec(query, args...)

	return err
}
