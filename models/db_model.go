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

type UserBase struct {
	User_id string `gorm:"primary_key;type:varchar(64);not null;unique"` // 用户uuid

	Username string `gorm:"type:varchar(64);not null;unique"` // 用户名
	Email string `gorm:"type:varchar(64);not null;unique"`    // 邮箱
	Password string `gorm:"type:varchar(64);not null"`    // 邮箱
	Logo string `gorm:"type:varchar(64)"`    // 头像

	Deleted string `gorm:"type:varchar(1)"`    // 是否删除
	Reset_uuid string `gorm:"type:varchar(64)"`    // 头像
	Salt string `gorm:"type:varchar(64)"`    // 头像
	Sysadmin_flag string `gorm:"type:varchar(1)"`
	Creation_time string `gorm:"type:varchar(1)"`
	Update_time string `gorm:"type:varchar(1)"`
}

func (UserBase) TableName() string {
  return "user"
}




