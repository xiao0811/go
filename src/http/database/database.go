package database

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var Eloquent *gorm.DB

func init() {
	var err error
	Eloquent, err = gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/" +
		"lelel?charset=utf8&parseTime=True&loc=Local&timeout=10ms")

	if err != nil {
		fmt.Println("Mysql connect error,", err)
		return
	}

	if Eloquent.Error != nil {
		fmt.Println("Database error,", Eloquent.Error)
		return
	}
}
