package dao

import (
	"database/sql"
	"os"
	"path"

	"github.com/bkcarlos/cliplabel/configs"
	"github.com/bkcarlos/cliplabel/logger"
	"github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func getLibSimplePath() string {
	p := os.Getenv("CLIP_LABEL_SIMPLE")
	if p == "" {
		// 运行环境有问题，不要启动
		ePath, err := os.Executable()
		if err != nil {
			panic(err)
		}

		p = path.Dir(ePath)
	}

	return path.Join(p, "lib")
}

// GetSqlitedb 获取数据库连接
//  @param filepath
//  @return *sql.DB
//  @return error
func GetSqlitedb(filepath string) (*sql.DB, error) {
	libPath := getLibSimplePath()
	libSimpleFile := path.Join(libPath, "libsimple", "libsimple")
	sql.Register("sqlite3_simple",
		&sqlite3.SQLiteDriver{
			Extensions: []string{
				libSimpleFile,
			},
		})

	return sql.Open("sqlite3_simple", filepath)
}

// Init 初始化
func InitSqlLiteDB() {
	var err error
	dbPath := path.Join(configs.Conf.Sqlite.DBPath, configs.Conf.Sqlite.DBName)
	libPath := getLibSimplePath()
	libSimpleFile := path.Join(libPath, "libsimple", "libsimple")
	const CustomDriverName = "sqlite3_simple"
	sql.Register("sqlite3_simple",
		&sqlite3.SQLiteDriver{
			Extensions: []string{
				libSimpleFile,
			},
		})

	gromLogger := logger.NewGormLogger(logger.GetLogLevel())
	ClipLabelDB, err = gorm.Open(&sqlite.Dialector{
		DriverName: CustomDriverName,
		DSN:        dbPath,
	}, &gorm.Config{
		Logger: gromLogger,
	})
	if err != nil {
		logger.Errors("init db is err:%v", err)
		panic(err)
	}
}
