package route

import (
	"ddd/middleware"
	"ddd/presentation"
	"ddd/view"

	"github.com/gin-gonic/gin"
)

// エンドポイントのルーティング
func routing(engine *gin.Engine, handlers Handlers) {
	// midLog all
	engine.Use(middleware.LoggingMid())

	// endpoints  // handlersを使って、作成済みの依存関係を利用

	// root page
	engine.GET("/", presentation.ShowRootPage) // /

	// confirmation
	engine.GET("/cfm_req", presentation.ConfirmationReq) // /cfm_req

	// test
	engine.GET("/test", presentation.Test) // /test

	// ver1グループ
	v1 := engine.Group("/v1")
	{
		// usersグループ
		users := v1.Group("/users")
		{
			// ユーザー登録
			users.POST("/register", handlers.UserHandler.RegisterUserHandler) // /v1/users/register
		}
	}
}

// エンジンを作成して返す
func SetupRouter(handlers Handlers) (*gin.Engine, error) {
	// エンジンを作成
	engine := gin.Default()

	// 静的ファイル設定
	err := view.LoadingStaticFile(engine)
	if err != nil {
		return nil, err
	}

	// マルチパートフォームのメモリ使用制限を設定
	engine.MaxMultipartMemory = 8 << 20 // 20bit左シフトで8MiB

	// ルーティング
	routing(engine, handlers)

	// router設定されたengineを返す。
	return engine, nil
}
