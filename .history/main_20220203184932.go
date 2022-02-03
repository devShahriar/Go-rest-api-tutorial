package main

import (
	"net/http"

	"github.com/devShahriar/Go-rest-api-tutorial/handlers"
	"github.com/labstack/echo/v4"
)

func main() {

	server := echo.New()
	http.HandleFunc("/ping", func(rw http.ResponseWriter, r *http.Request) {

	})
	server.GET("/getDetails", handlers.GetUserDetails)

}
