// マイグレーション

package orm

import (
	"ddd/common/logging"
	"ddd/infrastructure/orm/model"

	"xorm.io/xorm"
)

// マイグレーションするテーブル一覧を追加していく
func tableModels() []interface{} {
	return []interface{}{
		new(model.User),
	}
}

// 外部キーを設定
func initFK(db *xorm.Engine, models []interface{}) error { // テーブルモデルインスタンスのスライスを受け取る
	for _, model := range models { // それぞれのテーブルのFKを処理
		if modelImplFKs, ok := model.(interface{ FKs() []string }); ok { // テーブル構造体のメソッドがFKsを実装しているかアサーションして確かめる
			for _, fkQuery := range modelImplFKs.FKs() { // テーブルの複数のFK制約（FKs関数が文字列スライスを返す）をそれぞれ処理
				if _, err := db.Exec(fkQuery); err != nil { // FK制約追加のクエリ文字列を直接実行
					logging.ErrorLog("Failed to execute to registration FK constraint.", err)
					return err
				}
			}
		}
	}

	return nil
}

// マイグレーション関連
func MigrationTable(db *xorm.Engine) error {
	// テーブルがないなら自動で作成 // xormがテーブル作成時に列名をスネークケースにしてくれる  // 列情報の追加変更は反映するが列の削除は反映しない
	if exists, _ := db.IsTableExist(&model.User{}); !exists { // この判定で、外部キー設定済みのテーブルの再Sync2時に外部キーのインデックスを消せないエラーを防いでいる
		// マイグレーションするテーブル一覧を取得
		tableModels := tableModels()

		// テーブルの登録
		err := db.Sync2( // ここにテーブルを追加
			tableModels..., // ...でスライスを展開して可変長引数として渡す
		)
		if err != nil {
			logging.ErrorLog("Failed to sync database.", err)
			return err
		}

		// FK制約追加
		err = initFK(db, // FK制約があればここに追加
			tableModels,
		)
		if err != nil {
			logging.ErrorLog("Failed to set foreign key.", err)
			return err
		}
	}

	// サンプルデータ作成
	RegisterSample(db)

	return nil
}
