package middleware

import (
	"errors"
	"net/http"

	"github.com/bingtianbaihua/hfs/log"
)

type LogConfig struct{}

type LogAdapter struct {
	xl *log.Logger
}

func NewLogHandler(cfg *LogConfig) (*LogAdapter, error) {
	if cfg == nil {
		return nil, errors.New("config can not be empty")
	}
	return &LogAdapter{
		xl: log.NewLogger(),
	}, nil
}

func (ad *LogAdapter) HandleTask(w http.ResponseWriter, r *http.Request, stk func(http.ResponseWriter, *http.Request)) {
	log.Info("request method is:%v, url is:%v", r.Method, r.URL.String())
	stk(w, r)
	log.Info("resp header is:%v", w.Header())
}
