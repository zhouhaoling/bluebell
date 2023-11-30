package models

import "time"

type Community struct {
	CommunityID   int    `json:"id" gorm:"column:community_id"`
	CommunityName string `json:"name" gorm:"column:community_name"`
}

type CommunityDetail struct {
	CommunityID   int64     `json:"id" gorm:"column:community_id"`
	CommunityName string    `json:"name" gorm:"column:community_name"`
	Introduction  string    `json:"introduction,omitempty" gorm:"column:introduction"`
	CreateTime    time.Time `json:"create_time" gorm:"column:create_time"`
}

// TableName 对应数据库中的表名
func (CommunityDetail) TableName() string {
	return "community"
}
