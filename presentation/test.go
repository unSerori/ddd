package presentation

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// /
func ShowRootPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"topTitle":  "Route /",                                                            // ヘッダ内容 h1
		"mainTitle": "Hello.",                                                             // メインのタイトル h2
		"time":      time.Now(),                                                           // 時刻
		"message":   "This is an API server written in Golang for safety check purposes.", // message
	})
}
