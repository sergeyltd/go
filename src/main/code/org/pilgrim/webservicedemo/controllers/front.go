package controllers

import "net/http"

// RegisterControllers to register
func RegisterControllers() {
	uc := newUserController()

	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)

	ac := newAdminController()

	http.Handle("/admin", *ac)
	http.Handle("/admin/", *ac)
}
