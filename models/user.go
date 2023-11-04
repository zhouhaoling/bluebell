package models

type User struct {
	UserID   int64 `gorm:"column:user_id"`
	Username string
	Password string
}
