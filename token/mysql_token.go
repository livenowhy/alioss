package token

import (
	"github.com/golang/glog"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/liuzhangpei/alioss/models"
)

type MysqlConfig struct {
	Host string `yaml:"host,omitempty"`
	Port string `yaml:"port,omitempty"`
	Charset string `yaml:"charset,omitempty"`
	User string `yaml:"user,omitempty"`
	Pawd string `yaml:"pawd,omitempty"`
	Cydb string `yaml:"cydb,omitempty"`
	DataSourceName  string `yaml:"dataSourceName,omitempty"`
}

type TokenMysqlAuthorizer struct {
	GormEngine       string  //
	Config           *MysqlConfig   // 数据库配置
}

func (myc *MysqlConfig)NewTokenAuthorizer() (*TokenMysqlAuthorizer, error) {

	glog.V(2).Infof("add lzp Auth NewACLMysqlAuthorizer: ")
	engine_var := myc.User + ":" + myc.Pawd + "@tcp(" + myc.Host + ":" + myc.Port + ")/" + myc.Cydb + "?charset=" + myc.Charset

	authorizer := &TokenMysqlAuthorizer{
		GormEngine: engine_var,
		Config: myc,
	}
	return authorizer, nil
}


// 验证token
func (ma *TokenMysqlAuthorizer) Authenticate(token string) (retbool bool, err error, vt *models.Visit_Token) {
	fmt.Println("ss")
	fmt.Println(ma.GormEngine)


	db_session, err := gorm.Open("mysql", ma.GormEngine)

	glog.V(2).Infof("lzp --> CreateEngine GormEngine : %s", ma.GormEngine)

	if err != nil {
		fmt.Println("lzp --> CreateEngine : ")
		glog.V(2).Infof("lzp --> CreateEngine : %s", err.Error())
		return false, MysqlUnable, nil
	}
	defer db_session.Close()

	var VisitToken models.Visit_Token

	db := db_session.Where(&models.Visit_Token{Token: token}).First(&VisitToken)

	if db.Error != nil {
		fmt.Println("---lzp db.Error != nil   ")
		fmt.Println(db.Error.Error())
		return false, db.Error, nil
	}
	fmt.Println("db.RowsAffected")
	fmt.Println(db.RowsAffected)
	if db.RowsAffected >= 1 {
		return true, nil, &VisitToken
	} else {
		return false, MysqlNoMatch, nil
	}
}

func (ma *TokenMysqlAuthorizer) Name() string  {
	return "mysql_token"
}

func (myc *MysqlConfig)CheckToken(token string) (retbool bool, err error, vt *models.Visit_Token) {
	TokenA, err :=  myc.NewTokenAuthorizer()
	if err != nil {
		return false, err, nil
	} else {
		return TokenA.Authenticate(token)
	}
}