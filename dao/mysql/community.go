package mysql

import (
	"bluebell/models"
	"errors"

	"gorm.io/gorm"

	"go.uber.org/zap"
)

// GetCommunityList 查询Community的sql语句
func GetCommunityList() (communityList []*models.Community, err error) {
	results := db.Select("community_id", "community_name").Find(&communityList)
	if results.RowsAffected == 0 {
		zap.L().Warn("there is no community in db")
		err = nil
	}
	return
}

// GetCommunityDetailByID 根据ID查询社区详情
func GetCommunityDetailByID(id int64) (community *models.CommunityDetail, err error) {
	result := db.Where("community_id", id).First(&community)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			err = ErrorInvalidID
		}
	}
	return community, err
}
