package redis

import (
	"strconv"
	"time"
)

// var ctx = context.Background()
const data = 2 * time.Hour

// SetToken 插入token
func SetToken(userID int64, token string) (err error) {
	err = rdb.Set(strconv.FormatInt(userID, 10), token, data).Err()
	if err != nil {
		return err
	}
	return
}

// GetTokenExist 查找token
func GetTokenExist(userID int64, token string) bool {
	rtoken, err := rdb.Get(strconv.FormatInt(userID, 10)).Result()
	if err != nil {
		return false
	}
	if token != rtoken {
		return false
	}
	return true
}
