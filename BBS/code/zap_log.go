package code

import (
	"fmt"
	"github.com/robot007num/go/bbs/global"
	"github.com/robot007num/go/bbs/pkg/log"
	"github.com/robot007num/go/bbs/utils"
	"os"
)

func Zap() *log.Logger {
	if ok, _ := utils.PathExists(global.GVA_CONFIG.Zap.Director); !ok { //判断是否有Director文件夹
		fmt.Printf("create %v directory\n", global.GVA_CONFIG.Zap.Director)
		_ = os.Mkdir(global.GVA_CONFIG.Zap.Director, os.ModePerm)
	}

	var tops = []log.TeeOption{
		{
			Filename: "",
			Ropt: log.RotateOptions{
				MaxSize:    1,
				MaxAge:     1,
				MaxBackups: 3,
				Compress:   true,
			},
			Lef: func(lvl log.Level) bool {
				return lvl <= log.InfoLevel
			},
		},
		{
			Filename: "",
			Ropt: log.RotateOptions{
				MaxSize:    1,
				MaxAge:     1,
				MaxBackups: 3,
				Compress:   true,
			},
			Lef: func(lvl log.Level) bool {
				return lvl > log.InfoLevel
			},
		},
	}

	logger := log.NewTeeWithRotate(tops)
	log.ResetDefault(logger)

	return logger
}
