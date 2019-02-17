package controllers

import (
	"net/http"

	"github.com/weirdwiz/osiris/models"
	u "github.com/weirdwiz/osiris/utils"
)

// NewSkill : adds new skill to user
// func NewSkill(w http.ResponseWriter, r *http.Request) {
// 	user := r.Context().Value("user").(uint)
// 	var skills []int
// 	r.Body
// 	user.SetSkill()
// }

// GetSkills gets all skills
func GetSkills(w http.ResponseWriter, r *http.Request) {
	data := models.GetAllSkills()
	resp := u.Message(true, "success")
	resp["skills"] = data
	u.Respond(w, resp)
}
