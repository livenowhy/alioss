package models


import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// 用户登录token
type Visit_Token struct {
	Id string `gorm:"primary_key;type:varchar(64);not null;unique"`  // column name will be `id` Set field as not nullable and unique

	Tokenid string `gorm:"type:varchar(64);unique"`  // token id
	Token string `gorm:"type:varchar(1024)"`    // token

	// 0 token没有进行退出操作,  1进行了操作退出操作,已经失效
	Deleted int

	Creation_time string `gorm:"type:varchar(64)"`   // 添加该记录时间
	Update_time   string `gorm:"type:varchar(64)"`  // 更新该记录时间
}

func (Visit_Token) TableName() string {
  return "visit_token"
}