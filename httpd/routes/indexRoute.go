package routes

import (
	"rest-framework/httpd/controller"
	"rest-framework/httpd/response"

	"github.com/gin-gonic/gin"
)

type (
	//RestAPI  --
	RestAPI struct{}
)

var (
	resp response.RestResponse
	ctrl controller.RestController
)

// IndexRoute  --
func (r *RestAPI) IndexRoute(e *gin.Engine) {

	e.GET("/", ctrl.IndexCtrl)
	e.GET("/provinsi", ctrl.ProvCtrl)
	e.GET("/kota/:id", ctrl.KotaCtrl)
	e.GET("/kecamatan/:id", ctrl.KecamatanCtrl)
	e.GET("/kelurahan/:id", ctrl.KelurahanCtrl)

}
