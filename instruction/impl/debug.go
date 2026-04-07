package impl

import (
	"github.com/AlanRostem/mu-8/logger"
	"github.com/AlanRostem/mu-8/num"
)

func pcDebugf(pc num.DByte, format string, args ...any) {
	args = append([]any{pc}, args...)
	logger.Debugf("0x%04X: "+format, args...)
}
