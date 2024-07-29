package presentation

import (
	"ddd/common/logging"
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

// cfmreq
func ConfirmationReq(c *gin.Context) {
	logging.SimpleLog("method: ", c.Request.Method, "\n")
	logging.SimpleLog("url: ", c.Request.URL, "\n")
	// logging.SimpleLog("tls ver: ", c.Request.TLS.Version, "\n")
	logging.SimpleLog("header: ", c.Request.Header, "\n")
	logging.SimpleLog("body: ", c.Request.Body, "\n")
	logging.SimpleLog("url query: ", c.Request.URL.Query(), "\n")
	logging.SimpleLog("\n")
}
