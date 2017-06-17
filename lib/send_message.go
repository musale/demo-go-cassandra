package lib

import (
	"encoding/json"
	"net/http"
)

//SendMessage on /messages
func SendMessage(w http.ResponseWriter, r *http.Request) {
	if r.Method == GET {
		json.NewEncoder(w).Encode(Response{
			Status: "failed", Message: "GET not allowed on /messages/send",
		})
		return
	}
	json.NewEncoder(w).Encode(Response{
		Status: "success", Message: "message received",
	})
	return
}
