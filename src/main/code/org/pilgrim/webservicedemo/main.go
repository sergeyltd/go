package main

import (
	"net/http"
	"github.com/pluralsight/webservicedemo/controllers"
)

func main(){
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}