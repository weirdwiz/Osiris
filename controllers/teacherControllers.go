package controllers

import (
	"encoding/json"
	"net/http"

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
