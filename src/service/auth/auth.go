package auth

import (
	"WeatherAgent/src/model"
	"WeatherAgent/src/utils/dbService"
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/bwmarrin/snowflake"
)

// 雪花算法生成ID
var node *snowflake.Node

func Register(user *model.User) (err error) {
	sqlStr := "select count(user_id) from user where username = ?"
	var count int64
	err = dbService.GetMysqlClient().Raw(sqlStr, user.UserName).Scan(&count).Error
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
	passWd := fmt.Sprintf("%x", md5.Sum([]byte(user.Password)))
	insertSql := "insert into user (user_id,username,password) values (?,?,?)"
	err = dbService.GetMysqlClient().Exec(insertSql, id, user.UserName, passWd).Error
	if err != nil {
		errors.New("failed to create user")
	}
	return
}
