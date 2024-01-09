package task

import (
	"context"
	"github.com/davidwartell/go-logger-facade/logger"
	"runtime/debug"
	"strings"
)

type Task interface {
	StartTask()
	StopTask()
}

type BaseTask struct{}

// LogTraceStruct deprecated
//
//goland:noinspection GoUnusedExportedFunction
func LogTraceStruct(taskName string, msg string, fields ...logger.Field) {
	fieldsWithService := append(fields, logger.String("task", taskName))
	logger.Instance().Trace(getTaskLogPrefix(taskName, msg), fieldsWithService...)
}

// LogDebugStruct deprecated
//
//goland:noinspection GoUnusedExportedFunction
func LogDebugStruct(taskName string, msg string, fields ...logger.Field) {
	fieldsWithService := append(fields, logger.String("task", taskName))
	logger.Instance().Debug(getTaskLogPrefix(taskName, msg), fieldsWithService...)
}

// LogInfoStruct deprecated
//
//goland:noinspection GoUnusedExportedFunction
func LogInfoStruct(taskName string, msg string, fields ...logger.Field) {
	fieldsWithService := append(fields, logger.String("task", taskName))
	logger.Instance().Info(getTaskLogPrefix(taskName, msg), fieldsWithService...)
}

// LogWarnStruct deprecated
//
//goland:noinspection GoUnusedExportedFunction
func LogWarnStruct(taskName string, msg string, fields ...logger.Field) {
	fieldsWithService := append(fields, logger.String("task", taskName))
	logger.Instance().Warn(getTaskLogPrefix(taskName, msg), fieldsWithService...)
}

// LogWarnStructIgnoreCancel deprecated
//
//goland:noinspection GoUnusedExportedFunction
func LogWarnStructIgnoreCancel(ctx context.Context, taskName string, msg string, fields ...logger.Field) {
	fieldsWithService := append(fields, logger.String("task", taskName))
	logger.Instance().WarnIgnoreCancel(ctx, getTaskLogPrefix(taskName, msg), fieldsWithService...)
}

// LogErrorStruct deprecated
//
//goland:noinspection GoUnusedExportedFunction
func LogErrorStruct(taskName string, msg string, fields ...logger.Field) {
	fieldsWithService := append(fields, logger.String("task", taskName))
	logger.Instance().Error(getTaskLogPrefix(taskName, msg), fieldsWithService...)
}

// LogErrorStructIgnoreCancel deprecated
//
//goland:noinspection GoUnusedExportedFunction
func LogErrorStructIgnoreCancel(ctx context.Context, taskName string, msg string, fields ...logger.Field) {
	fieldsWithService := append(fields, logger.String("task", taskName))
	logger.Instance().ErrorIgnoreCancel(ctx, getTaskLogPrefix(taskName, msg), fieldsWithService...)
}

//goland:noinspection GoUnusedExportedFunction
func HandlePanic(taskName string) {
	if err := recover(); err != nil {
		LogErrorStruct(taskName, "panic occurred", logger.Any("err", err), logger.String("stacktrace", string(debug.Stack())))
	}
}

//goland:noinspection GoUnusedExportedFunction
func LogTrace(taskName string, msg string, fields ...logger.Field) {
	fieldsWithService := append(fields, logger.String("task", taskName))
	logger.Instance().Trace(getTaskLogPrefix(taskName, msg), fieldsWithService...)
}

//goland:noinspection GoUnusedExportedFunction
func LogDebug(taskName string, msg string, fields ...logger.Field) {
	fieldsWithService := append(fields, logger.String("task", taskName))
	logger.Instance().Debug(getTaskLogPrefix(taskName, msg), fieldsWithService...)
}

//goland:noinspection GoUnusedExportedFunction
func LogInfo(taskName string, msg string, fields ...logger.Field) {
	fieldsWithService := append(fields, logger.String("task", taskName))
	logger.Instance().Info(getTaskLogPrefix(taskName, msg), fieldsWithService...)
}

//goland:noinspection GoUnusedExportedFunction
func LogInfoIgnoreCancel(ctx context.Context, taskName string, msg string, fields ...logger.Field) {
	fieldsWithService := append(fields, logger.String("task", taskName))
	logger.Instance().InfoIgnoreCancel(ctx, getTaskLogPrefix(taskName, msg), fieldsWithService...)
}

//goland:noinspection GoUnusedExportedFunction
func LogWarn(taskName string, msg string, fields ...logger.Field) {
	fieldsWithService := append(fields, logger.String("task", taskName))
	logger.Instance().Warn(getTaskLogPrefix(taskName, msg), fieldsWithService...)
}

//goland:noinspection GoUnusedExportedFunction
func LogWarnIgnoreCancel(ctx context.Context, taskName string, msg string, fields ...logger.Field) {
	fieldsWithService := append(fields, logger.String("task", taskName))
	logger.Instance().WarnIgnoreCancel(ctx, getTaskLogPrefix(taskName, msg), fieldsWithService...)
}

//goland:noinspection GoUnusedExportedFunction
func LogError(taskName string, msg string, fields ...logger.Field) {
	fieldsWithService := append(fields, logger.String("task", taskName))
	logger.Instance().Error(getTaskLogPrefix(taskName, msg), fieldsWithService...)
}

//goland:noinspection GoUnusedExportedFunction
func LogErrorIgnoreCancel(ctx context.Context, taskName string, msg string, fields ...logger.Field) {
	fieldsWithService := append(fields, logger.String("task", taskName))
	logger.Instance().ErrorIgnoreCancel(ctx, getTaskLogPrefix(taskName, msg), fieldsWithService...)
}

func getTaskLogPrefix(taskName string, format string) string {
	var sb strings.Builder
	sb.WriteString("[")
	sb.WriteString(taskName)
	sb.WriteString("] ")
	sb.WriteString(format)
	return sb.String()
}
