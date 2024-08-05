// Package logger Provides structured & deprecated unstructured logging facade to facilitate logging to different
// providers and multiple instances.
package logger

import (
	"context"
	"github.com/mattn/go-colorable"
	"github.com/natefinch/lumberjack"
	backupLogger "github.com/sirupsen/logrus"
	"go.uber.org/atomic"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

const (
	debugConsoleKey = "debug-console"
	fileKey         = "file"
	jsonStdoutKey   = "json-stdout"

	DefaultAppShortName = "hello-world"
	taskName            = "Logging Service"
)

type LogInstance struct {
	logger  *zap.Logger
	level   zap.AtomicLevel
	enabled *atomic.Bool
}

type Logger struct {
	startMutex sync.RWMutex // locks start/stop
	started    bool
	cfg        atomic.Pointer[loggerConfig]
}

func (s *Logger) config() *loggerConfig {
	cfg := s.cfg.Load()
	if cfg == nil {
		// panic no config loaded
		panic("logger: no config loaded")
	}
	return cfg
}

func (s *Logger) setConfig(cfg *loggerConfig) {
	s.cfg.Store(cfg)
}

var instance *Logger
var once sync.Once

func (s *Logger) StartTask(opts ...LoggingOption) {
	var err error
	s.startMutex.Lock()
	if s.started {
		// if already started do nothing
		s.startMutex.Unlock()
		return
	}

	cfg := s.config().clone()

	// apply options
	for _, opt := range opts {
		opt(cfg.options)
	}

	//
	// debug console logger
	//
	cfg.instances[debugConsoleKey] = &LogInstance{
		enabled: atomic.NewBool(false),
	}
	cfg.instances[debugConsoleKey].level = zap.NewAtomicLevelAt(zapcore.Level(InfoLevel))

	consoleEncoderConfig := zap.NewDevelopmentEncoderConfig()
	//consoleEncoderConfig.FunctionKey = "function"		// uncomment this to enable calling function like: github.com/foo/bar/foo/slogger.(*Logger).ErrorUnstruct
	consoleEncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	consoleEncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000000000 UTCZ07:00")

	cfg.instances[debugConsoleKey].logger = zap.New(
		zapcore.NewCore(
			zapcore.NewConsoleEncoder(consoleEncoderConfig),
			zapcore.AddSync(colorable.NewColorableStdout()),
			cfg.instances[debugConsoleKey].level,
		),
		zap.AddCallerSkip(1),
		zap.Development(),
		zap.AddCaller(),
		zap.AddStacktrace(zapcore.WarnLevel),
	)

	//
	// json stdout logger
	//
	cfg.instances[jsonStdoutKey] = &LogInstance{
		enabled: atomic.NewBool(false),
	}
	cfg.instances[jsonStdoutKey].level = zap.NewAtomicLevelAt(zapcore.Level(InfoLevel))

	jsonStdoutEncoderConfig := zap.NewProductionEncoderConfig()
	jsonStdoutEncoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.UTC().Format(time.RFC3339Nano))
	}
	jsonStdoutLoggerCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(jsonStdoutEncoderConfig),
		zapcore.AddSync(os.Stdout),
		cfg.instances[jsonStdoutKey].level,
	)
	if cfg.options.samplingEnabled {
		jsonStdoutLoggerCore = zapcore.NewSamplerWithOptions(jsonStdoutLoggerCore, cfg.options.samplingOptions.Tick, cfg.options.samplingOptions.First, cfg.options.samplingOptions.Thereafter)
	}
	cfg.instances[jsonStdoutKey].logger = zap.New(jsonStdoutLoggerCore, zap.AddStacktrace(zap.ErrorLevel), zap.AddCaller(), zap.AddCallerSkip(1))

	//
	// file logger
	//
	cfg.instances[fileKey] = &LogInstance{
		enabled: atomic.NewBool(false),
	}
	cfg.instances[fileKey].level = zap.NewAtomicLevelAt(zapcore.Level(ErrorLevel))

	var exPath string
	exPath, err = os.Executable()
	if err != nil {
		backupLogger.Errorf("error creating log file failed to get current working directory: %v", err)
	}
	workingDir := filepath.Dir(exPath)
	logFilePath := filepath.Join(workingDir, cfg.options.productNameShort+".log")
	lumberjackSink := zapcore.AddSync(&lumberjack.Logger{
		Filename:   logFilePath,
		MaxSize:    10, // megabytes
		MaxBackups: 5,
		MaxAge:     28, // days
	})

	fileEncoderConfig := zap.NewProductionEncoderConfig()
	fileEncoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.UTC().Format(time.RFC3339Nano))
	}

	fileLoggerCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(fileEncoderConfig),
		lumberjackSink,
		cfg.instances[fileKey].level,
	)

	if cfg.options.samplingEnabled {
		fileLoggerCore = zapcore.NewSamplerWithOptions(fileLoggerCore, cfg.options.samplingOptions.Tick, cfg.options.samplingOptions.First, cfg.options.samplingOptions.Thereafter)
	}
	cfg.instances[fileKey].logger = zap.New(fileLoggerCore, zap.AddStacktrace(zap.ErrorLevel), zap.AddCaller(), zap.AddCallerSkip(1))

	s.setConfig(cfg)

	s.started = true
	s.startMutex.Unlock()
	s.Info(getTaskLogPrefix(taskName, "started"))
}

func (s *Logger) Sync() {
	cfg := s.config()
	for _, logInstance := range cfg.instances {
		_ = logInstance.logger.Sync()
	}
}

