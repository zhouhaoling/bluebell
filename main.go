package main

import (
	"bluebell/dao/mysql"
	"bluebell/dao/redis"
	"bluebell/logger"
	"bluebell/pkg/snowflake"
	"bluebell/router"
	"bluebell/setting"
	"fmt"

	"go.uber.org/zap"
)

func main() {
	//1.加载配置文件（远程加载）
	if err := setting.Init(); err != nil {
		fmt.Println("init settings failed, err:", err)
		return
	}
	//2.初始日志
	if err := logger.Init(setting.Conf.LogConfig, setting.Conf.Mode); err != nil {
		fmt.Println("init logger failed, err:", err)
		return
	}
	defer zap.L().Sync()
	zap.L().Debug("logger init success...")

	//3.初始化MySQL连接
	if err := mysql.Init(setting.Conf.MySQLConfig); err != nil {
		fmt.Println("init mysql failed, err:", err)
		return
	}
	defer mysql.Close()
	//4.初始化Redis连接
	if err := redis.Init(setting.Conf.RedisConfig); err != nil {
		fmt.Println("init redis failed, err:", err)
		return
	}
	defer redis.Close()

	if err := snowflake.Init(setting.Conf.StartTime, setting.Conf.MachineID); err != nil {
		fmt.Println("init snowflake failed, err:", err)
		return
	}
	//5.注册路由
	r := router.SetUpRoutes()
	err := r.Run(fmt.Sprintf(":%d", setting.Conf.Port))
	if err != nil {
		fmt.Printf("run logic failed, err:%v\n", err)
		return
	}
}
