package models

import (
	"fmt"
)

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

// func (s *Skill) GetTeachers() ()

//GetAllSkills gets all skills
func GetAllSkills() []Skill {
	skills := make([]Skill, 10)
	GetDB().Table("skills").Find(&skills)
	// if err != nil {
	// 	fmt.Println("Fuck me ", err)
	// 	return nil
	// }
	fmt.Println(skills)
	sk := make(map[int]string, 10)
	for _, i := range skills {
		sk[i.ID] = i.Name
	}
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
