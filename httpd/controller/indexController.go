package controller

import (
	"rest-framework/httpd/models"
	"rest-framework/httpd/response"

	"github.com/gin-gonic/gin"
)

type (
	// RestController -- reciver Controllrt
	RestController struct{}
)

var (
	resp  response.RestResponse
	model models.RestModels
)

// IndexCtrl --
func (r *RestController) IndexCtrl(res *gin.Context) {
	out := resp.RespOK(map[string]interface{}{
		"Services": "new Micro",
	})
	res.JSON(out.Code, out)

}

// ProvCtrl --
func (r *RestController) ProvCtrl(res *gin.Context) {
	var (
		prov  []models.WilayahProvinsi
		param = res.Query("q")
	)

	model.GetProv(&prov, param)
	out := resp.RespOK(prov)
	res.JSON(out.Code, out)
}

// KotaCtrl --
func (r *RestController) KotaCtrl(res *gin.Context) {
	var (
		kota  []models.WilayahKota
		param = res.Param("id")
	)

	model.GetKota(&kota, param)
	out := resp.RespOK(kota)
	res.JSON(out.Code, out)
}

// KecamatanCtrl --
func (r *RestController) KecamatanCtrl(res *gin.Context) {
	var (
		kec   []models.WilayahKecamatan
		param = res.Param("id")
	)

	model.GetKecamatan(&kec, param)
	out := resp.RespOK(kec)
	res.JSON(out.Code, out)
}

// KelurahanCtrl --
func (r *RestController) KelurahanCtrl(res *gin.Context) {
	var (
		kel   []models.WilayahKelurahan
		param = res.Param("id")
	)

	model.GetKelurahan(&kel, param)
	out := resp.RespOK(kel)
	res.JSON(out.Code, out)
}
