package settings

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
	"os"
)

//mapstructure 标签用于指定在解析配置文件时，这些字段对应的键名。例如，配置文件中的 host 键会映射到 Host 字段

var Conf = new(AppConfig)

type AppConfig struct {
	Mode         string `mapstructure:"mode"`
	Port         int    `mapstructure:"port"`
	Name         string `mapstructure:"name"`
	Version      string `mapstructure:"version"`
	StartTime    string `mapstructure:"start_time"`
	MachineID    int    `mapstructure:"machine_id"`
	*LogConfig   `mapstructure:"log"`
	*MySQLConfig `mapstructure:"mysql"`
	//todo
	//*RedisConfig `mapstructure:"redis"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"` //日志文件的最大大小（以MB为单位）。当日志文件达到这个大小时，系统将自动创建一个新的日志文件。
	MaxLine    int    `mapstructure:"max_line"`
	MaxAge     int    `mapstructure:"max_age"`     //日志文件的最大保留时间（以天为单位）。超过这个时间的日志文件将被自动删除。
	MaxBackups int    `mapstructure:"max_backups"` //MaxBackups (int): 日志文件的最大备份数量。当日志文件数量超过这个数量时，最旧的日志文件将被删除。
}

type MySQLConfig struct {
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	Dbname       string `mapstructure:"dbname"`
	MaxOpenConns int    `mapstructure:"max_open_conns"` //数据库连接池中最大打开的连接数。
	MaxIdleConns int    `mapstructure:"max_idle_conns"` //数据库连接池中最大空闲的连接数。
}

func generateConfig() {
	config := AppConfig{
		Mode:      "localhost",
		Port:      8088,
		Name:      "MyApp",
		Version:   "1.0.0",
		StartTime: "2023-10-01T00:00:00Z",
		MachineID: 1,
		LogConfig: &LogConfig{
			Level:      "info",
			Filename:   "app.log",
			MaxSize:    10,
			MaxAge:     7,
			MaxBackups: 3,
		},
		MySQLConfig: &MySQLConfig{
			Host:         "localhost",
			User:         "root",
			Password:     "password",
			Dbname:       "yiyuser",
			Port:         3306,
			MaxOpenConns: 100,
			MaxIdleConns: 10,
		},
	}

	// 创建 conf 目录
	err := os.MkdirAll("./conf", os.ModePerm)
	if err != nil {
		fmt.Printf("Failed to create conf directory: %v\n", err)
		return
	}

	file, err := os.Create("./conf/config.yaml")
	if err != nil {
		fmt.Printf("Failed to create config file: %v\n", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	data, err := yaml.Marshal(&config)
	if err != nil {
		fmt.Printf("Failed to marshal config: %v\n", err)
		return
	}

	_, err = file.Write(data)
	if err != nil {
		fmt.Printf("Failed to write to config file: %v\n", err)
		return
	}

	fmt.Println("Config file generated successfully.")
}

func Init() error {

	// 生成配置文件
	generateConfig()
	//读取配置文件
	viper.SetConfigFile("./conf/config.yaml")
	//读取环境变量
	viper.WatchConfig()
	//监听配置文件变化
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件发生变更")
		err := viper.Unmarshal(&Conf)
		if err != nil {
			return
		}
	})
	//查找并读取配置文件
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("配置文件读取失败，err: %v", err))

	}
	//把读取到的配置信息反序列化到conf变量中
	if err := viper.Unmarshal(Conf); err != nil {

		panic(fmt.Errorf("配置文件解析失败，err: %v", err))
	}
	return err
}
