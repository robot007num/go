package code

import (
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/robot007num/go/bbs/code/internal"
	"github.com/robot007num/go/bbs/global"
	"github.com/spf13/viper"
	"os"
)

//Viper
// 优先级: 命令行 > 环境变量 > 默认值
func Viper(path ...string) *viper.Viper {
	var config string
	pLen := len(path)

	if pLen == 0 {
		//命令行解析
		flag.StringVar(&config, "configFile", "", "path")
		flag.Parse()

		if config == "" { //命令行参数为空
			configEnv := os.Getenv(internal.ConfigEnv)

			if configEnv == "" {
				switch gin.Mode() {
				case gin.DebugMode:
					config = internal.ConfigDefaultFile
					fmt.Printf("您正在使用gin模式的%s环境名称,config的路径为%s\n", gin.EnvGinMode, internal.ConfigDefaultFile)
				case gin.ReleaseMode:
					config = internal.ConfigReleaseFile
					fmt.Printf("您正在使用gin模式的%s环境名称,config的路径为%s\n", gin.EnvGinMode, internal.ConfigReleaseFile)
				case gin.TestMode:
					config = internal.ConfigTestFile
					fmt.Printf("您正在使用gin模式的%s环境名称,config的路径为%s\n", gin.EnvGinMode, internal.ConfigTestFile)
				} //switch
			}

			if configEnv != "" { // internal.ConfigEnv 常量存储的环境变量不为空 将值赋值于config
				config = configEnv
				fmt.Printf("您正在使用%s环境变量,config的路径为%s\n", internal.ConfigEnv, config)
			}

		} else { //命令行参数不为空
			fmt.Printf("您正在使用命令行的-c参数传递的值,config的路径为%s\n", config)
		}

	} // if pLen==0

	if pLen > 0 {

	}

	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType(internal.ConfigFileType)
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err = v.Unmarshal(&global.GVA_CONFIG); err != nil {
			fmt.Println(err)
		}
	})

	if err = v.Unmarshal(&global.GVA_CONFIG); err != nil {
		fmt.Println(err)
	}

	return v
}
