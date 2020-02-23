package gorm_demo

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type Member struct {
	Id            int       `gorm:"primary_key"`
	LocationId    int       `json:"location_id"`
	ApplicationId int       `json:"application_id"` // iOS：3测试，5生产。
	PushId        string    `json:"push_id"`
	Nickname      string    `json:"nickname"`
	Avatar        string    `json:"avatar"`
	AvatarId      int       `json:"avatar_id"`
	CreatedAt     time.Time `json:"created_at"`
	ActiveAt      time.Time `json:"active_at"`
	UpdatedAt     time.Time
	Uid           string `json:"uid"`
}

func Start() {
	db, err := gorm.Open("mysql", "root:abc123@/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.LogMode(true)

	memb := Member{}
	err = db.Where("id = ?", 100).First(&memb).Error
	if err != nil {
		panic(err)
	}
	fmt.Println(memb)
}
