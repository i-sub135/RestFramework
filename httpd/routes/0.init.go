package routes

import (
	"rest-framework/httpd/controller"
	"rest-framework/httpd/response"
)

type (
	//RestAPI  --
	RestAPI struct{}
)

var (
	resp response.RestResponse
	ctrl controller.RestController
)
