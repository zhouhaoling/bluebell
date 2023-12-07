package mysql

import (
	"bluebell/models"

	"gorm.io/gorm"
)

func CreateComment(comment *models.Comment) error {
	tx := db.Create(comment)
	if tx.RowsAffected > 0 {
		return nil
	}
	return tx.Error
}

func GetCommentByPostIdCount(postID int64) (int64, error) {
	var comments []models.Comment
	tx := db.Table("comment").Where("post_id = ?", postID).Find(&comments)
	if tx.Error == gorm.ErrRecordNotFound {
		return 0, nil
	}
	return tx.RowsAffected, tx.Error
}

func GetCommentByPostID(postID int64) ([]models.Comment, error) {
	var comments []models.Comment
	tx := db.Table("comment").Where("post_id = ?", postID).Find(&comments)
	return comments, tx.Error
}
