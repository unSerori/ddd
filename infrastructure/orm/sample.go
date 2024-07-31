// サンプルデータ
package orm

import (
	"ddd/infrastructure/orm/model"

	"xorm.io/xorm"
)

// テストデータ
// 外部キーの参照先テーブルを先に登録する必要がある。
func sampleData() []interface{} {
	return []interface{}{ // サンプル
		// ユーザー
		&model.User{
			UserUuid:    "9efeb117-1a34-4012-b57c-7f1a4033adb9",
			UserName:    "test teacher",
			Age:         18,
			MailAddress: "test-teacher@gmail.com",
			Password:    "$2a$10$Ig/s1wsrXBuZ7qvjudr4CeQFhqJTLQpoAAp1LrBNh5jX9VZZxa3R6", // C@tt
			JtiUuid:     "42c28ac4-0ba4-4f81-8813-814dc92e2f40",
			IconPath:    "b298dc3c-67bb-4c93-a8e6-f2ae6338ef4c",
		},
		&model.User{
			UserUuid:    "3cac1684-c1e0-47ae-92fd-6d7959759224",
			UserName:    "test pupil",
			Age:         20,
			MailAddress: "test-pupil@gmail.com",
			Password:    "$2a$10$8hJGyU235UMV8NjkozB7aeHtgxh39wg/ocuRXW9jN2JDdO/MRz.fW", // C@tp
			JtiUuid:     "14dea318-8581-4cab-b233-995ce8e1a948",
			IconPath:    "691820ca-8e84-4926-b203-e3f0b3a2f0b0",
		},
		&model.User{
			UserUuid:    "9efeb117-1a34-4012-b57c-7f1a4033adb9",
			UserName:    "test teacher",
			Age:         100,
			MailAddress: "test-teacher@gmail.com",
			Password:    "$2a$10$Ig/s1wsrXBuZ7qvjudr4CeQFhqJTLQpoAAp1LrBNh5jX9VZZxa3R6", // C@tt
			JtiUuid:     "42c28ac4-0ba4-4f81-8813-814dc92e2f40",
			IconPath:    "e16d9717-309f-4039-85ec-f881cb7dcdf1",
		},
		&model.User{
			UserUuid:    "868c0804-cf1b-43e2-abef-08f7ef58fcd0",
			UserName:    "test parent",
			Age:         30,
			MailAddress: "test-parent@gmail.com",
			Password:    "$2a$10$8hJGyU235UMV8NjkozB7aeHtgxh39wg/ocuRXW9jN2JDdO/MRz.fW", // C@tp
			JtiUuid:     "0553853f-cbcf-49e2-81d6-a4c7e4b1b470",
			IconPath:    "199be890-0ef6-4689-9efb-e1be03e0deaa",
		},

		// 部屋
	}
}

// サンプルデータ作成
func RegisterSample(db *xorm.Engine) {
	// テスト用データ作成
	for _, sample := range sampleData() {
		db.Insert(sample)
	}
}
