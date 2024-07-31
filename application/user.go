// サービスのユースケースを書く(:処理の流れ)

package application

import "ddd/domain"

// アプリケーション層の構造体
type UserService struct {
	l *domain.UserLogic // ビジネスロジック
}

// ファクトリー関数
func NewUserService(l *domain.UserLogic) *UserService {
	return &UserService{l: l}
}

// ユーザー登録
func (s *UserService) RegisterUserService() {
	s.l.CreateUser(domain.User{})
}
