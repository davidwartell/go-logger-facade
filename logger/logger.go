package logger

import (
	"context"
	"fmt"
	"os"
)

func (s *Logger) Trace(msg string, fields ...Field) {
	cfg := s.config()
	for _, logInstance := range cfg.instances {
		if logInstance.enabled.Load() {
			logInstance.logger.Debug(msg, fieldsToZapFields(fields...)...)
		}
	}
}

func (s *Logger) Debug(msg string, fields ...Field) {
	cfg := s.config()
	for _, logInstance := range cfg.instances {
		if logInstance.enabled.Load() {
			logInstance.logger.Debug(msg, fieldsToZapFields(fields...)...)
		}
	}
}

func (s *Logger) Info(msg string, fields ...Field) {
	cfg := s.config()
	for _, logInstance := range cfg.instances {
		if logInstance.enabled.Load() {
			logInstance.logger.Info(msg, fieldsToZapFields(fields...)...)
		}
	}
}

func (s *Logger) InfoIgnoreCancel(ctx context.Context, msg string, fields ...Field) {
	if ctx.Err() != nil {
		return
	}
	if fieldsContainContextCancelled(fields...) {
		return
	}
	s.Info(msg, fields...)
}

func (s *Logger) Warn(msg string, fields ...Field) {
	cfg := s.config()
	for _, logInstance := range cfg.instances {
		if logInstance.enabled.Load() {
			logInstance.logger.Warn(msg, fieldsToZapFields(fields...)...)
		}
	}
}

func (s *Logger) WarnIgnoreCancel(ctx context.Context, msg string, fields ...Field) {
	if ctx.Err() != nil {
		return
	}
	if fieldsContainContextCancelled(fields...) {
		return
	}
	s.Warn(msg, fields...)
}

func (s *Logger) Error(msg string, fields ...Field) {
	cfg := s.config()
	for _, logInstance := range cfg.instances {
		if logInstance.enabled.Load() {
			logInstance.logger.Error(msg, fieldsToZapFields(fields...)...)
		}
	}
}

func (s *Logger) ErrorIgnoreCancel(ctx context.Context, msg string, fields ...Field) {
	if ctx.Err() != nil {
		return
	}
	if fieldsContainContextCancelled(fields...) {
		return
	}
	s.Error(msg, fields...)
}

func (s *Logger) Panic(msg string, fields ...Field) {
	cfg := s.config()
	var foundLogger bool
	for _, logInstance := range cfg.instances {
		if logInstance.enabled.Load() {
			foundLogger = true
			logInstance.logger.Panic(msg, fieldsToZapFields(fields...)...)
		}
	}
	if !foundLogger {
		panic(msg)
	}
}

func (s *Logger) DPanic(msg string, fields ...Field) {
	cfg := s.config()
	var foundLogger bool
	for _, logInstance := range cfg.instances {
		if logInstance.enabled.Load() {
			foundLogger = true
			logInstance.logger.DPanic(msg, fieldsToZapFields(fields...)...)
		}
	}
	if !foundLogger {
		panic(msg)
	}
}

func (s *Logger) Fatal(msg string, fields ...Field) {
	cfg := s.config()
	var foundLogger bool
	for _, logInstance := range cfg.instances {
		if logInstance.enabled.Load() {
			foundLogger = true
			logInstance.logger.Fatal(msg, fieldsToZapFields(fields...)...)
		}
	}
	if !foundLogger {
		fmt.Println(msg)
		os.Exit(1)
	}
}

// Deprecated: use structured logging instead.
func (s *Logger) TraceUnstruct(args ...interface{}) {
	cfg := s.config()
	for _, logInstance := range cfg.instances {
		if logInstance.enabled.Load() {
			logInstance.logger.Sugar().Debug(args...)
		}
	}
}

// Deprecated: use structured logging instead.
func (s *Logger) DebugUnstruct(args ...interface{}) {
	cfg := s.config()
	for _, logInstance := range cfg.instances {
		if logInstance.enabled.Load() {
			logInstance.logger.Sugar().Debug(args...)
		}
	}
}

// Deprecated: use structured logging instead.
func (s *Logger) InfoUnstruct(args ...interface{}) {
	cfg := s.config()
	for _, logInstance := range cfg.instances {
		if logInstance.enabled.Load() {
			logInstance.logger.Sugar().Info(args...)
		}
	}
}

// Deprecated: use structured logging instead.
func (s *Logger) WarnUnstruct(args ...interface{}) {
	cfg := s.config()
	for _, logInstance := range cfg.instances {
		if logInstance.enabled.Load() {
			logInstance.logger.Sugar().Warn(args...)
		}
	}
}

// Deprecated: use structured logging instead.
func (s *Logger) WarnIgnoreCancelUnstruct(ctx context.Context, args ...interface{}) {
	if ctx.Err() != nil {
		return
	}
	s.WarnUnstruct(args)
}

// Deprecated: use structured logging instead.
func (s *Logger) ErrorUnstruct(args ...interface{}) {
	cfg := s.config()
	for _, logInstance := range cfg.instances {
		if logInstance.enabled.Load() {
			logInstance.logger.Sugar().Error(args...)
		}
	}
}

