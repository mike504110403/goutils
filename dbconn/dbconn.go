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
	dbList                        = make(map[DBName]DBConfig)
	connMaxIdleTime time.Duration = 10 * time.Minute
	connMaxLifetime time.Duration = 1 * time.Hour
	maxOpenConns    int           = 100
	maxIdleConns    int           = 10
	// 預設連線設定
	defaultConfig = Config{
		ConnMaxIdleTime: &connMaxIdleTime,
		ConnMaxLifetime: &connMaxLifetime,
		MaxOpenConns:    &maxOpenConns,
		MaxIdleConns:    &maxIdleConns,
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
func setDBConfig(db *sql.DB, config *Config) {
	if config == nil {
		config = &defaultConfig
	}

	setConnMaxIdleTime(db, config.ConnMaxIdleTime, defaultConfig.ConnMaxIdleTime)
	setConnMaxLifetime(db, config.ConnMaxLifetime, defaultConfig.ConnMaxLifetime)
	setMaxOpenConns(db, config.MaxOpenConns, defaultConfig.MaxOpenConns)
	setMaxIdleConns(db, config.MaxIdleConns, defaultConfig.MaxIdleConns)
}

// 設定連線最大閒置時間
func setConnMaxIdleTime(db *sql.DB, value, defaultValue *time.Duration) {
	if value != nil {
		db.SetConnMaxIdleTime(*value)
	} else {
		db.SetConnMaxIdleTime(*defaultValue)
	}
}

// 設定連線最大存活時間
func setConnMaxLifetime(db *sql.DB, value, defaultValue *time.Duration) {
	if value != nil {
		db.SetConnMaxLifetime(*value)
	} else {
		db.SetConnMaxLifetime(*defaultValue)
	}
}

// 設定最大開啟連線數
func setMaxOpenConns(db *sql.DB, value, defaultValue *int) {
	if value != nil {
		db.SetMaxOpenConns(*value)
	} else {
		db.SetMaxOpenConns(*defaultValue)
	}
}

// 設定最大閒置連線數
func setMaxIdleConns(db *sql.DB, value, defaultValue *int) {
	if value != nil {
		db.SetMaxIdleConns(*value)
	} else {
		db.SetMaxIdleConns(*defaultValue)
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
