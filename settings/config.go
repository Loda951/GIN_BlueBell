package settings

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Config = new(AppConfig)

type AppConfig struct {
	Name         string `mapstructure:"name"`
	Mode         string `mapstructure:"mode"`
	Version      string `mapstructure:"version"`
	Port         int    `mapstructure:"port"`
	*LogConfig   `mapstructure:"log"`
	*MySQLConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	FileName   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

type MySQLConfig struct {
	Host         string `mapstructure:"host"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	DBname       string `mapstructure:"dbname"`
	Port         int    `mapstructure:"port"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Password string `mapstructure:"password"`
	Port     int    `mapstructure:"port"`
	DB       int    `mapstructure:"db"`
	PoolSize int    `mapstructure:"pool_size"`
}

func InitConfig() error {
	viper.SetConfigName("config")     // 配置文件名称(无扩展名)
	viper.SetConfigType("yaml")       // 指定配置文件类型(专用于从远程获取配置信息时指定配置文件类型)
	viper.AddConfigPath("./settings") // 还可以在工作目录中查找配置
	err := viper.ReadInConfig()       // 查找并读取配置文件
	if err != nil {                   // 处理读取配置文件的错误
		fmt.Printf("viper.ReadInConfig() failed err:%v\n", err)
		return err
	}
	if err := viper.Unmarshal(Config); err != nil {
		fmt.Printf("viper.Unmarshal() failed err:%v\n", err)
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("settings file changed:", e.Name)
		if err := viper.Unmarshal(Config); err != nil {
			fmt.Printf("viper.Unmarshal() failed err:%v\n", err)
		}
	})
	return nil
}
