package log

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"wxapp/config"
	"wxapp/pkg/fs"
)

var Logger *zap.Logger

func LoadLog() {
	path := config.GetConfig().Zap.Director
	if err := fs.IsNotExistMkDir(path); err != nil {
		panic(fmt.Sprintf("create log dir failed: %v\n", err))
	}
	if config.GetConfig().Zap.LogInConsole {

	}
	logf := z.GetLogger()
	Logger = logf
	fmt.Println("loading log success...")
}

func (z *_zap) GetLogger() *zap.Logger {
	var log *zap.Logger
	if config.GetConfig().Zap.LogInConsole {
		core := zapcore.NewCore(z.GetEncoder(), os.Stdout, zap.DebugLevel)
		log = zap.New(core, zap.AddCaller())
	} else {
		cores := z.GetZapCores()
		log = zap.New(zapcore.NewTee(cores...), zap.AddCaller())
	}
	return log
}
