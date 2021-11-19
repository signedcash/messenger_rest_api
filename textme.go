package textme

type Profile struct {
	Id       int    `json:"-"`
	UserId   int    `json:"user_id"`
	Descript string `json:"descript"`
	Age      int    `json:"age"`
	Country  string `json:"country"`
	City     string `json:"city"`
}

type Message struct {
	Id        int    `json:"-" db:"id"`
	ChatId    int    `json:"chat_id" db:"chat_id" binding:"required"`
	SenderId  int    `json:"sender_id" db:"sender_id" binding:"required"`
	Content   string `json:"content" db:"content" binding:"required"`
	CreatedAt string `json:"created_at" db:"created_at" binding:"required"`
	State     int    `json:"state" db:"state" binding:"required"`
	Type      int    `json:"type" db:"type" binding:"required"`
}

type Chat struct {
	Id      int `json:"-"`
	User1Id int `json:"user1_id"`
	User2Id int `json:"user2_id"`
}
