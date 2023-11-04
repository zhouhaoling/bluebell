package mysql

import "bluebell/models"

func CreatePost(p *models.Post) error {
	result := db.Create(&p)
	return result.Error
}

func GetPostById(pid int64) (post *models.Post, err error) {
	post = new(models.Post)
	result := db.Where("post_id = ?", pid).First(post)
	return post, result.Error
}

func GetPostList(page, size int) (posts []*models.Post, err error) {
	posts = make([]*models.Post, 0, 2)
	result := db.Limit(size).Offset((page - 1) * size).Find(&posts)
	err = result.Error
	return
}
