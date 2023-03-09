package tool

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

/*
AppConfig 配置
*/
type AppConfig struct {
	GithubToken      string         `json:"githubToken"`
	DingDing         Ding           `json:"dingDing"`
	EnterpriseWeChat WeChat         `json:"enterpriseWeChat"`
	Database         DatabaseConfig `json:"database"`
}

type Ding struct {
	Enable    bool     `json:"enable"`
	Webhook   string   `json:"webhook"`
	SecretKey string   `json:"secretKey"`
	Mobiles   []string `json:"mobiles"`
}

type WeChat struct {
	Enable         bool     `json:"enable"`
	Webhook        string   `json:"webhook"`
	MentionUsers   []string `json:"mentionUsers"`
	MentionMobiles []string `json:"mentionMobiles"`
}

type DatabaseConfig struct {
	Driver   string `json:"driver"`
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	DbName   string `json:"dbname"`
	Charset  string `json:"charset"`
	ShowSql  bool   `json:"show_sql"`
}

var appConf AppConfig

func GetAppConfig() *AppConfig {
	return &appConf
}

/*
ToolsConf 配置
*/
type ToolsConf struct {
	ToolsList   []string `yaml:"toolsList"`
	KeywordList []string `yaml:"keywordList"`
	UserList    []string `yaml:"userList"`
	BlackUser   []string `yaml:"blackUser"`
}

var toolConf ToolsConf

func GetToolsConf() *ToolsConf {
	return &toolConf
}

// ParseAppConfig 解析 app.json 配置文件
func ParseAppConfig(filepath string) *AppConfig {
	viper.SetConfigFile(filepath)
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error tool.json file: %s \n", err))
	}

	// 监控配置文件变化
	viper.WatchConfig()
	// 注意！！！配置文件发生变化后要同步到全局变量Conf
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("夭寿啦~ app配置文件被人修改啦...")
		if err := viper.Unmarshal(&appConf); err != nil {
			panic(fmt.Errorf("unmarshal conf failed, err:%s \n", err))
		}
	})

	// 将读取的配置信息保存至全局变量Conf
	if err := viper.Unmarshal(&appConf); err != nil {
		panic(fmt.Errorf("unmarshal conf failed, err:%s \n", err))
	}
	//fmt.Println(appConf)
	return &appConf
}

// ParseToolConf 解析 tool.json 配置文件
func ParseToolConf(filePath string) *ToolsConf {
	viper.SetConfigFile(filePath)
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error tool.json file: %s \n", err))
	}

	// 监控配置文件变化
	viper.WatchConfig()
	// 注意！！！配置文件发生变化后要同步到全局变量Conf
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("夭寿啦~ tool配置文件被人修改啦...")
		if err := viper.Unmarshal(&toolConf); err != nil {
			panic(fmt.Errorf("unmarshal conf failed, err:%s \n", err))
		}
	})

	// 将读取的配置信息保存至全局变量 toolConf
	if err := viper.Unmarshal(&toolConf); err != nil {
		panic(fmt.Errorf("unmarshal conf failed, err:%s \n", err))
	}
	return &toolConf
}
