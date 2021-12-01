package textme

import "errors"

type User struct {
	Id       int    `json:"-" db:"id"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Name     string `json:"name" binding:"required"`
	ImgUrl   string `json:"img_url"`
}

type UserInfo struct {
	Id       int    `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
	Name     string `json:"name" db:"name"`
	ImgUrl   string `json:"img_url" db:"img_url"`
}

type UpdateUserInput struct {
	Name   *string `json:"name"`
	ImgUrl *string `json:"img_url"`
}

func (i UpdateUserInput) Validate() error {
	if i.Name == nil && i.ImgUrl == nil {
		return errors.New("update structure has no values")
	}

	return nil
}
