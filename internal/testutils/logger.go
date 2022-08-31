package testutils

import "github.com/TcMits/ent-clean-template/pkg/infrastructure/logger"

var _ logger.Interface = &NullLogger{}

type NullLogger struct{}

func (NullLogger) Info(string, ...any) {}
func (NullLogger) Warn(string, ...any) {}
func (NullLogger) Error(any, ...any)   {}
func (NullLogger) Fatal(any, ...any)   {}
func (NullLogger) Debug(any, ...any)   {}
