package mysql

import (
	"bluebell/setting"
	"database/sql"
	"fmt"
	"time"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var db *gorm.DB
var sqlDB *sql.DB

func Init(cfg *setting.MySQLConfig) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DB,
	)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		zap.L().Error("connect mysql failed", zap.Error(err))
		return
	}

	sqlDB, err = db.DB()

	if err != nil {
		zap.L().Error("connect DB pool failed", zap.Error(err))
		return
	}
	sqlDB.SetMaxIdleConns(cfg.MaxIdleCOnns)
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return
}

func Close() {
	_ = sqlDB.Close()
}
