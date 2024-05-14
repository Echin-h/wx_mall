package log

import (
	"go.uber.org/zap/zapcore"
	"time"
	"wxapp/config"
)

type Zap struct {
	Prefix       string `default:"gin-vue-admin" yaml:"prefix" json:"prefix"`
	Level        string `default:"panic" yaml:"level" json:"level"`
	Path         string `yaml:"path" json:"path"`
	Director     string `default:"director" yaml:"director" json:"director"`
	EncoderLevel string `default:"LowercaseLevelEncoder" yaml:"encoder_level" json:"encoder_level"`
	MaxAge       int    `default:"7" yaml:"max_age" json:"max_age"`
	ShowLine     bool   `default:"true" yaml:"show_line" json:"show_line"`
	LogInConsole bool   `default:"true" yaml:"log_in_console" json:"log_in_console" `
	Encoder      string `default:"console" yaml:"encoder" json:"encoder"`
}

func GetLevelEncoder() zapcore.LevelEncoder {
	lvl := config.GetConfig().Logger.EncoderLevel
	switch lvl {
	case "LowercaseLevelEncoder":
		return zapcore.LowercaseColorLevelEncoder
	case "LowercaseColorLevelEncoder":
		return zapcore.LowercaseColorLevelEncoder
	case "CapitalLevelEncoder":
		return zapcore.CapitalLevelEncoder
	case "CapitalColorLevelEncoder":
		return zapcore.CapitalColorLevelEncoder
	default:
		return zapcore.LowercaseLevelEncoder
	}
}

func GetTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	prefix := config.GetConfig().Logger.Prefix
	t.Format(prefix + ":" + "2006-01-02 15:04:05")
}

func LevelEnable(level string) zapcore.Level {
	switch level {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	case "dpanic":
		return zapcore.DPanicLevel
	case "panic":
		return zapcore.PanicLevel
	case "fatal":
		return zapcore.FatalLevel
	default:
		return zapcore.InfoLevel
	}
}
