package code

import (
	"fmt"
	"github.com/robot007num/go/bbs/global"
	"github.com/robot007num/go/bbs/pkg/log"
	"github.com/robot007num/go/bbs/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func Zap() (MyLog *zap.Logger) {
	if ok, _ := utils.PathExists(global.GVA_CONFIG.Zap.Director); !ok { //判断是否有Director文件夹
		fmt.Printf("create %v directory\n", global.GVA_CONFIG.Zap.Director)
		_ = os.Mkdir(global.GVA_CONFIG.Zap.Director, os.ModePerm)
	}

	cores := log.GetZapCores()
	MyLog = zap.New(zapcore.NewTee(cores...))

	if global.GVA_CONFIG.Zap.ShowLine {
		MyLog = MyLog.WithOptions(zap.AddCaller())
	}

	return MyLog
}
