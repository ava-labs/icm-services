package log

import (
	"os"
	"sync/atomic"

	"github.com/ava-labs/avalanchego/utils/logging"
	"go.uber.org/zap"
)

var root atomic.Value

func init() {
	root.Store(logging.NewLogger(
		"default-logger",
		logging.NewWrappedCore(
			logging.Info,
			os.Stdout,
			logging.Plain.ConsoleEncoder(),
		),
	))
}

// SetLevel sets the level of the global logger
func SetLevel(level logging.Level) {
	l := Root()
	l.SetLevel(level)
	root.Store(l)
}

// SetDefault sets the default global logger
func SetDefault(l logging.Logger) {
	root.Store(l)
}

// Root returns the root logger
func Root() logging.Logger {
	return root.Load().(logging.Logger)
}

func Trace(msg string, fields ...zap.Field) {
	Root().Trace(msg, fields...)
}

func Debug(msg string, fields ...zap.Field) {
	Root().Debug(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	Root().Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	Root().Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	Root().Error(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	Root().Fatal(msg, fields...)
	os.Exit(1)
}

func With(ctx ...zap.Field) logging.Logger {
	return Root().With(ctx...)
}
