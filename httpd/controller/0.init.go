package controller

import (
	"rest-framework/httpd/models"
	"rest-framework/httpd/response"
)

type (
	// RestController -- reciver Controllrt
	RestController struct{}
)

var (
	resp  response.RestResponse
	model models.RestModels
)
