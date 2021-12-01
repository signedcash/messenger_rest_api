package textme

import "errors"

type Profile struct {
	Id       int    `json:"-" db:"id"`
	UserId   int    `json:"user_id" db:"user_id" binding:"required"`
	Descript string `json:"descript" db:"descript" binding:"required"`
	Age      int    `json:"age" db:"age" binding:"required"`
	Country  string `json:"country" db:"country" binding:"required"`
	City     string `json:"city" db:"city" binding:"required"`
}

type Message struct {
	Id        int    `json:"-" db:"id"`
	ChatId    int    `json:"chat_id" db:"chat_id" binding:"required"`
	SenderId  int    `json:"sender_id" db:"sender_id" binding:"required"`
	Content   string `json:"content" db:"content" binding:"required"`
	CreatedAt string `json:"created_at" db:"created_at" binding:"required"`
}

type Chat struct {
	Id      int `json:"id" db:"id"`
	User1Id int `json:"user1_id" db:"user1_id" binding:"required"`
	User2Id int `json:"user2_id" db:"user2_id" binding:"required"`
	Vision1 int `json:"vision_1" db:"vision_1" binding:"required"`
	Vision2 int `json:"vision_2" db:"vision_2" binding:"required"`
}

type UpdateProfileInput struct {
	Descript *string `json:"descript"`
	Age      *int    `json:"age"`
	Country  *string `json:"country"`
	City     *string `json:"city"`
}

func (i UpdateProfileInput) Validate() error {
	if i.Descript == nil && i.Age == nil && i.Country == nil && i.City == nil {
		return errors.New("update structure has no values")
	}

	return nil
}

type UpdateMessageInput struct {
	Content   *string `json:"content"`
	CreatedAt *string `json:"created_at"`
}

func (i UpdateMessageInput) Validate() error {
	if i.Content == nil && i.CreatedAt == nil {
		return errors.New("update structure has no values")
	}

	return nil
}

type UpdateChatInput struct {
	Vision1 *int `json:"vision_1"`
	Vision2 *int `json:"vision_2"`
}

func (i UpdateChatInput) Validate() error {
	if i.Vision1 == nil && i.Vision2 == nil {
		return errors.New("update structure has no values")
	}

	return nil
}
