// ビジネスロジック
package domain

import (
	"ddd/common/logging"
	"ddd/utility/custom"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/mail"
	"os"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

// ビジネスロジックの構造体 // infrastructure層の処理を呼び出したいが、このままではdomain<->infrastructure間で循環参照してしまう。そのため実際の実装(infra)に依存するのではなく、提供元の同レイヤー内のrepository interface(実装された処理関数の型と呼べる)を利用することで、repositoryを介して具体的な実装(:infrastructure)を利用できる。
type UserLogic struct {
	r UserRepository
}

// ファクトリー関数
func NewUserLogic(r UserRepository) *UserLogic { // 依存先のインスタンスを受け取る  // 実体を返す
	return &UserLogic{r: r} // 構造体に依存先のインスタンスを設定
}

// ビジネスロジック // アプリケーション層からのエンティティへの影響をビジネスルールに従って実現する

// 年齢制限
func (l *UserLogic) AgeLimit(age int) error {
	ageLimit := 18
	if age > ageLimit { // 18以上
		return custom.NewErr(custom.ErrTypeTooYoung, custom.WithMsg(fmt.Sprintf("%d%s%d", age, ">", ageLimit)))
	}

	return nil
}

// メアドの形式
func (l *UserLogic) ValidMail(email string) error {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return err
	}

	return nil
}

// パスワードは12オクテット以上72オクテット以下
func (l *UserLogic) ValidPass(pass string) error {
	passCount := len(pass)
	if passCount > 72 {
		return custom.NewErr(custom.ErrTypeTooLongPass)
	}
	if passCount < 12 {
		return custom.NewErr(custom.ErrTypeTooShortPass)
	}

	return nil
}

// ユーザー情報のチェック
func (l *UserLogic) ValidUserInfo(user User) error {
	// 年齢の確認
	err := l.AgeLimit(user.Age)
	if err != nil {
		return err
	}

	// メアド形式の確認
	err = l.ValidMail(user.Mail)
	if err != nil {
		return err
	}

	// パスワード確認
	err = l.ValidPass(user.Pass)
	if err != nil {
		return err
	}

	return nil
}

// 画像のファイルサイズの制限(context由来)
func (l *UserLogic) ValidFileSize(file multipart.FileHeader) error {
	// ファイルサイズの制限
	var maxSize int64                                                              // 上限設定値
	maxSize = 5242880                                                              // default値10MB
	if maxSizeByEnv := os.Getenv("MULTIPART_IMAGE_MAX_SIZE"); maxSizeByEnv != "" { // 空文字でなければ数値に変換する
		var err error
		maxSizeByEnvInt, err := strconv.Atoi(maxSizeByEnv) // 数値に変換
		if err != nil {
			return err
		}
		maxSize = int64(maxSizeByEnvInt) // int64に変換
	}
	if file.Size > maxSize { // ファイルサイズと比較する
		return custom.NewErr(custom.ErrTypeFileSizeTooLarge)
	}

	return nil
}

// 画像の種類を特定し、拡張子を返す
func (l *UserLogic) ValidFileMime(reqFile multipart.FileHeader) (string, error) {
	// 許可されたMIMEタイプかどうかを確認、許可されていた場合は一致したタイプを返すネスト関数を定義
	validMime := func(mimetype string) (bool, string) {
		// 有効なファイルタイプを定義
		var allowedMimeTypes = []string{
			"image/png",
			"image/jpeg",
			"image/jpg",
			"image/gif",
		}

		// ひとつずつ比較
		for _, allowedMimeType := range allowedMimeTypes {
			if strings.EqualFold(allowedMimeType, mimetype) { // 大文字小文字を無視して文字列比較
				logging.InfoLog("True validMime", "True validMime/mimetype: "+mimetype+", allowedMimeType: "+allowedMimeType)
				return true, allowedMimeType // 一致した時点で早期リターン
			}
		}

		logging.InfoLog("False validMime", "False validMime/mimetype: "+mimetype)

		return false, ""
	}

	// 画像リクエストのContent-Typeから形式(png, jpg, jpeg, gif)の確認
	mimeType := reqFile.Header.Get("Content-Type") // リクエスト画像のmime typeを取得
	ok, _ := validMime(mimeType)                   // 許可されたMIMEタイプか確認
	if !ok {
		return "", custom.NewErr(custom.ErrTypeInvalidFileFormat, custom.WithMsg("the Content-Type of the request image is invalid"))
	}
	// ファイルのバイナリからMIMEタイプを推測し確認、拡張子を取得
	buffer := make([]byte, 512) // バイトスライスのバッファを作成
	file, err := reqFile.Open() // multipart.Formを実装するFileオブジェクトを直接取得  // このバイナリはファイルタイプの特定とファイル保存書き込み処理で使う
	if err != nil {
		return "", err
	}
	defer file.Close()                                 // 終了後破棄
	file.Read(buffer)                                  // ファイルをバッファに読み込む  // 読み込んだバイト数とエラーを返す
	mimeTypeByBinary := http.DetectContentType(buffer) // 読み込んだバッファからコンテントタイプを取得
	ok, validType := validMime(mimeTypeByBinary)       // 許可されたMIMEタイプか確認
	if !ok {
		return "", custom.NewErr(custom.ErrTypeInvalidFileFormat, custom.WithMsg("the Content-Type inferred from the request image binary is invalid"))
	}

	// 画像の種類を取得して拡張子として保存
	return strings.Split(validType, "/")[1], nil
}

// 画像保存
func (l *UserLogic) UploadIcon(form multipart.Form) (string, error) {
	// filesスライスからimage fieldsのひとつめを取得
	image := form.File["image"][0]

	// 保存先ディレクトリの確保
	dst := "./upload/t_material"
	_ = dst
	// l.r.CreateDstDir(dst, 0644)

	// バリデーション

	err := l.ValidFileSize(*image) // ファイルサイズの制限
	if err != nil {
		return "", err
	}

	fileExt, err := l.ValidFileMime(*image) // ファイルのMIMEチェックと拡張子取得
	if err != nil {
		return "", err
	}

	// ファイル名をuuidで作成
	fileNameWithoutExt, err := uuid.NewRandom() // 新しいuuidの生成
	if err != nil {
		return "", err
	}
	fileName := fileNameWithoutExt.String() + "." + fileExt // ファイルネームを生成
	filePath := dst + "/" + fileName                        // ファイルパスを生成

	//　保存
	file, err := image.Open()
	if err != nil {
		return "", err
	}
	err = l.r.UpLoadImage(filePath, file)
	if err != nil {
		return "", err
	}

	// 成果物返却
	return fileNameWithoutExt.String(), nil

}

// ユーザー登録
func (l *UserLogic) CreateUser(user User) (string, error) {
	id, err := l.r.CreateUser(user)
	if err != nil {
		return "", err
	}

	return id, nil
}
