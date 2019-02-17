package models

//Teacher struct for form input when inputing a new teacher and shit
type Teacher struct {
	UserID     uint   `json:"userID"`
	SkillID    int    `json:"skillID"`
	Desc       string `json:"description"`
	Experience string `json:"experience"`
}
