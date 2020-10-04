package log

import (
	"io"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func newZap(routineID string, writer io.Writer) *zap.SugaredLogger {
	conf := zap.NewDevelopmentEncoderConfig()
	conf.EncodeTime = zapcore.RFC3339TimeEncoder
	conf.EncodeLevel = zapcore.CapitalLevelEncoder

	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(conf),
		zapcore.AddSync(writer),
		zapcore.DebugLevel,
	)

	logger := zap.New(core, zap.AddCaller()).Named(routineID)
	zap.RedirectStdLog(logger)

	return logger.Sugar()
}
