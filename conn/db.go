package conn

import (
	"fmt"
	"social-backend/config"
	"gorm.io/driver/mysql"
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectDB() {
	conf := config.Db()
	log.Info("connecting to mysql at ", conf.Host, ":", conf.Port, "...")
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True",
		conf.User,
		conf.Pass,
		conf.Host,
		conf.Port,
		conf.Schema,
	)
	dB, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database: ", err)
	}
	log.Info("database connected")
	db = dB
}

func Db() *gorm.DB {
	return db
}