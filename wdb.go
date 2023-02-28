package lucky

import (
	"github.com/kevin-zhangzh/lucky-key/schema"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
	"path"
)

const (
	sqliteName = "seed.db"
)

type Wdb struct {
	Db *gorm.DB
}

func NewSqliteDb(dbDir string) *Wdb {
	if err := os.MkdirAll(dbDir, os.ModePerm); err != nil {
		panic(err)
	}
	db, err := gorm.Open(sqlite.Open(path.Join(dbDir, sqliteName)), &gorm.Config{
		Logger:          logger.Default.LogMode(logger.Silent),
		CreateBatchSize: 200,
	})
	if err != nil {
		panic(err)
	}
	log.Info("connect sqlite db success")
	return &Wdb{Db: db}

}

// when use sqlite,same index name in different table will lead to migrate failed,

func (w *Wdb) Migrate() error {
	return w.Db.AutoMigrate(&schema.Asset{})
}

func (w *Wdb) InsertAsset(asset schema.Asset) error {
	return w.Db.Create(&asset).Error
}
