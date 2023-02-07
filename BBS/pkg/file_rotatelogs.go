package pkg

import (
	"github.com/robot007num/go/bbs/global"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"path"
	"time"
)

// GetWriteSyncer 获取 zapcore.WriteSyncer
// 使用lumberjack来分割
func GetWriteSyncer(level string) zapcore.WriteSyncer {
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   path.Join(global.GVA_CONFIG.Zap.Director, time.Now().Format("2006-01-02"), level+".log"),
		MaxSize:    global.GVA_CONFIG.Zap.Size,
		MaxBackups: global.GVA_CONFIG.Zap.BakeUp,
		MaxAge:     global.GVA_CONFIG.Zap.MaxAge,
		Compress:   global.GVA_CONFIG.Zap.Compress,
	})

	if global.GVA_CONFIG.Zap.LogInConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(w))
	}

	return w
}
