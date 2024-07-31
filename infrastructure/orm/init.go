// ORMの初期化

package orm

import (
	"ddd/common/logging"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql" // ブランクインポート呼ばれ、コード上では直接利用しないが、使用しているライブラリ等が必要としているドライバなどのinit関数に対して、_が初期化を強制する。もしコード上でこのライブラリを直接利用する場合、そちらでインポートされるのでこの行は不要だが、わかりやすくするために明示的にインポートする
	"xorm.io/xorm"
)

// ORMを使ってDB鯖に接続
func DBConnect() (*xorm.Engine, error) {
	// 環境変数から取得
	dbUser := os.Getenv("MYSQL_USER")
	dbPass := os.Getenv("MYSQL_PASSWORD")
	dbHost := os.Getenv("MYSQL_HOST")
	dbPort := os.Getenv("MYSQL_PORT")
	dbDB := os.Getenv("MYSQL_DATABASE")

	// Mysqlに接続
	db, err := xorm.NewEngine( // dbとエラーを取得
		"mysql", // dbの種類"root:root@tcp(db:3306)/cgroup?charset=utf8"
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", dbUser, dbPass, dbHost, dbPort, dbDB), // 接続情報
	)
	if err != nil { // エラー処理
		logging.ErrorLog("Couldnt connect to the db server."+fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", dbUser, dbPass, dbHost, dbPort, dbDB), err)
		return nil, err
	} else {
		logging.SuccessLog("Could connect to the db server." + fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", dbUser, dbPass, dbHost, dbPort, dbDB))
	}

	return db, nil
}

// ORMの初期化処理
func InitDB() (*xorm.Engine, error) {
	// 接続
	db, err := DBConnect()
	if err != nil {
		return nil, err
	}

	// テーブル作成
	err = MigrationTable(db)
	if err != nil {
		logging.ErrorLog("Failed migration.", err)
		return nil, err
	}

	// 設定
	db.ShowSQL(true)       // SQL文の表示
	db.SetMaxOpenConns(10) // 接続数

	return db, nil
}
