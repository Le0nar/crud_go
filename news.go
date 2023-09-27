package news

type News struct {
	Id          int       `json:"-" db:"id"`
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Date        string `json:"-"`
	UserId 		int       `json:"userId" binding:"required"`
}
