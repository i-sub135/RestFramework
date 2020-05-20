package models

import (
	"rest-framework/httpd/db"

	"github.com/jinzhu/gorm"
)

type (
	// RestModels -- reciver
	RestModels struct{}
)

var (
	d   db.RestDBConf
	sql *gorm.DB = d.Conenect()
)
