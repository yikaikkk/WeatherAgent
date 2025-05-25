package model

type User struct {
	Id       uint   `gorm:"primary_key" json:"id"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
	UserId   string `json:"user_id"`
}
