package util

import (
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"
)

// Config 配置文件定义
type Config struct {
	MysqlDsn      string `yaml:"MYSQL_DSN"`
	SessionSecret string `yaml:"SESSION_SECRET"`
	GinMode       string `yaml:"GIN_MODE"`
	LogLevel      string `yaml:"LOG_LEVEL"`
	AdminPassword string `yaml:"ADMIN_PASSWORD"`
}

// ReadYamlConfig 读取yaml配置
func ReadYamlConfig(path string) (*Config, error) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		Log().Error("获取程序执行目录出错：" + err.Error())
	}
	if path == "" {
		path = strings.Replace(dir, "\\", "/", -1)
	}
	conf := &Config{}
	f, err := os.Open(path + "/config.yaml")
	if err != nil {
		return nil, err
	}
	defer f.Close()
	_ = yaml.NewDecoder(f).Decode(conf)

	return conf, nil
}
