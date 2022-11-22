package app

import (
	_ "gorm-example/api" //注册所有controller
	"gorm-example/common"
	"gorm-example/config"
)

func Run() {
	config.MysqlInit()
	_ = common.R.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
