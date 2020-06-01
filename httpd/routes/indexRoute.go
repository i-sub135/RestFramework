package routes

import (
	"github.com/gin-gonic/gin"
)

// IndexRoute  --
func (r *RestAPI) IndexRoute(e *gin.Engine) {

	e.GET("/", ctrl.IndexCtrl)
	e.GET("/provinsi", ctrl.ProvCtrl)
	e.GET("/kota/:id", ctrl.KotaCtrl)
	e.GET("/kecamatan/:id", ctrl.KecamatanCtrl)
	e.GET("/kelurahan/:id", ctrl.KelurahanCtrl)

}
