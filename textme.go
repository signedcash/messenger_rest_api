package textme

import "time"

type DateType time.Time

type Profiles struct {
	Id       int    `json:"-"`
	UserId   int    `json:"user_id"`
	Descript string `json:"descript"`
	Age      int    `json:"age"`
	Country  string `json:"country"`
	City     string `json:"city"`
}

type Messages struct {
	Id        int      `json:"-"`
	ChatId    int      `json:"chat_id"`
	SenderId  int      `json:"sender_id"`
	Content   string   `json:"descript"`
	CreatedAt DateType `json:"age"`
	Country   string   `json:"country"`
	City      string   `json:"city"`
}

type Chats struct {
	Id      int `json:"-"`
	User1Id int `json:"user1_id"`
	User2Id int `json:"user2_id"`
}
