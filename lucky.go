package lucky

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-co-op/gocron"
	"github.com/kevin-zhangzh/lucky-key/common"
	"time"
)

var log = common.NewLog("otc")

type Lucky struct {
	cli       *ethclient.Client
	wdb       *Wdb
	scheduler *gocron.Scheduler
}

func NewLucky(dbDir, rpcUrl string) *Lucky {
	cli, err := ethclient.Dial(rpcUrl)
	if err != nil {
		panic(err)
	}
	wdb := NewSqliteDb(dbDir)
	err = wdb.Migrate()
	if err != nil {
		panic(err)
	}
	return &Lucky{
		cli:       cli,
		wdb:       wdb,
		scheduler: gocron.NewScheduler(time.UTC),
	}
}

func (l *Lucky) Run() {
	l.runJobs()
}
