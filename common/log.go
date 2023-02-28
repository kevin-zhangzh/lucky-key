package common

import (
	"github.com/inconshreveable/log15"
)

func NewLog(serverName string) log15.Logger {
	lg := log15.New("module", serverName)

	// 默认的 logger handle
	h := lg.GetHandler()
	// 集成 sentry 的 logger handle

	lg.SetHandler(h)

	return lg
}
