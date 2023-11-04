package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
)

// GetCommunityList 查找到所有的community并返回
func GetCommunityList() ([]*models.Community, error) {
	return mysql.GetCommunityList()
}

// GetCommunityDetail 查找到id对应的communityDetail并返回
func GetCommunityDetail(id int64) (*models.CommunityDetail, error) {
	return mysql.GetCommunityDetailByID(id)
}