// Deprecated: use structured logging instead.
func (s *Logger) ErrorIgnoreCancelUnstruct(ctx context.Context, args ...interface{}) {
	if ctx.Err() != nil {
		return
	}
	s.ErrorUnstruct(args)
}

// Deprecated: use structured logging instead.
func (s *Logger) PanicUnstruct(args ...interface{}) {
	cfg := s.config()
	var foundLogger bool
	for _, logInstance := range cfg.instances {
		if logInstance.enabled.Load() {
			foundLogger = true
			logInstance.logger.Sugar().Panic(args...)
		}
	}
	if !foundLogger {
		if len(args) >= 1 {
			if str, ok := args[0].(string); ok {
				panic(str)
			}
		}
	}
}

// Deprecated: use structured logging instead.
func (s *Logger) DPanicUnstruct(args ...interface{}) {
	cfg := s.config()
	var foundLogger bool
	for _, logInstance := range cfg.instances {
		if logInstance.enabled.Load() {
			foundLogger = true
			logInstance.logger.Sugar().DPanic(args...)
		}
	}
	if !foundLogger {
		if len(args) >= 1 {
			if str, ok := args[0].(string); ok {
				panic(str)
			}
		}
	}
}

// Deprecated: use structured logging instead.
func (s *Logger) FatalUnstruct(args ...interface{}) {
	cfg := s.config()
	var foundLogger bool
	for _, logInstance := range cfg.instances {
		if logInstance.enabled.Load() {
			foundLogger = true
			logInstance.logger.Sugar().Fatal(args...)
		}
	}
	if !foundLogger {
		if len(args) >= 1 {
			if str, ok := args[0].(string); ok {
				fmt.Println(str)
			}
		}
		os.Exit(1)
	}
}

// Deprecated: use structured logging instead.
func (s *Logger) TracefUnstruct(format string, args ...interface{}) {
	cfg := s.config()
	for _, logInstance := range cfg.instances {
		if logInstance.enabled.Load() {
			logInstance.logger.Sugar().Debugf(format, args...)
		}
	}
}

// Deprecated: use structured logging instead.
func (s *Logger) DebugfUnstruct(format string, args ...interface{}) {
	cfg := s.config()
	for _, logInstance := range cfg.instances {
		if logInstance.enabled.Load() {
			logInstance.logger.Sugar().Debugf(format, args...)
		}
	}
}

// Deprecated: use structured logging instead.
func (s *Logger) InfofUnstruct(format string, args ...interface{}) {
	cfg := s.config()
	for _, logInstance := range cfg.instances {
		if logInstance.enabled.Load() {
			logInstance.logger.Sugar().Infof(format, args...)
		}
	}
}

// Deprecated: use structured logging instead.
func (s *Logger) WarnfUnstruct(format string, args ...interface{}) {
	cfg := s.config()
	for _, logInstance := range cfg.instances {
		if logInstance.enabled.Load() {
			logInstance.logger.Sugar().Warnf(format, args...)
		}
	}
}

// Deprecated: use structured logging instead.
func (s *Logger) WarnfIgnoreCancelUnstruct(ctx context.Context, format string, args ...interface{}) {
	if ctx.Err() != nil {
		return
	}
	s.WarnfUnstruct(format, args)
}

// Deprecated: use structured logging instead.
func (s *Logger) ErrorfUnstruct(format string, args ...interface{}) {
	cfg := s.config()
	for _, logInstance := range cfg.instances {
		if logInstance.enabled.Load() {
			logInstance.logger.Sugar().Errorf(format, args...)
		}
	}
}

// Deprecated: use structured logging instead.
func (s *Logger) ErrorfIgnoreCancelUnstruct(ctx context.Context, format string, args ...interface{}) {
	if ctx.Err() != nil {
		return
	}
	s.ErrorfUnstruct(format, args)
}

// Deprecated: use structured logging instead.
func (s *Logger) PanicfUnstruct(format string, args ...interface{}) {
	cfg := s.config()
	var foundLogger bool
	for _, logInstance := range cfg.instances {
		if logInstance.enabled.Load() {
			foundLogger = true
			logInstance.logger.Sugar().Panicf(format, args...)
		}
	}
	if !foundLogger {
		panic(fmt.Sprintf(format, args...))
	}
}

// Deprecated: use structured logging instead.
func (s *Logger) DPanicfUnstruct(format string, args ...interface{}) {
	cfg := s.config()
	var foundLogger bool
	for _, logInstance := range cfg.instances {
		if logInstance.enabled.Load() {
			foundLogger = true
			logInstance.logger.Sugar().DPanicf(format, args...)
		}
	}
	if !foundLogger {
		panic(fmt.Sprintf(format, args...))
	}
}

// Deprecated: use structured logging instead.
func (s *Logger) FatalfUnstruct(format string, args ...interface{}) {
	cfg := s.config()
	var foundLogger bool
	for _, logInstance := range cfg.instances {
		if logInstance.enabled.Load() {
			foundLogger = true
			logInstance.logger.Sugar().Fatalf(format, args...)
		}
	}
	if !foundLogger {
		fmt.Println(fmt.Sprintf(format, args...))
		os.Exit(1)
	}
}
