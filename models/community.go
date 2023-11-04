package models

import "time"

type Community struct {
	ID   int    `json:"id" gorm:"column:community_id"`
	Name string `json:"name" gorm:"column:community_name"`
}

type CommunityDetail struct {
	ID           int64     `json:"id" gorm:"column:community_id"`
	Name         string    `json:"name" gorm:"column:community_name"`
	Introduction string    `json:"introduction,omitempty"`
	CreateTime   time.Time `json:"create_time" gorm:"column:create_time"`
}

// TableName 对应数据库中的表名
func (CommunityDetail) TableName() string {
	return "community"
}
