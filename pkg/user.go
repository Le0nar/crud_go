package news

type User struct {
	Id       int    `json:"-"`
	Login    string `json:"login"`
	Email    string `json:"email"`
	Password string `json:"password"`
}