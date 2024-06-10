package repo

import (
	"context"
	"time"

	"gorm.io/gorm/logger"

	"github.com/dcalsky/kong_service_demo/internal/common/logs"
)

type GormLogger struct {
}

func (s *GormLogger) LogMode(level logger.LogLevel) logger.Interface {
	return s
}

func (s *GormLogger) Info(ctx context.Context, msg string, i ...interface{}) {
	logs.Infof(ctx, msg, i)
}

func (s *GormLogger) Warn(ctx context.Context, msg string, i ...interface{}) {
	logs.Warnf(ctx, msg, i)
}

func (s *GormLogger) Error(ctx context.Context, msg string, i ...interface{}) {
	logs.Errorf(ctx, msg, i)
}

func (s *GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	elapsed := time.Since(begin)
	sql, rows := fc()
	if err != nil {
		logs.Errorf(ctx, "[gorm] duration: %d ms, affected rows: %d, sql: %s", elapsed.Milliseconds(), rows, sql)
	} else {
		logs.Infof(ctx, "[gorm] duration: %d ms, affected rows: %d, sql: %s", elapsed.Milliseconds(), rows, sql)
	}
}
