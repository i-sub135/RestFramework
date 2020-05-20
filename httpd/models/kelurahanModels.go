package models

import (
	"fmt"

	"github.com/getsentry/sentry-go"
)

// GetKelurahan --
func (r *RestModels) GetKelurahan(resp *[]WilayahKelurahan, param string) {
	var (
		kel []WilayahKelurahan
	)
	err := sql.Table("t_wilayah_kelurahan").
		Where("kode like ?", param+".%").
		Find(&kel).Error
	if err != nil {
		sentry.CaptureException(err)
		fmt.Printf("%+v", err)
	}
	*resp = kel
}
