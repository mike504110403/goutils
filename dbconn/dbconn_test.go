package dbconn

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	testDsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?%s",
		"user_name",
		"user_password",
		"server_ip",
		"server_port",
		"db_name",
		"dsn_option",
	)
	// 設定測試用的資料庫連線設定
	cfgList := map[DBName]DBConfig{
		DBName("test_db"): {
			DBDriver:  DBDriverMySQL,
			DSNSource: testDsn,
		},
	}

	// 建立新的資料庫連線
	New(cfgList)

	// 取得連線
	db, err := DBName("test_db").DB()
	if err != nil {
		t.Fatalf("無法取得資料庫連線：%v", err)
	}
	defer db.Close()

	// 嘗試連線資料庫
	err = db.Ping()
	if err != nil {
		t.Fatalf("無法連線資料庫：%v", err)
	}

	// 測試通過
	t.Logf("成功連線到資料庫")
}
