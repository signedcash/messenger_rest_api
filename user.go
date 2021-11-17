package textme

type User struct {
	Id       int    `json:"-"`
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
	ImgUrl   string `json:"img_url"`
}
