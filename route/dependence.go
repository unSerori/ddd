// 構造体インスタンスを用いた依存性注入をライブラリで管理
package route

import (
	"ddd/application"
	"ddd/common/logging"
	"ddd/domain"
	"ddd/infrastructure"
	"ddd/infrastructure/orm"
	"ddd/presentation"

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
	"xorm.io/xorm"
)

// ドメイン層のエンティティを追加したら、各層のファクトリー関数をprovidersスライスに追加し、ハンドラーの構造体をHandlersに追加

// 依存関係を登録したいファクトリー関数を追加していく
func providers() []interface{} {
	return []interface{}{
		// ユーザー登録
		domain.NewUserLogic,
		infrastructure.NewUserPersistence,
		application.NewUserService,
		presentation.NewUserHandler,
	}
}

// 依存設定を一括で行うための構造体(:ファクトリー関数でつながった依存関係の一番上の層)  // これをrouter設定側で使えば依存関係をながなが書かなくていい
type Handlers struct {
	dig.In // 継承

	// ハンドラーの構造体
	UserHandler *presentation.UserHandler // ユーザー登録
}

// 依存関係を登録
func BuildContainer() *dig.Container {
	// コンテナを作成
	container := dig.New()

	// ORMをコンテナに渡し、依存関係を登録
	container.Provide(
		func() *xorm.Engine {
			db, err := orm.InitDB() // 設定済みのORMインスタンスを無名関数でラップしたものを受け取り、コンテナにインスタンスを登録
			if err != nil {
				logging.ErrorLog("Couldnt receive router engine.", err) // エラー内容を出力し早期リターン
				panic(err)
			}

			return db
		},
	)

	// 登録する依存関係を取得
	providers := providers()

	// スライスから各項目の依存関係を登録し、エラーハンドリング
	for _, provider := range providers {
		if err := container.Provide(provider); err != nil {
			logging.ErrorLog("Dependency registration failed.", nil)
			panic(err)
		}
	}

	// ルーティング設定をコンテナに渡し、依存関係を登録
	container.Provide(
		func(handlers Handlers) *gin.Engine {
			router, err := SetupRouter(handlers) // router設定されたengineを無名関数でラップしたものを受け取り、ルーティングを登録
			if err != nil {
				logging.ErrorLog("Couldnt receive router engine.", err) // エラー内容を出力し早期リターン
				panic(err)
			}
			return router
		},
	)

	return container
}
