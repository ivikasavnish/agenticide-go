package logger

import (
"os"

"go.uber.org/zap"
"go.uber.org/zap/zapcore"
)

type Logger struct {
*zap.Logger
sugar *zap.SugaredLogger
}

func NewLogger(level string) (*Logger, error) {
zapLevel := zapcore.InfoLevel
switch level {
case "debug":
zapLevel = zapcore.DebugLevel
case "info":
zapLevel = zapcore.InfoLevel
case "warn":
zapLevel = zapcore.WarnLevel
case "error":
zapLevel = zapcore.ErrorLevel
}

encoderConfig := zapcore.EncoderConfig{
TimeKey:        "time",
LevelKey:       "level",
NameKey:        "logger",
CallerKey:      "caller",
MessageKey:     "msg",
StacktraceKey:  "stacktrace",
LineEnding:     zapcore.DefaultLineEnding,
EncodeLevel:    zapcore.CapitalColorLevelEncoder,
EncodeTime:     zapcore.ISO8601TimeEncoder,
EncodeDuration: zapcore.StringDurationEncoder,
EncodeCaller:   zapcore.ShortCallerEncoder,
}

core := zapcore.NewCore(
zapcore.NewConsoleEncoder(encoderConfig),
zapcore.AddSync(os.Stdout),
zapLevel,
)

logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))

return &Logger{
Logger: logger,
sugar:  logger.Sugar(),
}, nil
}

func (l *Logger) Named(name string) *Logger {
return &Logger{
Logger: l.Logger.Named(name),
sugar:  l.sugar.Named(name),
}
}

func (l *Logger) Sugar() *zap.SugaredLogger {
return l.sugar
}

func (l *Logger) Sync() error {
return l.Logger.Sync()
}
