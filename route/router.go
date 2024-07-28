package route

import "github.com/gin-gonic/gin"

// エンドポイントのルーティング
func routing(engine *gin.Engine, handlers Handlers) {
	// midLog all

	// endpoints  // handlersを使って、作成済みの依存関係を利用
}

// エンジンを作成して返す
func SetupRouter(handlers Handlers) (*gin.Engine, error) {
	// エンジンを作成
	engine := gin.Default()

	// マルチパートフォームのメモリ使用制限を設定
	engine.MaxMultipartMemory = 8 << 20 // 20bit左シフトで8MiB

	// ルーティング
	routing(engine, handlers)

	// router設定されたengineを返す。
	return engine, nil
}
