package zsw

import (
	"fmt"

	"github.com/streamingfast/logging"
	"go.uber.org/zap"
)

var traceEnabled = logging.IsTraceEnabled("zswchain-go", "github.com/zhongshuwen/zswchain-go")
var zlog = zap.NewNop()

func init() {
	logging.Register("github.com/zhongshuwen/zswchain-go", &zlog)
}

func EnableDebugLogging(l *zap.Logger) {
	traceEnabled = true
	zlog = l
}

type logStringerFunc func() string

func (f logStringerFunc) String() string { return f() }

func typeField(field string, v interface{}) zap.Field {
	return zap.Stringer(field, logStringerFunc(func() string {
		return fmt.Sprintf("%T", v)
	}))
}

func newLogger(production bool) (l *zap.Logger) {
	if production {
		l, _ = zap.NewProduction()
	} else {
		l, _ = zap.NewDevelopment()
	}
	return
}

// NewLogger a wrap to newLogger
func NewLogger(production bool) *zap.Logger {
	return newLogger(production)
}
