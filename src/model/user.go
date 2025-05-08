package model

type User struct {
	UserID   uint   `gorm:"primary_key" json:"id"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
}
