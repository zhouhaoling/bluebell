package mysql

import (
	"bluebell/models"
	"fmt"

	"gorm.io/gorm/clause"
)

// CreatePost 创建帖子
func CreatePost(p *models.Post) error {
	result := db.Create(&p)
	return result.Error
}

// GetPostById 根据帖子id查询帖子
func GetPostById(pid int64) (post *models.Post, err error) {
	post = new(models.Post)
	result := db.Where("post_id = ?", pid).First(post)
	return post, result.Error
}

// GetPostList 获取帖子列表
func GetPostList(page, size int) (posts []*models.Post, err error) {
	posts = make([]*models.Post, 0, 2)
	result := db.Order("create_time desc").Limit(size).Offset((page - 1) * size).Find(&posts)
	err = result.Error
	return
}

// GetPostListByIDs 根据给定id列表查询帖子数据
func GetPostListByIDs(ids []string) ([]*models.Post, error) {
	var posts []*models.Post
	//按照指定的ids顺序，先查出来，再做匹配,这样查询出来的顺序还是传入的ids的顺序
	tx := db.Clauses(clause.OrderBy{
		Expression: clause.Expr{SQL: "FIELD(post_id,?)", Vars: []interface{}{ids}, WithoutParentheses: true},
	}).Where("post_id in ?", ids).Find(&posts)

	//fmt.Println(posts[0].ID)
	return posts, tx.Error
}

func GetPostListTotalCount(p *models.ParamSearchList) (count int64, err error) {
	var posts []models.Post
	search := "%" + p.Search + "%"
	tx := db.Where("title like ? ", search).Or("content like ?", search).Find(&posts)
	fmt.Println("tx.RowsAffected :", tx.RowsAffected)
	return tx.RowsAffected, tx.Error
}

func GetPostListByKeywords(p *models.ParamSearchList) (posts []*models.Post, err error) {
	posts = make([]*models.Post, 0, 2)
	size := int(p.Size)
	offset := int(p.Page-1) * size
	search := "%" + p.Search + "%"
	tx := db.Where("title like ?", search).Or("content like ?", search).
		Order("create_time desc").
		Limit(size).Offset(offset).Find(&posts)
	return posts, tx.Error
}
