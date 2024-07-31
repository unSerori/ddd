// ロジックから呼び出される具体的な永続化処理
package infrastructure

import (
	"ddd/domain"
	"fmt"
	"mime/multipart"
	"os"

	"xorm.io/xorm"
)

// インフラストラクチャ層の構造体
type UserPersistence struct {
	db *xorm.Engine
}

// ファクトリー関数
func NewUserPersistence(db *xorm.Engine) domain.UserRepository {
	return &UserPersistence{db: db} // 返す構造体インスタンスのメソッドは、ファクトリー関数の返り血のインターフェースをすべて実装している(implements)ので、型が違うが無問題
}

// このエンティティのリポジトリインターフェースをすべて実装

// ファイルディレクトリ操作

// 指定されたディレクトリを作成
func (p *UserPersistence) CreateDstDir(dst string, fileMode os.FileMode) error {
	return nil
}

// ファイルをディレクトリに保存
func (p *UserPersistence) UpLoadImage(filePath string, file multipart.File) error {
	return nil
}

// ORM操作

// ユーザー登録
func (p *UserPersistence) CreateUser(user domain.User) (string, error) {
	fmt.Println("com cfm: OK!")
	return "", nil
}
