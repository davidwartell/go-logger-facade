package logger

import (
	"context"
	"errors"
	"fmt"
	"os"
)

type contextKey int

const (
	loggerContextKey contextKey = iota + 1
	loggerFieldsContextKey
)

var allContextParams = []contextKey{
	loggerContextKey,
	loggerFieldsContextKey,
}

type ContextLogger struct {
	logger *Logger
	fields *ContextFields
}

type ContextFields struct {
	fields []Field
}

func WithAllValues(toCtx context.Context, fromCtxWithValues context.Context) context.Context {
	for _, paramName := range allContextParams {
		val := fromCtxWithValues.Value(paramName)
		if val != nil {
			toCtx = context.WithValue(toCtx, paramName, val)
		}
	}
	return toCtx
}

func WithLogger(ctx context.Context, logger *Logger) context.Context {
	return context.WithValue(ctx, loggerContextKey, logger)
}

func WithFields(ctx context.Context, fs ...Field) context.Context {
	cfields := new(ContextFields)
	existingCFields, found := fieldsOf(ctx)
	if found {
		cfields.fields = append(cfields.fields, existingCFields.fields...)
	}
	cfields.fields = append(cfields.fields, fs...)
	return context.WithValue(ctx, loggerFieldsContextKey, cfields)
}

func OfMust(ctx context.Context) (clogger *ContextLogger) {
	var err error
	clogger, err = Of(ctx)
	if err != nil {
		panic("logger not found in context")
	}
	return
}

func Of(ctx context.Context) (clogger *ContextLogger, err error) {
	logger, loggerExists := ctx.Value(loggerContextKey).(*Logger)
	if !loggerExists {
		err = errors.New("logger not found in context")
		return
	}
	clogger = &ContextLogger{
		logger: logger,
	}
	clogger.fields, _ = fieldsOf(ctx)
	return
}

func fieldsOf(ctx context.Context) (fields *ContextFields, found bool) {
	fields, found = ctx.Value(loggerFieldsContextKey).(*ContextFields)
	return
}

func (s *ContextLogger) Trace(msg string, fields ...Field) {
	if s.fields != nil {
		fields = append(fields, s.fields.fields...)
	}
	cfg := s.logger.config()
	for _, logInstance := range cfg.instances {
		if logInstance.enabled.Load() {
			logInstance.logger.Debug(msg, fieldsToZapFields(fields...)...)
		}
	}
}

func (s *ContextLogger) Debug(msg string, fields ...Field) {
	if s.fields != nil {
		fields = append(fields, s.fields.fields...)
	}
	cfg := s.logger.config()
	for _, logInstance := range cfg.instances {
		if logInstance.enabled.Load() {
			logInstance.logger.Debug(msg, fieldsToZapFields(fields...)...)
		}
	}
}

func (s *ContextLogger) Info(msg string, fields ...Field) {
	if s.fields != nil {
		fields = append(fields, s.fields.fields...)
	}
	cfg := s.logger.config()
	for _, logInstance := range cfg.instances {
		if logInstance.enabled.Load() {
			logInstance.logger.Info(msg, fieldsToZapFields(fields...)...)
		}
	}
}

func (s *ContextLogger) InfoIgnoreCancel(ctx context.Context, msg string, fields ...Field) {
	if ctx.Err() != nil {
		return
	}
	if s.fields != nil {
		fields = append(fields, s.fields.fields...)
	}
	cfg := s.logger.config()
	for _, logInstance := range cfg.instances {
		if logInstance.enabled.Load() {
			logInstance.logger.Info(msg, fieldsToZapFields(fields...)...)
		}
	}
}

func (s *ContextLogger) Warn(msg string, fields ...Field) {
	if s.fields != nil {
		fields = append(fields, s.fields.fields...)
	}
	cfg := s.logger.config()
	for _, logInstance := range cfg.instances {
		if logInstance.enabled.Load() {
			logInstance.logger.Warn(msg, fieldsToZapFields(fields...)...)
		}
	}
}

func (s *ContextLogger) WarnIgnoreCancel(ctx context.Context, msg string, fields ...Field) {
	if ctx.Err() != nil {
		return
	}
	if s.fields != nil {
		fields = append(fields, s.fields.fields...)
	}
	cfg := s.logger.config()
	for _, logInstance := range cfg.instances {
		if logInstance.enabled.Load() {
			logInstance.logger.Warn(msg, fieldsToZapFields(fields...)...)
		}
	}
}

func (s *ContextLogger) Error(msg string, fields ...Field) {
	if s.fields != nil {
		fields = append(fields, s.fields.fields...)
	}
	cfg := s.logger.config()
	for _, logInstance := range cfg.instances {
		if logInstance.enabled.Load() {
			logInstance.logger.Error(msg, fieldsToZapFields(fields...)...)
		}
	}
}

func (s *ContextLogger) ErrorIgnoreCancel(ctx context.Context, msg string, fields ...Field) {
	if ctx.Err() != nil {
		return
	}
	if s.fields != nil {
		fields = append(fields, s.fields.fields...)
	}
	cfg := s.logger.config()
	for _, logInstance := range cfg.instances {
		if logInstance.enabled.Load() {
			logInstance.logger.Error(msg, fieldsToZapFields(fields...)...)
		}
	}
}

func (s *ContextLogger) Panic(msg string, fields ...Field) {
	if s.fields != nil {
		fields = append(fields, s.fields.fields...)
	}
	cfg := s.logger.config()
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

func (s *ContextLogger) DPanic(msg string, fields ...Field) {
	if s.fields != nil {
		fields = append(fields, s.fields.fields...)
	}
	cfg := s.logger.config()
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

func (s *ContextLogger) Fatal(msg string, fields ...Field) {
	if s.fields != nil {
		fields = append(fields, s.fields.fields...)
	}
	cfg := s.logger.config()
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
