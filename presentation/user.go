// ユーザーインターフェース(:リクエストの受け取りとレスポンスの返却)のハンドラー

package presentation

import (
	"ddd/application"

	"github.com/gin-gonic/gin"
)

// プレゼンテーション層の構造体
type UserHandler struct {
	s *application.UserService // 依存先層の構造体 依存先層のポインタ型
}

// ファクトリー関数 // これと構造体で各層をつなぐ
func NewUserHandler(s *application.UserService) *UserHandler {
	return &UserHandler{s: s}
}

// ユーザー登録
func (h *UserHandler) RegisterUserHandler(c *gin.Context) {
	h.s.RegisterUserService()
}
