package models


import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// 用户登录token
type Visit_Token struct {
	//Id string `gorm:"primary_key;type:varchar(64);not null;unique"`  // column name will be `id` Set field as not nullable and unique

	ID         int64  `db:"id" json:"id"` // 主键id
	UserId     string `db:"user_id" json:"user_id"`      // 用户 uuid
	UserName   string `db:"user_name" json:"user_name"`    // 用户名
	OrgId      string `db:"org_id" json:"org_id"`       // 组织 id
	OrgName    string `db:"org_name" json:"org_name"`     // 组织名
	TokenId    string `db:"token_uuid" json:"token_uuid"`   // 随机字符串
	Token      string `db:"token" json:"-" `        // token,  添加索引这样快; json 直接忽略字段
	Role       int    `db:"role_uuid" json:"role_uuid"`    // 角色
	CreateTime string `db:"create_time" json:"create_time"`  // 创建时间
	UpdateTime string `db:"update_time" json:"update_time"`  // 更新时间
	Expiration int    `db:"expiration" json:"expiration"`   // 多长时间之后过期
	Deleted    int    `db:"deleted" json:"deleted"`      // 0 token没有进行退出操作,  1进行了操作退出操作,已经失效

}

func (Visit_Token) TableName() string {
  return "access_token"
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




