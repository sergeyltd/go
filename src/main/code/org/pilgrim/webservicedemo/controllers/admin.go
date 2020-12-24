package controllers

import (
	"net/http"
	"regexp"
)

type adminController struct {
	adminIDPattern *regexp.Regexp
}

func (ac adminController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from the Admin Controller!"))
}

func newAdminController() *adminController {
	return &adminController{
		adminIDPattern: regexp.MustCompile(`^/admin/(\d+)/?`),
	}
}
