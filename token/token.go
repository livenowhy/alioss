package token




import "errors"

type Labels map[string][]string


// token认证接口
type Authenticator interface {

	// token 认证
	Authenticate(token string) (bool, Labels, error)

	Name() string
}


var MysqlNoMatch = errors.New("There is no matching user in the database")
var MysqlWrongPass = errors.New("The user name and password and don't match in the database")
var MysqlUnable = errors.New("Unable to communicate with Mysql.")