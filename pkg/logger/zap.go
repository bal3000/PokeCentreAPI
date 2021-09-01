package logger

import (
	"io"
	"runtime"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var encoderCfg = zapcore.EncoderConfig{
	MessageKey: "msg",
	NameKey:    "name",

	LevelKey:    "level",
	EncodeLevel: zapcore.LowercaseLevelEncoder,

	CallerKey:    "caller",
	EncodeCaller: zapcore.ShortCallerEncoder,
}

type ZapLogger struct {
	logger *zap.Logger
}

func NewZapLogger(w io.Writer) (ZapLogger, func()) {
	zl := zap.New(
		zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderCfg),
			zapcore.Lock(zapcore.AddSync(w)),
			zapcore.DebugLevel,
		),
		zap.AddCaller(),
		zap.Fields(zap.String("version", runtime.Version())),
	)

	return ZapLogger{logger: zl}, func() {
		_ = zl.Sync()
	}
}

func (l ZapLogger) LogError(msg string, err error) {
	l.logger.Error(msg, zap.Error(err))
}

func (l ZapLogger) LogDebug(msg string) {
	l.logger.Debug(msg)
}

func (l ZapLogger) LogWarning(msg string) {
	l.logger.Warn(msg)
}

func (l ZapLogger) LogInfo(msg string) {
	l.logger.Info(msg)
}
