package initialize

import (
	"github.com/bkcarlos/cliplabel/configs"
	"github.com/bkcarlos/cliplabel/dao"
	"github.com/bkcarlos/cliplabel/logger"
)

// Init 全局 init
func Init() {
	configs.Init()
	logger.InitLogger(configs.Conf.Log.FilePath, configs.Conf.Log.Level, configs.Conf.Server.Name)
	dao.InitSqlLiteDB()
}
