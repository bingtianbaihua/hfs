package middleware

import (
	"errors"
	"net/http"

	"github.com/bingtianbaihua/hfs/log"
)

const (
	TrueClientIP   = "True-Client-IP"
	TrueRealIP     = "True-Real-IP"
	XForwardedFor  = "X-Forwarded-For"
	XOriginatingIP = "X-Originating-IP"
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
	log.Info("request method is: %v, url is: %v, remote ip: %v", r.Method, r.URL.String(), ad.handleActualRequest(w, r))
	stk(w, r)
	log.Info("resp header is: %v", w.Header())
}

func (ad *LogAdapter) handleActualRequest(w http.ResponseWriter, r *http.Request) string {
	var ipAddress string
	var ipSources = []string{
		r.Header.Get(TrueClientIP),
		r.Header.Get(TrueRealIP),
		r.Header.Get(XForwardedFor),
		r.Header.Get(XOriginatingIP),
	}

	for _, ip := range ipSources {
		if ip != "" {
			ipAddress = ip
			break
		}
	}
	return ipAddress
}
