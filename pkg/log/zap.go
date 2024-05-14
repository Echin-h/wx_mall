package log

import (
	"go.uber.org/zap/zapcore"
	"os"
	"wxapp/config"
)

var z = &_zap{}

type _zap struct{}

func (z *_zap) New() {

}

func (z *_zap) NewEncoderCore(level zapcore.Level) {
	writer := z.GenWriter()
	zapcore.NewCore(z.GenEncoder(), writer, level)
}

func (z *_zap) GenEncoder() zapcore.Encoder {
	if config.GetConfig().Logger.Encoder == "json" {
		return zapcore.NewJSONEncoder(z.GetConfig())
	}
	return zapcore.NewConsoleEncoder(z.GetConfig())
}

func (z *_zap) GenWriter() zapcore.WriteSyncer {
	return zapcore.AddSync(os.Stdout)
}

func (z *_zap) GetConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  "stacktrace",
		SkipLineEnding: true,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    GetLevelEncoder(),
		EncodeTime:     GetTimeEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}
}
