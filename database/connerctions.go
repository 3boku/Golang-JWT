package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() {
	_, err := gorm.Open(mysql.Open("root:0815nuny@/yt_go_auth"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
}
