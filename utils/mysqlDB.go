package utils

import (
    _ "github.com/lib/pq"
	_ "github.com/go-sql-driver/mysql"

	"github.com/jmoiron/sqlx"

)

func CreateMysqlDb(dataSourceName string) *sqlx.DB{
	return sqlx.MustConnect("mysql", dataSourceName)
}

func UpdateLogo(user_id, logo, dataSourceName string) error {

	db := CreateMysqlDb(dataSourceName)

	const findProjectSql = `UPDATE user SET logo=? WHERE user_id=?`
	tx := db.MustBegin()
    tx.MustExec(findProjectSql, logo, user_id)
	return tx.Commit()

}

