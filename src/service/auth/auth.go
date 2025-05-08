package auth

import (
	"WeatherAgent/src/model"
	"WeatherAgent/src/utils/dbService"
	"crypto/md5"
	"errors"
	"github.com/bwmarrin/snowflake"
)

// 雪花算法生成ID
var node *snowflake.Node

func Register(user *model.User) (err error) {
	sqlStr := "select count(userid) from user where username = ?"
	var count int64
	err = dbService.Mysql_client.Raw(sqlStr, user.UserName).Scan(&count).Error
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("username already exists")
	}
	node, err = snowflake.NewNode(1)
	if err != nil {
		return errors.New("failed to create snowflake node")
	}
	id := node.Generate()
	passWd := md5.Sum([]byte(user.Password))
	insertSql := "insert into user (userid,username,password) values (?,?,?)"
	err = dbService.Mysql_client.Exec(insertSql, id, user.UserName, passWd).Error
	if err != nil {
		errors.New("failed to create user")
	}
	return
}
