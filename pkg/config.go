package pkg

import (
	"fmt"
	"github.com/spf13/viper"
)

/*type ChaogeConf struct {
	AllConfig Config `yaml:"allConfig"`
}*/

type ConfigS struct {
	GithubToken string `yaml:"githubToken"`
	DingDing    Ding   `yaml:"dingding"`
	Feishu      Fei    `yaml:"feishu"`
}

type Ding struct {
	Enable    string `yaml:"enable"`
	Webhook   string `yaml:"webhook"`
	SecretKey string `yaml:"secretKey"`
	AppName   string `yaml:"app_name"`
}

type Fei struct {
	Enable  string `yaml:"enable"`
	Webhook string `yaml:"webhook"`
	AppName string `yaml:"app_name"`
}

func SettingS() {
	var dsb = ConfigS{}
	vip := viper.New()
	vip.SetConfigFile("./config/config.yaml")
	vip.SetConfigType("yaml")

	if err := vip.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := vip.Unmarshal(&dsb); err != nil {
		panic(err)
	}
	fmt.Println(dsb)
	fmt.Println(dsb.Feishu.AppName)
	fmt.Println(dsb.DingDing.AppName)
}
