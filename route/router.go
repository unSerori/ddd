package route

import (
	"ddd/common/logging"
	"ddd/middleware"
	"ddd/presentation"

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

// ファイルを設定
func loadingStaticFile(engine *gin.Engine) {
	// テンプレートと静的ファイルを読み込む
	engine.LoadHTMLGlob("view/*.html")
	engine.Static("/styles", "./view/styles") // クライアントがアクセスするURL, サーバ上のパス
	engine.Static("/scripts", "./view/scripts")
	logging.SuccessLog("Routing completed, start the server.")
}

// エンジンを作成して返す
func SetupRouter(handlers Handlers) (*gin.Engine, error) {
	// エンジンを作成
	engine := gin.Default()

	// マルチパートフォームのメモリ使用制限を設定
	engine.MaxMultipartMemory = 8 << 20 // 20bit左シフトで8MiB

	// ルーティング
	routing(engine, handlers)

	// 静的ファイル設定
	loadingStaticFile(engine)

	// router設定されたengineを返す。
	return engine, nil
}
