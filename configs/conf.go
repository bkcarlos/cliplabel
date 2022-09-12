package configs

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"gopkg.in/yaml.v2"
)

// Conf 全局配置
var Conf CConfigs

// TConfigs config 配置
type CConfigs struct {
	Server CServer `yaml:"server"` // 服务器相关的配置
	Sqlite CSqlite `yaml:"sqlite"` // sqlite 数据库相关的配置
	Files  CFileS  `yaml:"files"`  // 文件管理相关的配置
	Log    CLog    `yaml:"log"`    // 日志相关的配置
}

// TServe 服务器配置
type CServer struct {
	Port string `yaml:"port"` // 服务运行的端口
	Name string `yaml:"name"` // 服务名
}

// CSqlite sqlite 的配置
type CSqlite struct {
	DBName string `yaml:"db_name"` // 数据库名字
	DBPath string `yaml:"db_path"` // 数据保存的位置
}

// CFileS 文件保存相关的设置，默认保存位置，数据的分片，分区，跨分区
type CFileS struct {
	CopyIn   bool   `yaml:"copy_in"`   // 是否将文件拷贝到程序目录下
	CopyPath string `yaml:"copy_path"` // 文件的根目录
	IsSync   bool   `yaml:"is_sync"`   // 是否同步到远程保存
	SyncTyps string `yaml:"sync_type"` // 同步的方式，百度云，阿里云
}

// CLog Log 日志配置
type CLog struct {
	Level    string `yaml:"level"`
	FilePath string `yaml:"file_path"`
}

// getpath 获取配置文件夹
//  @return string
func getpath() string {
	p := os.Getenv("CLIP_LABEL_PATH")
	if p == "" {
		// 运行环境有问题，不要启动
		ePath, err := os.Executable()
		if err != nil {
			panic(err)
		}

		return path.Dir(ePath)
	}
	return p
}

// Init 初始化配置
func Init() {
	// 配置文件默认放到当前运行的目录下
	p := getpath()
	configFilePath := path.Join(p, "config.yaml")

	fmt.Printf("used config file path:%v\n", configFilePath)

	yamFile, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		panic(err)
	}

	fmt.Printf("config:%v\n", string(yamFile))
	err = yaml.Unmarshal(yamFile, &Conf)
	if err != nil {
		panic(fmt.Errorf("unmarshal config is err:%v", err))
	}
}
