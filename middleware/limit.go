package middleware

import (
	"bytes"
	"ddd/common/logging"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// リクエストボディ容量を制限するミドルウェア
func LimitReqBodySize(maxBytesSize int64) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Printf("maxBytesSize: %v\n", maxBytesSize)
		// Content-Lengthからの確認
		if ctx.Request.ContentLength > maxBytesSize {
			// エラーログ
			logging.ErrorLog("Multipart request bodies are too big.", errors.New("request size "+strconv.Itoa(int(maxBytesSize))+"bytes"))
			// レスポンス
			resStatusCode := http.StatusRequestEntityTooLarge
			ctx.JSON(resStatusCode, gin.H{
				"srvResMsg":  http.StatusText(resStatusCode),
				"srvResData": gin.H{},
			})
			ctx.Abort() // リクエスト処理を中止
			return
		}

		// リクエストのボディを実際に読み込んで判定
		buf := make([]byte, maxBytesSize)             // 制限の分だけ読み込めるバッファを用意
		n, err := io.ReadFull(ctx.Request.Body, buf)  // バッファの容量分だけ読み込む  // nは読み込めたバイト数
		if err != nil && err != io.ErrUnexpectedEOF { // EOF以外のエラーが発生した場合は内部エラー
			// エラーログ
			logging.ErrorLog("Internal Server Error.", err)
			// レスポンス
			resStatusCode := http.StatusInternalServerError
			ctx.JSON(resStatusCode, gin.H{
				"srvResMsg":  http.StatusText(resStatusCode),
				"srvResData": gin.H{},
			})
			ctx.Abort() // リクエスト処理を中止
			return
		}
		if int64(n) == maxBytesSize && err == nil { // 制限サイズと同じまで読み込めてしまったら413
			// エラーログ
			logging.ErrorLog("Payload Too Large.", err)
			// レスポンス
			resStatusCode := http.StatusRequestEntityTooLarge
			ctx.JSON(resStatusCode, gin.H{
				"srvResMsg":  http.StatusText(resStatusCode),
				"srvResData": gin.H{},
			})
			ctx.Abort() // リクエスト処理を中止
			return
		}

		// 読み取ったデータをリクエストボディに再設定
		ctx.Request.Body = io.NopCloser(bytes.NewBuffer(buf[:n]))

		ctx.Next() // エンドポイントの処理に移行
	}
}
