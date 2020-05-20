package main

import (
	"os"
	"rest-framework/httpd/middleware"
	"rest-framework/httpd/routes"

	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	SenDNS := os.Getenv("SENTRY_DNS")
	err = sentry.Init(sentry.ClientOptions{
		Dsn: SenDNS,
	})
	if err != nil {
		panic(err)
	}
}

func main() {
	mode := os.Getenv("MODE")
	gin.SetMode(mode)

	// initial route Engine
	route := gin.Default()

	// Initial Middleware
	var middle middleware.RestMiddleware

	// Register Middleware
	route.Use(
		middle.ErrorHandle(),
	)

	// initial Routes
	var app routes.RestAPI
	/*
		+ Register Route
	*/
	app.IndexRoute(route)

	/*
		- Register Route
	*/

	route.Run()

}
