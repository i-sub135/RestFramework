package models

import (
	"fmt"

	"github.com/getsentry/sentry-go"
)

// GetKota --
func (r *RestModels) GetKota(resp *[]WilayahKota, param string) {
	var (
		kota []WilayahKota
	)
	err := sql.Table("t_wilayah_kota").
		Where("kode like ?", param+".%").
		Find(&kota).Error
	if err != nil {
		sentry.CaptureException(err)
		fmt.Printf("%+v", err)
	}
	*resp = kota
}
