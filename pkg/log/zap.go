package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"os"
	"time"
	"wxapp/config"
	"wxapp/pkg/fs"
)

var z = new(_zap)

type _zap struct{}

func (z *_zap) GetZapCores() []zapcore.Core {
	cores := make([]zapcore.Core, 0, 7)
	for level := TransmitLvl(); level <= zap.FatalLevel; level++ {
		cores = append(cores, z.NewEncoderCore(level, z.GetLevelPriority(level)))
	}
	return cores
}

func (z *_zap) GetTimeEncoder(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
	prefix := config.GetConfig().Zap.Prefix
	encoder.AppendString(prefix + " :" + t.Format("2006/01/02 - 15:04:05"))
}

func (z *_zap) NewEncoderCore(l zapcore.Level, level zap.LevelEnablerFunc) zapcore.Core {
	writer := z.GenWriter(l.String())
	return zapcore.NewCore(z.GetEncoder(), writer, level)
}

func (z *_zap) GetEncoder() zapcore.Encoder {
	if config.GetConfig().Zap.Encoder == "json" {
		return zapcore.NewJSONEncoder(z.GetConfig())
	}
	return zapcore.NewConsoleEncoder(z.GetConfig())
}

func (z *_zap) GenWriter(level string) zapcore.WriteSyncer {
	dir := config.GetConfig().Zap.Director
	path := dir + "/" + level + ".log"
	_ = fs.IsNotExistMkFile(path)
	log.Printf("log init success")
	maxSize := config.GetConfig().Zap.MaxSize
	maxBackups := config.GetConfig().Zap.MaxBackups
	maxAge := config.GetConfig().Zap.MaxAge
	comp := config.GetConfig().Zap.Compress

	lumberJack := &lumberjack.Logger{
		Filename:   path,
		MaxSize:    maxSize,
		MaxBackups: maxBackups,
		MaxAge:     maxAge,
		Compress:   comp,
	}

	writeSyncer := zapcore.AddSync(lumberJack)
	if config.GetConfig().Zap.LogInConsole {
		writeSyncer = zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), writeSyncer)
	}

	log.Println("writeSyncer init success")
	return writeSyncer
}

func (z *_zap) GetConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		SkipLineEnding: false,
		EncodeLevel:    GetLevelEncoder(),
		EncodeTime:     z.GetTimeEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}
}

func (z *_zap) GetLevelPriority(level zapcore.Level) zap.LevelEnablerFunc {
	switch level {
	case zapcore.DebugLevel:
		return func(level zapcore.Level) bool {
			return level == zap.DebugLevel
		}
	case zapcore.InfoLevel:
		return func(level zapcore.Level) bool {
			return level == zap.InfoLevel
		}
	case zapcore.WarnLevel:
		return func(level zapcore.Level) bool {
			return level == zap.WarnLevel
		}
	case zapcore.ErrorLevel:
		return func(level zapcore.Level) bool {
			return level == zap.ErrorLevel
		}
	case zapcore.DPanicLevel:
		return func(level zapcore.Level) bool {
			return level == zap.DPanicLevel
		}
	case zapcore.PanicLevel:
		return func(level zapcore.Level) bool {
			return level == zap.PanicLevel
		}
	case zapcore.FatalLevel:
		return func(level zapcore.Level) bool {
			return level == zap.FatalLevel
		}
	default:
		return func(level zapcore.Level) bool {
			return level == zap.DebugLevel
		}
	}
}
