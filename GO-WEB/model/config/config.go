package config

type AllConfig struct {
	App
	MySql
	Redis
	Log
}

var allConfig AllConfig

type App struct {
	Name    string `mapstructure:"name"`
	Version string `mapstructure:"version"`
}

type MySql struct {
	User        string `mapstructure:"user"`
	PassWord    string `mapstructure:"password"`
	Host        string `mapstructure:"host"`
	Port        int    `mapstructure:"port"`
	DbName      string `mapstructure:"dbname"`
	Max_Connect int    `mapstructure:"max_con"`
}

type Redis struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}

type Log struct {
	ComPress     bool   `mapstructure:"compress"`
	InfoFileName string `mapstructure:"info_filename"`
	ErrFileName  string `mapstructure:"err_filename"`
	GinFileName  string `mapstructure:"gin_filename"`
	LogMaxSize   int    `mapstructure:"max_size"`
	LogSaveDay   int    `mapstructure:"max_age"`
	LogBackups   int    `mapstructure:"max_backups"`
}

// InitAllConfig 初始化allConfig
func (a AllConfig) InitAllConfig() {
	allConfig = a
}

//GetAllConfig 获取allConfig(让外部无法修改里面的内容)
func GetAllConfig() AllConfig {
	return allConfig
}
