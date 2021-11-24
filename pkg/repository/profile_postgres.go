package repository

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	textme "github.com/signedcash/messenger_rest_api"
)

type ProfilePostgres struct {
	db *sqlx.DB
}

func NewProfilePostgres(db *sqlx.DB) *ProfilePostgres {
	return &ProfilePostgres{db: db}
}

func (r *ProfilePostgres) GetByUserId(userId int) (textme.Profile, error) {
	var profile textme.Profile

	query := fmt.Sprintf(`SELECT *
						  FROM %s
						  WHERE user_id = $1`,
		profilesTable)
	err := r.db.Get(&profile, query, userId)

	return profile, err
}

func (r *ProfilePostgres) Update(userId int, input textme.UpdateProfileInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Descript != nil {
		setValues = append(setValues, fmt.Sprintf("descript=$%d", argId))
		args = append(args, *input.Descript)
		argId++
	}

	if input.Age != nil {
		setValues = append(setValues, fmt.Sprintf("age=$%d", argId))
		args = append(args, *input.Age)
		argId++
	}

	if input.Country != nil {
		setValues = append(setValues, fmt.Sprintf("country=$%d", argId))
		args = append(args, *input.Country)
		argId++
	}

	if input.City != nil {
		setValues = append(setValues, fmt.Sprintf("city=$%d", argId))
		args = append(args, *input.City)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE user_id = $%d",
		profilesTable, setQuery, argId)

	args = append(args, userId)
	_, err := r.db.Exec(query, args...)

	return err
}
