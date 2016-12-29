package utils

import (
    _ "github.com/lib/pq"
	_ "github.com/go-sql-driver/mysql"

	"github.com/jmoiron/sqlx"

)

func CreateMysqlDb(dataSourceName string) *sqlx.DB{
	return sqlx.MustConnect("mysql", dataSourceName)
}

func UpdateUserAvatars(user_id, logo, dataSourceName string) error {
	// 更新用户头像
	db := CreateMysqlDb(dataSourceName)
	const findProjectSql = `UPDATE user SET logo=? WHERE user_id=?`
	tx := db.MustBegin()
    tx.MustExec(findProjectSql, logo, user_id)
	return tx.Commit()
}



func UpdateMirrorIcon(user_id, image_id, logo, dataSourceName string) error {
	// 更新镜像头像
	db := CreateMysqlDb(dataSourceName)
	const findProjectSql = `UPDATE image_repository SET logo=? WHERE uid=? and uuid=?`
	tx := db.MustBegin()
    tx.MustExec(findProjectSql, logo, user_id, image_id)
	return tx.Commit()

}
