package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/weirdwiz/osiris/models"
	u "github.com/weirdwiz/osiris/utils"
)

// //NewStrength adds new skill to user
// func NewStrength(w http.ResponseWriter, r *http.Request) {
// 	userID := r.Context().Value("user").(uint)
// 	user := models.GetUser(userID)
// 	var sk models.SkillList
// 	json.NewDecoder(r.Body).Decode(&sk)

// 	res := user.SetStrengths(sk.List)
// 	u.Respond(w, res)
// }

// GetSkills gets all skills
func GetSkills(w http.ResponseWriter, r *http.Request) {
	data := models.GetAllSkills()
	resp := u.Message(true, "success")
	resp["skills"] = data
	u.Respond(w, resp)
}

// TeachNew route to create a new teacher
func TeachNew(w http.ResponseWriter, r *http.Request) {
	f := models.Teacher{}
	userID := r.Context().Value("user").(uint)
	json.NewDecoder(r.Body).Decode(&f)
	f.UserID = userID
	models.GetDB().Create(f)
	u.Respond(w, u.Message(true, "success"))
}
