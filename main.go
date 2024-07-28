package main // package

import (
	"ddd/common/logging" // main method

	"github.com/gin-gonic/gin"
)

func main() {
	// 初期化処理
	initInstances, err := Init() // add "initInstances, " when changing to ddd
	if err != nil {
		return
	}
	// 破棄処理
	defer logging.LogFile().Close() // defer文でこの関数終了時に破棄
	logging.SuccessLog("Successful server init process.")

	// 鯖起動
	err = initInstances.Container.Invoke( // 依存性注入コンテナから必要な依存解決を解決し、渡されたコールバック関数にcontainerが持つエンジンの実体を渡す
		func(r *gin.Engine) { // containerが持つエンジンを受け取り鯖を起動
			r.Run(":4561") // 指定したポートで鯖起動
		},
	)
	if err != nil {
		logging.ErrorLog("Failed to start server", err)
		panic(err)
	}
}
