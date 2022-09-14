package initialize

import (
	"github.com/bkcarlos/cliplabel/configs"
	"github.com/bkcarlos/cliplabel/dao"
	"github.com/bkcarlos/cliplabel/logger"
	"github.com/bkcarlos/cliplabel/models"
)

// Init 全局 init
func Init() {
	// 读取配置文件
	configs.Init()
	// 初始化日志文件
	logger.InitLogger(configs.Conf.Log.FilePath, configs.Conf.Log.Level, configs.Conf.Server.Name)
	// 初始化数据库链接
	dao.Init()
	// 初始化表
	models.Init()
}
