package news

type News struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Date        int    `json:"date"`
}
