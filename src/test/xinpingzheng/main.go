package main

import (
	"fmt"
	"strconv"
	"test/xinpingzheng/database"
)

type User struct {
	ID         uint   `json:"id";gorm:"id"`
	OpenID     string `json:"open_id";gorm:"openid"`
	RealName   string `json:"real_name";gorm:"real_name"`
	NickName   string `json:"nick_name";gorm:"nickname"`
	Sex        uint8  `json:"sex";gorm:"sex"`
	City       string `json:"city"`
	Province   string `json:"province";gorm:"province"`
	Country    string `json:"country";gorm:"country"`
	HeadImgURL string `json:"head_img_url";gorm:"headimgurl"`
	Phone      string `json:"phone";gorm:"phone"`
	IDCard     string `json:"id_card";gorm:"id_card"`
}

func (User) TableName() string {
	return "xpz_borrow_users"
}

func main() {
	var users []User
	database.Eloquent.Find(&users)
	var i int
	for _, user := range users {
		if user.IDCard != "" {
			sexIndex := user.IDCard[len(user.IDCard)-2 : len(user.IDCard)-1]
			index, _ := strconv.Atoi(sexIndex)

			if index%2 == 0 {
				i++
				fmt.Println(i, user.RealName, user.Phone, user.IDCard)
			}
		}
	}
}
