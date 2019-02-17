package models

// SkillList list for skills (hack lol)
type SkillList struct {
	List []int `json:"skills"`
}

// Skill struct contains the skill
type Skill struct {
	ID          int        `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"decription"`
	Teachers    []*Account `json:"teachers" gorm:"many2many:account_teachers"`
}

// GetTeachersForSkill pretty obvious
func (s *Skill) GetTeachersForSkill() []*Teacher {
	t := make([]*Teacher, 0)
	GetDB().Table("teachers").Where("skill_id = ?", s.ID).Find(&t)
	return t
}

//GetAllSkills gets all skills
func GetAllSkills() []Skill {
	skills := make([]Skill, 10)
	GetDB().Table("skills").Find(&skills)
	return skills
}

// GetSkill returns the skill using id number
func GetSkill(u uint) *Skill {
	s := &Skill{}
	GetDB().Table("skills").Where("id = ?", u).First(s)
	if s.Name == "" {
		return nil
	}
	return s
}
