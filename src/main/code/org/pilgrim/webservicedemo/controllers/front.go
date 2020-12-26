package controllers

import (
	"encoding/json"
	"net/http"
)

// RegisterControllers to register
func RegisterControllers() {
	uc := newUserController()

	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)

	ac := newAdminController()

	http.Handle("/admin", *ac)
	http.Handle("/admin/", *ac)
}


func encodeResponseAsJSON(data interface{}, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	enc := json.NewEncoder(w)
	enc.Encode(data)
}
