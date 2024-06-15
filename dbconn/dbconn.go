package dbconn

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	mlog "github.com/mike504110403/goutils/log"
)

var (
	// 連線池
	dbList = make(map[DBName]DBConfig)
	// 預設連線設定
	defaultConfig = DefaultConfig{
		ConnMaxIdleTime: 10 * time.Minute,
		ConnMaxLifetime: 1 * time.Hour,
		MaxOpenConns:    100,
		MaxIdleConns:    10,
	}
)

// 初始化連線
func New(cfgList map[DBName]DBConfig) {
	for dbName, dbCfg := range cfgList {
		db := &sql.DB{}
		dbCfg.Once.Do(func() {
			switch dbCfg.DBDriver {
			case DBDriverMySQL:
				if _db, err := sql.Open(string(dbCfg.DBDriver), dbCfg.DSNSource.(string)); err != nil {
					mlog.Fatal(fmt.Sprintf("SQL Connection Error: %s", err.Error()))
					os.Exit(1)
				} else {
					db = _db
					setDBConfig(db, dbCfg.ConnConfig)
				}

			}
		})
		dbCfg.db = db
		dbList[dbName] = dbCfg
	}
}

// 設定連線池參數
func setDBConfig(db *sql.DB, config Config) {
	if config.ConnMaxIdleTime != nil {
		db.SetConnMaxIdleTime(*config.ConnMaxIdleTime)
	} else {
		db.SetConnMaxIdleTime(defaultConfig.ConnMaxIdleTime)
	}

	if config.ConnMaxLifetime != nil {
		db.SetConnMaxLifetime(*config.ConnMaxLifetime)
	} else {
		db.SetConnMaxLifetime(defaultConfig.ConnMaxLifetime)
	}

	if config.MaxOpenConns != nil {
		db.SetMaxOpenConns(*config.MaxOpenConns)
	} else {
		db.SetMaxOpenConns(defaultConfig.MaxOpenConns)
	}

	if config.MaxIdleConns != nil {
		db.SetMaxIdleConns(*config.MaxIdleConns)
	} else {
		db.SetMaxIdleConns(defaultConfig.MaxIdleConns)
	}
}

// 取得連線
func (dbName DBName) DB() (*sql.DB, error) {
	return dbList[dbName].db, dbList[dbName].db.Ping()
}

// 取得事務連線
func (dbName DBName) TX() (*sql.Tx, error) {
	return dbList[dbName].db.Begin()
}
