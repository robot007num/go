package config

type AllConfig struct {
	App
	MySql
}

var allConfig AllConfig

type App struct {
	Name    string `mapstructure:"name"`
	Version string `mapstructure:"version"`
}

type MySql struct {
	User     string `mapstructure:"user"`
	PassWord string `mapstructure:"password"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	DbName   string `mapstructure:"dbname"`
}

// InitAllConfig 初始化allConfig
func (a AllConfig) InitAllConfig() {
	allConfig = a
}

//GetAllConfig 获取allConfig(让外部无法修改里面的内容)
func GetAllConfig() AllConfig {
	return allConfig
}
