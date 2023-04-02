package main

import (
	"fmt"
	"gin_forum/conf"
	"gin_forum/dao/mysql"
	"gin_forum/dao/redis"
	"gin_forum/logger"
	"gin_forum/pkg/snowflake"
	"gin_forum/router"

	"go.uber.org/zap"
)

func main() {
	//1. 加载配置
	if err := conf.Init("conf/config.yaml"); err != nil {
		fmt.Printf("init setting failed,err :v%\n", err)
		return
	}
	//2. 初始化日志
	if err := logger.Init(conf.Conf.LogConfig, conf.Conf.Mode); err != nil {
		fmt.Printf("init logger setting failed, err:%v\n", err)
		return
	}
	defer zap.L().Sync()
	zap.L().Debug("logger init success...")
	//3. 初始化Mysql连接
	if err := mysql.Init(conf.Conf.MySQLConfig); err != nil {
		fmt.Printf("init mysql setting failed, err:%v\n", err)
		return
	}
	defer mysql.Close()
	//4. 初始化Redis连接

	if err := redis.Init(conf.Conf.RedisConfig); err != nil {
		fmt.Printf("init redis setting failed, err:%v\n", err)
		return
	} else {
		fmt.Println("init redis setting")
	}

	if err := snowflake.Init(conf.Conf.SnowflakeStartTime, conf.Conf.SnowflakeMachineID); err != nil {
		fmt.Printf("init snowflake failed, err:%v\n", err)
		return
	}
	r := router.SetupRouter(conf.Conf.Mode)
	err := r.Run(fmt.Sprintf(":%d", conf.Conf.Port))
	if err != nil {
		fmt.Printf("server run failedm err,err:%s", err)
	}

}
