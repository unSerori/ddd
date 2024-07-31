// ユーザーテーブル

package model

// ユーザテーブル  // モデルを構造体で定義
type User struct { // typeで型の定義, structは構造体
	UserUuid    string `xorm:"varchar(36) pk" json:"userUUID"`                  // ユーザのUUID
	UserName    string `xorm:"varchar(25) not null" json:"userName"`            // 名前
	Age         int    `xorm:"INT not null" json:"age"`                         // 年齢
	MailAddress string `xorm:"varchar(256) not null unique" json:"mailAddress"` // メアド
	Password    string `xorm:"varchar(60) not null" json:"password"`            // bcrypt化されたパスワード
	JtiUuid     string `xorm:"varchar(36) unique" json:"jwtUUID"`               // jwtクレームのuuid
	IconPath    string `xorm:"varchar(36) unique" json:"iconPath"`
}

// テーブル名
func (User) TableName() string {
	return "users"
}

// FK制約 制約がなくても空で定義する(:制約クエリ実行時にまとめて呼び出すときにinterfaceを利用しているから)
func (User) FKs() []string {
	return []string{ // "ALTER TABLE {TABLE_NAME} ADD FOREIGN KEY ({F_KEY}) REFERENCES {REF_TABLE_NAME}({REF_KEY}) ON DELETE CASCADE ON UPDATE CASCADE",

	}
}
