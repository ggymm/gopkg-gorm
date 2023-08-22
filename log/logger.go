package log

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/utils"

	"github.com/ggymm/gopkg/log"
)

type CustomLog struct {
	log      zerolog.Logger
	LogLevel logger.LogLevel
}

func NewCustomLog(filename string) *CustomLog {
	return &CustomLog{
		log:      log.InitCustom(filename),
		LogLevel: logger.Info,
	}
}

func (l *CustomLog) LogMode(level logger.LogLevel) logger.Interface {
	l.LogLevel = level
	return l
}

func (l *CustomLog) Info(_ context.Context, msg string, data ...interface{}) {
	l.log.Info().Str("file", utils.FileWithLineNum()).Msgf(msg, data...)
}

func (l *CustomLog) Warn(_ context.Context, msg string, data ...interface{}) {
	l.log.Warn().Str("file", utils.FileWithLineNum()).Msgf(msg, data...)
}

func (l *CustomLog) Error(_ context.Context, msg string, data ...interface{}) {
	l.log.Error().Str("file", utils.FileWithLineNum()).Msgf(msg, data...)
}

// Trace print sql message
func (l *CustomLog) Trace(_ context.Context, begin time.Time, fc func() (string, int64), err error) {
	elapsed := time.Since(begin)
	costTime := fmt.Sprintf("%.3fms", float64(elapsed.Nanoseconds())/1e6)

	file := utils.FileWithLineNum()

	sql, rows := fc()
	sql = strings.Replace(sql, "\"", "'", -1)

	if err != nil {
		// 忽略记录不存在的错误
		if !errors.Is(err, logger.ErrRecordNotFound) {
			l.log.Error().Err(errors.WithStack(err)).Str("sql", sql).Str("source", file).
				Str("costTime", costTime).Int64("rowsAffected", rows).Msg("SQLTrace")
		}
	} else {
		l.log.Info().Str("sql", sql).Str("source", file).
			Str("costTime", costTime).Int64("rowsAffected", rows).Msg("SQLTrace")
	}
}
