package models

type User struct {
	UserID       int64  `gorm:"column:user_id"`
	UserName     string `gorm:"column:username"`
	Password     string
	AccessToken  string
	RefreshToken string
}
