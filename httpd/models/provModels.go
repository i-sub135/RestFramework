package models

import (
	"fmt"

	"github.com/getsentry/sentry-go"
)

// GetProv --
func (r *RestModels) GetProv(resp *[]WilayahProvinsi, params string) {
	var prov []WilayahProvinsi
	err := sql.Table("t_wilayah_propinsi").
		Where("nama like ?", "%"+params+"%").
		Find(&prov).Error
	if err != nil {
		sentry.CaptureException(err)
		fmt.Printf("%+v", err)
	}
	*resp = prov
}
