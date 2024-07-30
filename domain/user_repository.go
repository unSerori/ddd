// リポジトリインターフェース

package domain

import (
	"mime/multipart"
	"os"
)

// リポジトリインターフェース
type UserRepository interface {
	// ファイルディレクトリ操作
	CreateDstDir(dst string, fileMode os.FileMode) error    // 指定されたディレクトリを作成
	UpLoadImage(filePath string, file multipart.File) error // ファイルをディレクトリに保存

	// ORM操作
	CreateUser(user User) (string, error) // ユーザー登録
	// ユーザー取得
}
