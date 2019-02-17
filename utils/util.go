package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Message encodes into a message
func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

// Respond with json
func Respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	b, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Fucking shit ", err)
	}
	w.Write(b)
}