func (s *Logger) StopTask() {
	s.startMutex.Lock()
	if !s.started {
		// if not running do nothing
		s.startMutex.Unlock()
		return
	}
	s.started = false
	s.startMutex.Unlock()
	s.Sync()
	s.Info(getTaskLogPrefix(taskName, "stopped"))
}

// Instance deprecated
func Instance() *Logger {
	once.Do(func() {
		instance = NewLogger()
	})
	return instance
}

func NewLogger() *Logger {
	logger := new(Logger)
	logger.setConfig(&loggerConfig{
		instances: make(map[string]*LogInstance),
		options: &Options{
			productNameShort: DefaultAppShortName,
		},
	})
	return logger
}

func (s *Logger) AddLogger(key string, w io.Writer, newLevel Level, opts ...LoggingOption) {
	cfg := s.config()
	// if a logger already exists at this key do nothing
	_, exists := cfg.instances[key]
	if exists {
		return
	}
	cfg = cfg.clone()

	// apply options
	var addLoggerOpts Options
	for _, opt := range opts {
		opt(&addLoggerOpts)
	}

	cfg.instances[key] = &LogInstance{
		enabled: atomic.NewBool(true),
	}
	cfg.instances[key].level = zap.NewAtomicLevelAt(zapcore.Level(newLevel))

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.UTC().Format(time.RFC3339Nano))
	}

	newloggerCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.AddSync(w),
		cfg.instances[key].level,
	)
	if addLoggerOpts.samplingEnabled {
		newloggerCore = zapcore.NewSamplerWithOptions(newloggerCore, cfg.options.samplingOptions.Tick, cfg.options.samplingOptions.First, cfg.options.samplingOptions.Thereafter)
	}

	cfg.instances[key].logger = zap.New(newloggerCore, zap.AddStacktrace(zap.ErrorLevel), zap.AddCaller(), zap.AddCallerSkip(1))

	s.setConfig(cfg)
}

func (s *Logger) SetLoggerEnabled(key string, enabled bool) {
	cfg := s.config()
	if cfg.instances[key] != nil {
		cfg.instances[key].enabled.Store(enabled)
	}
}

func (s *Logger) SetConsoleLogging(enabled bool) {
	s.SetLoggerEnabled(debugConsoleKey, enabled)
}

func (s *Logger) SetJsonStdoutLogging(enabled bool) {
	s.SetLoggerEnabled(jsonStdoutKey, enabled)
}

func (s *Logger) SetFileLogging(enabled bool) {
	s.SetLoggerEnabled(fileKey, enabled)
}

func (s *Logger) SetFileLogLevel(newLevel Level) {
	cfg := s.config()
	if cfg.instances[fileKey] != nil {
		cfg.instances[fileKey].level.SetLevel(zapcore.Level(newLevel))
	}
}

func (s *Logger) SetConsoleLogLevel(newLevel Level) {
	cfg := s.config()
	if cfg.instances[debugConsoleKey] != nil {
		cfg.instances[debugConsoleKey].level.SetLevel(zapcore.Level(newLevel))
	}
}

func (s *Logger) SetJsonStdoutLogLevel(newLevel Level) {
	cfg := s.config()
	if cfg.instances[jsonStdoutKey] != nil {
		cfg.instances[jsonStdoutKey].level.SetLevel(zapcore.Level(newLevel))
	}
}

func (s *Logger) SetLogLevel(key string, newLevel Level) {
	cfg := s.config()
	if cfg.instances[key] != nil {
		cfg.instances[key].level.SetLevel(zapcore.Level(newLevel))
	}
}

// ErrorInLoggerWriter is used by log Writer sinks added with AddLogger() to log messages to standard console & file instances
// that are enabled so the error in the logger can be trapped somewhere and without an error loop in the logger that triggered it
func (s *Logger) ErrorInLoggerWriter(format string, args ...interface{}) {
	cfg := s.config()
	if cfg.instances[fileKey] != nil && cfg.instances[fileKey].enabled.Load() {
		cfg.instances[fileKey].logger.Sugar().Errorf(format, args...)
	}
	if cfg.instances[debugConsoleKey] != nil && cfg.instances[debugConsoleKey].enabled.Load() {
		cfg.instances[debugConsoleKey].logger.Sugar().Errorf(format, args...)
	}
	if cfg.instances[jsonStdoutKey] != nil && cfg.instances[jsonStdoutKey].enabled.Load() {
		cfg.instances[jsonStdoutKey].logger.Sugar().Errorf(format, args...)
	}
}

func (s *Logger) ErrorInLoggerWriterIgnoreCancel(ctx context.Context, format string, args ...interface{}) {
	if ctx.Err() != nil {
		return
	}
	s.ErrorInLoggerWriter(format, args)
}

func (s *Logger) IsLevelEnabled(level Level) bool {
	cfg := s.config()
	for _, logInstance := range cfg.instances {
		if level >= Level(logInstance.level.Level()) {
			return true
		}
	}
	return false
}

// IsInterruptError returns whether an error was returned by a process that
// was terminated by an interrupt signal (SIGINT).
//func IsInterruptError(err error) bool {
//	exitErr, ok := err.(*exec.ExitError)
//	if !ok || exitErr.ExitCode() >= 0 {
//		return false
//	}
//	status := exitErr.Sys().(syscall.WaitStatus)
//	return status.Signal() == syscall.SIGINT
//}

func fieldsToZapFields(fields ...Field) []zap.Field {
	var zapFields []zap.Field
	for _, f := range fields {
		zapFields = append(zapFields, zap.Field(f))
	}
	return zapFields
}

func getTaskLogPrefix(taskName string, format string) string {
	var sb strings.Builder
	sb.WriteString("[")
	sb.WriteString(taskName)
	sb.WriteString("] ")
	sb.WriteString(format)
	return sb.String()
}
