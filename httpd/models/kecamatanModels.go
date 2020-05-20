package models

import (
	"fmt"

	"github.com/getsentry/sentry-go"
)

// GetKecamatan --
func (r *RestModels) GetKecamatan(resp *[]WilayahKecamatan, params string) {
	var kec []WilayahKecamatan
	err := sql.Table("t_wilayah_kecamatan").
		Where("kode like ?", params+"%").
		Find(&kec).Error
	if err != nil {
		sentry.CaptureException(err)
		fmt.Printf("%+v", err)
	}
	*resp = kec
}
