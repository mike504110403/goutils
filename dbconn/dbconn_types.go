package dbconn

import (
	"database/sql"
	"sync"
	"time"
)

// DBDriver 資料庫驅動
type DBDriver string

const (
	DBDriverMySQL DBDriver = "mysql"
)

// DBName 資料庫名稱
type DBName string

// 連線設定
type Config struct {
	ConnMaxIdleTime *time.Duration
	ConnMaxLifetime *time.Duration
	MaxOpenConns    *int
	MaxIdleConns    *int
}

// 資料庫連線設定
type DBConfig struct {
	db         *sql.DB
	DBDriver   DBDriver
	DSNSource  any
	ConnConfig *Config
	Once       sync.Once
}
