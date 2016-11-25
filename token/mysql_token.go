package token

import (
	"github.com/golang/glog"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/liuzhangpei/alioss/models"
	"github.com/liuzhangpei/alioss/server"
)

type MysqlConfig struct {
	Host string `yaml:"host,omitempty"`
	Port string `yaml:"port,omitempty"`
	Charset string `yaml:"charset,omitempty"`
	User string `yaml:"user,omitempty"`
	Pawd string `yaml:"pawd,omitempty"`
	Cydb string `yaml:"cydb,omitempty"`
}

type TokenMysqlAuthorizer struct {
	GormEngine       string  //
	Config           *MysqlConfig   // 数据库配置
}

func NewTokenAuthorizer(c *MysqlConfig) (*TokenMysqlAuthorizer, error) {

	glog.V(2).Infof("add lzp Auth NewACLMysqlAuthorizer: ")
	engine_var := c.User + ":" + c.Pawd + "@tcp(" + c.Host + ":" + c.Port + ")/" + c.Cydb + "?charset=" + c.Charset

	authorizer := &TokenMysqlAuthorizer{
		GormEngine: engine_var,
		Config: c,
	}
	return authorizer, nil
}


// 验证token
func (ma *TokenMysqlAuthorizer) Authenticate(token string) (retbool bool, err error) {
	fmt.Println("ss")


	db_session, err := gorm.Open("mysql", ma.GormEngine)

	glog.V(2).Infof("lzp --> CreateEngine GormEngine : %s", ma.GormEngine)

	if err != nil {
		fmt.Println("lzp --> CreateEngine : ")
		glog.V(2).Infof("lzp --> CreateEngine : %s", err.Error())
		return false, MysqlUnable
	}
	defer db_session.Close()

	var VisitToken models.Visit_Token

	db := db_session.Where(&models.Visit_Token{Token: token}).First(&VisitToken)

	if db.Error != nil {
		glog.V(2).Info("---lzp db.Error != nil   ")
		return false, MysqlNoMatch
	}
	fmt.Println("db.RowsAffected")
	fmt.Println(db.RowsAffected)
	if db.RowsAffected >= 1 {
		return true, nil
	} else {
		return false, MysqlNoMatch
	}
}

func (ma *TokenMysqlAuthorizer) Name() string  {
	return "mysql_token"
}

func CheckToken(token string) (retbool bool, err error) {
	TokenA, err :=  NewTokenAuthorizer(&server.CONF.MysqlConf)
	if err != nil {
		return false, err
	} else {
		return TokenA.Authenticate("yJ1aWQiOiAiYWMwYjVhMTEtOTZhYS0zN2E1LTk5MmYtZTVhNDNmZTVjNTVkIiwgInVzZXJfb3JhZyI6ICJ6aGFuZ3NhaSIsICJ0b2tlbmlkIjogImIyOWI2YzFhNDI3MWQyYmVhMWQ2ZTY1YSIsICJ1c2VyX3V1aWQiOiAiYWMwYjVhMTEtOTZhYS0zN2E1LTk5MmYtZTVhNDNmZTVjNTVkIiwgImV4cGlyZXMiOiAxNDgwMjc0MTQyLjI5OTI0NiwgInVzZXJfcm9sZSI6ICIxIiwgInVzZXJfaXAiOiAiMTI3LjAuMC4xIiwgInVzZXJfb3JnYSI6ICJ6aGFuZ3NhaSIsICJyb2xlX3V1aWQiOiAyMDAsICJvcmdhX3V1aWQiOiAiYWMwYjVhMTEtOTZhYS0zN2E1LTk5MmYtZTVhNDNmZTVjNTVkIiwgInNhbHQiOiAiMzBmYTkzODc0NGRmMGU5YmI0NGZmMDJkIiwgImVtYWlsIjogIjEyM0BxcS5jb20iLCAidXNlcl9uYW1lIjogInpoYW5nc2FpIn0v0OPasrqBpyG_VLxfE2tq")
	}
}