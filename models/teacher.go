package models

//TeachList struct for form input when inputing a new teacher and shit
type TeachList struct {
	UserID     uint   `json:"userID"`
	SkillID    int    `json:"skillID"`
	Desc       string `json:"description"`
	Experience string `json:"experience"`
}
