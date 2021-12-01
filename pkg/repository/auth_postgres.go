package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	textme "github.com/signedcash/messenger_rest_api"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user textme.User) (int, error) {
	var id, profileId int

	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash, img_url) values ($1, $2, $3, $4) RETURNING id", usersTable)
	row := r.db.QueryRow(query, user.Name, user.Username, user.Password, "")
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	query = fmt.Sprintf("INSERT INTO %s (user_id, descript, age, country, city) values ($1, $2, $3, $4, $5) RETURNING id", profilesTable)
	row = r.db.QueryRow(query, id, "", 0, "", "")
	if err := row.Scan(&profileId); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthPostgres) GetUser(username, password string) (textme.User, error) {
	var user textme.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password_hash=$2", usersTable)
	err := r.db.Get(&user, query, username, password)
	return user, err
}
