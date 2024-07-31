// ユーザーインターフェース(:リクエストの受け取りとレスポンスの返却)

package presentation

import (
	"ddd/common/logging"
	"fmt"
	"net/http"
	"time"
	"unicode/utf8"

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

// test
func Test(c *gin.Context) {
	passes := []string{
		"aaaaaaaaaaa",
		"aaaaaaaaaaaa",
		"aaaaaaaaaaaaa",
		"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
		"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
		"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",

		"あああああああああああああああああああああああ",
		"ああああああああああああああああああああああああ",
		"あああああああああああああああああああああああああ",
	}

	for _, pass := range passes {
		fmt.Print(pass, ": ", len(pass), "\n")
		fmt.Print(pass, ": ", utf8.RuneCountInString(pass), "\n")
		fmt.Println()
	}
	// emails := []string{
	// 	"hoge@gmail.com",
	// 	"piyo.ta@gmail.com",
	// 	"piyo-ta@gamil.com",
	// 	"tyu320v9",
	// 	"8898@g.c",
	// 	"---@g.com",
	// 	"hoge@piyo",
	// 	"..@a",
	// 	"a@.",
	// }

	// for _, email := range emails {
	// 	_, err := mail.ParseAddress(email)
	// 	if err != nil {
	// 		fmt.Println(email + ": " + "no")
	// 	} else {
	// 		fmt.Println(email + ": " + "ok")
	// 	}
	// }
}
