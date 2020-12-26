package controllers

import (
	"encoding/json"
	"io"
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


func encodeResponseAsJSON(data interface{}, w io.Writer) {
	enc := json.NewEncoder(w)
	enc.Encode(data)
}
