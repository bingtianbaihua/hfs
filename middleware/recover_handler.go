package middleware

import (
	"net/http"

	"github.com/bingtianbaihua/hfs/log"
)

type RecoverAdapter struct{}

// NewRecovery creates a new instance of Recovery
func (rv *RecoverAdapter) HandleTask(w http.ResponseWriter, r *http.Request, stk func(http.ResponseWriter, *http.Request)) {
	log.Info("start catch panic...")
	defer func() {
		if err := recover(); err != nil {
			log.Warn("panic with error: %v\n", err)
		}
	}()

	stk(w, r)
}

func NewRecoverAdapter() *RecoverAdapter {
	return &RecoverAdapter{}
}
