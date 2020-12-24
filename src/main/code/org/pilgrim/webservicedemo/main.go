package main

import (
	"github.com/pluralsight/webservicedemo/controllers"
	"net/http"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
