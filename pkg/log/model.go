package log

import (
	"go.uber.org/zap/zapcore"
	"wxapp/config"
)

func GetLevelEncoder() zapcore.LevelEncoder {
	lvl := config.GetConfig().Zap.EncoderLevel
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

func TransmitLvl() zapcore.Level {
	Z := config.GetConfig().Zap
	switch Z.Level {
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
