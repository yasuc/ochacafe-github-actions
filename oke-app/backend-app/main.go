package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/tniita/ochacafe-github-actions/oke-app/backend-app/db"
	"github.com/tniita/ochacafe-github-actions/oke-app/backend-app/http"
)

func main() {
	db.SetupDB()
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		// アクセスを許可したいアクセス元
		AllowOrigins: []string{"*"},
		// アクセスを許可したいHTTPメソッド(以下の例だとPUTやDELETEはアクセスできません)
		AllowMethods: []string{
			"POST",
			"GET",
			"DELETE",
		},
		// 許可したいHTTPリクエストヘッダ
		AllowHeaders: []string{
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
		},
		// cookieなどの情報を必要とするかどうか
		AllowCredentials: true,
		// preflightリクエストの結果をキャッシュする時間
		MaxAge: 24 * time.Hour,
	}))
	router.GET("/items", http.GetAll)
	router.GET("/items/:id", http.GetItemById)
	router.POST("/items", http.UpdateItem)
	router.DELETE("/items/:id", http.DeleteItem)
	router.Run()
}
