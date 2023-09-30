package news

type News struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	// TODO: use timestamp instead of string for date+time type
	Date   string `json:"-"`
	UserId int    `json:"userId" binding:"required" db:"user_id"`
}
