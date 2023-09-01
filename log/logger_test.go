package log

import (
	"gorm.io/gorm/logger"
	"testing"
)

func TestNewCustomLog(t *testing.T) {
	t.Log(NewCustomLog("test.log").LogMode(logger.Info))
}
