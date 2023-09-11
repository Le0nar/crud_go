package news

// type AchievementList struct {
// 	Id          int    `json:id`
// 	Title       string `json:title`
// 	Description string `json:description`
// }

type News struct {
	Id          int    `json:id`
	Title       string `json:title`
	Description string `json:description`
	// TODO: mb use another type for "Date"
	Date string `json:date`
}
