package db

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"

	// mysql Driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// RestDBConf ---
type RestDBConf struct {
}

var (
	err     error
	dirver  string
	connect string
	db      *gorm.DB
)

func init() {
	er := godotenv.Load()
	if er != nil {
		log.Fatal("Main func --> Error loading .env file")
	}
	dirver = os.Getenv("DRIVER")
	connect = os.Getenv("CONNECT")
}

// Conenect -- db Conection
func (r *RestDBConf) Conenect() *gorm.DB {

	db, err = gorm.Open(dirver, connect)
	if err != nil {
		panic(err.Error())
	}
	db.DB().SetMaxIdleConns(5)
	db.DB().SetMaxOpenConns(10)

	return db
}
