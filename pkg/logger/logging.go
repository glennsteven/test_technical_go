// Package logger
package logger

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"

	"technical_test_go/technical_test_go/pkg/util"
)

var (
	conf = &Config{}
)

// Field log object
type Field struct {
	Key   string
	Value any
}

// FieldFunc log
type FieldFunc func(key string, value any) *Field

type Fields []Field

// NewFields create instance new field
func NewFields(p ...Field) Fields {
	x := Fields{}

	for i := 0; i < len(p); i++ {
		x.Append(p[i])
	}

	return x
}

// Append new field
func (f *Fields) Append(p Field) {
	*f = append(*f, p)
}

// Any log
func Any(k string, v any) Field {
	return Field{
		Key:   k,
		Value: v,
	}
}

// String log
func String(k string, v string) Field {
	return Field{
		Key:   k,
		Value: v,
	}
}

// EventName log
func EventName(v any) Field {
	return Field{
		Key:   EventNameKey,
		Value: v,
	}
}

// EventId log
func EventId(v any) Field {
	return Field{
		Key:   EventIdKey,
		Value: v,
	}
}

// MessageFormat message with custom argument
func MessageFormat(format string, args ...any) string {
	return fmt.Sprintf(format, args...)
}

func extract(args ...Field) map[string]any {
	data := map[string]any{}

	if len(args) == 0 {
		return data
	}

	for _, fl := range args {
		data[fl.Key] = fl.Value
	}
	return data
}

// Error log
func Error(arg any, fl ...Field) {
	logrus.WithFields(
		addField(logrus.Fields{
			EventKey: extract(fl...),
		}),
	).Error(arg)

}

func Info(arg any, fl ...Field) {
	logrus.WithFields(
		addField(logrus.Fields{
			EventKey: extract(fl...),
		}),
	).Info(arg)
}

func Debug(arg any, fl ...Field) {
	logrus.WithFields(
		addField(logrus.Fields{
			EventKey: extract(fl...),
		}),
	).Debug(arg)
}

// Fatal log
func Fatal(arg any, fl ...Field) {
	logrus.WithFields(
		addField(logrus.Fields{
			EventKey: extract(fl...),
		}),
	).Fatal(arg)
}

// Warn log
func Warn(arg any, fl ...Field) {
	logrus.WithFields(
		addField(logrus.Fields{
			EventKey: extract(fl...),
		}),
	).Warn(arg)
}

// Trace log
func Trace(arg any, fl ...Field) {
	logrus.WithFields(
		addField(logrus.Fields{
			EventKey: extract(fl...),
		}),
	).Trace(arg)
}

// AccessLog http accessing log
func AccessLog(arg any, fl ...Field) {
	logrus.WithFields(
		addField(logrus.Fields{
			EventKey: extract(fl...),
		}),
	).Info(arg)
}

// InfoWithContext log info with context
func InfoWithContext(ctx context.Context, arg any, fl ...Field) {
	logrus.WithFields(extractContext(ctx.Value("access"), map[string]any{
		EventKey: extract(fl...),
	})).WithContext(ctx).Info(arg)
}

// WarnWithContext log warn with context
func WarnWithContext(ctx context.Context, arg any, fl ...Field) {
	logrus.WithFields(extractContext(ctx.Value("access"), map[string]any{
		EventKey: extract(fl...),
	})).WithContext(ctx).Warn(arg)
}

// ErrorWithContext log error with context
func ErrorWithContext(ctx context.Context, arg any, fl ...Field) {
	logrus.WithFields(extractContext(ctx.Value("access"), map[string]any{
		EventKey: extract(fl...),
	})).WithContext(ctx).Error(arg)
}

// DebugWithContext log debug with context
func DebugWithContext(ctx context.Context, arg any, fl ...Field) {
	logrus.WithFields(extractContext(ctx.Value("access"), map[string]any{
		EventKey: extract(fl...),
	})).WithContext(ctx).Debug(arg)
}

// TraceWithContext log trace with context
func TraceWithContext(ctx context.Context, arg any, fl ...Field) {
	logrus.WithFields(extractContext(ctx.Value("access"), map[string]any{
		EventKey: extract(fl...),
	})).WithContext(ctx).Trace(arg)
}

func extractContext(i any, logField map[string]any) map[string]any {

	if Environment(conf.Environment) != "" {
		logField[EnvironmentKey] = Environment(conf.Environment)
	}

	if conf.ServiceName != "" {
		logField[ServiceKey] = conf.ServiceName
	}

	if util.IsSameType(i, logField) {
		x := i.(map[string]any)
		for k, v := range x {
			logField[k] = v
		}
	}
	return logField
}

func addField(f logrus.Fields) logrus.Fields {
	if Environment(conf.Environment) != "" {
		f[EnvironmentKey] = Environment(conf.Environment)
	}

	if conf.ServiceName != "" {
		f[ServiceKey] = conf.ServiceName
	}
	return f
}
