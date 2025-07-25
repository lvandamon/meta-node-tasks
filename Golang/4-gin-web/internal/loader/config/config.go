package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	App      AppConfig      `yaml:"app" mapstructure:"app"`
	Database DatabaseConfig `yaml:"database" mapstructure:"database"`
	Log      LogConfig      `yaml:"log" mapstructure:"log"`
	Jwt      JwtConfig      `yaml:"jwt" mapstructure:"jwt"`
}

type AppConfig struct {
	Port int `mapstructure:"port"`
}

type DatabaseConfig struct {
	Driver   string `yaml:"driver" mapstructure:"driver"`
	Source   string `yaml:"source" mapstructure:"source"`
	Username string `yaml:"username" mapstructure:"username"`
	Password string `yaml:"password" mapstructure:"password"`
	Host     string `yaml:"host" mapstructure:"host"`
	Port     int    `yaml:"port" mapstructure:"port"`
	Dbname   string `yaml:"dbname" mapstructure:"dbname"`
}

type LogConfig struct {
	Format       string `yaml:"format" mapstructure:"format"`
	Level        string `yaml:"level" mapstructure:"level"`
	ReportCaller bool   `yaml:"reportCaller" mapstructure:"reportCaller"`
}

type JwtConfig struct {
	Secret string `yaml:"secret" mapstructure:"secret"`
}

var Conf *Config

// LoadConfig 加载配置文件
func LoadConfig() error {

	// 设置配置文件路径和名称
	viper.AddConfigPath("./configs")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	// 读取配置文件
	err := viper.ReadInConfig()
	if err != nil {
		return fmt.Errorf("读取配置文件失败: %v", err)
	}

	// 将配置文件内容解析到 Conf 变量中
	Conf = &Config{}
	err = viper.Unmarshal(Conf)
	if err != nil {
		return fmt.Errorf("解析配置文件失败: %v", err)
	}

	return nil
}
