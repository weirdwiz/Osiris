package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/weirdwiz/osiris/models"
	u "github.com/weirdwiz/osiris/utils"
)

// TeacherNew route to create a new teacher
func TeacherNew(w http.ResponseWriter, r *http.Request) {
	f := models.Teacher{}
	userID := r.Context().Value("user").(uint)
	json.NewDecoder(r.Body).Decode(&f)
	f.UserID = userID
	models.GetDB().Create(f)
	u.Respond(w, u.Message(true, "success"))
}

//GetAllTeachers Sends all teachers regarding the skill
func GetAllTeachers(w http.ResponseWriter, r *http.Request) {
	sid := mux.Vars(r)["id"]
	i, _ := strconv.ParseUint(sid, 10, 16)
	s := models.GetSkill(uint(i))
	t := s.GetTeachersForSkill()
	resp := u.Message(true, "success")
	resp["teachers"] = t
	u.Respond(w, resp)
}
