package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var Eloquent *gorm.DB

func init() {
	var err error
	// root  Xhwlkj1204$
	Eloquent, err = gorm.Open("mysql",
		"root:Xhwlkj1204$@tcp(47.101.181.255:3306)/xinpingzheng?charset=utf8" +
		"&parseTime=True&loc=Local")

	if err != nil {
		fmt.Println("Mysql connect error,", err)
		return
	}

	if Eloquent.Error != nil {
		fmt.Println("Database error,", Eloquent.Error)
		return
	}
}
