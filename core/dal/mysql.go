package dal

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB *gorm.DB
)

func InitMySQL() (err error) {
	DB, err = gorm.Open("mysql", "root:123456@(121.41.93.250:3307)/alert_demo?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Printf("connect msyql failed, err: %v \n", err)
		return
	}
	// 测试是否能够连通
	return DB.DB().Ping()
}
