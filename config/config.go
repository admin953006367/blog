package config

import (
	"github.com/BurntSushi/toml"
	"os"
)

type tomlConfig struct {
	Viewer Viewer
	System SystemConfig
}

type Viewer struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Logo        string   `json:"logo"`
	Navigation  []string `json:"navigation"`
	Bilibili    string   `json:"bilibili"`
	Avatar      string   `json:"avatar"`
	UserName    string   `json:"user_name"`
	UserDesc    string   `json:"user_desc"`
}

type SystemConfig struct {
	AppName         string
	Version         float32
	CurrentDir      string
	CdnURL          string
	QiniuAccessKey  string
	QiniuSecretKey  string
	Valine          bool
	ValineAppid     string
	ValineAppkey    string
	ValineServerURL string
}

var Cfg *tomlConfig

func init() {
	//启动的时候，执行init
	Cfg = new(tomlConfig)
	Cfg.System.AppName = "blog"
	Cfg.System.Version = 1.0
	currentDir, _ := os.Getwd()
	Cfg.System.CurrentDir = currentDir
	_, err := toml.DecodeFile("config/config.toml", &Cfg)
	if err != nil {
		//异常的话直接关闭
		panic(err)
	}
}
