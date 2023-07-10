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
		AllowOrigins: []string{"*"},
		AllowMethods: []string{
			"POST",
			"GET",
			"DELETE",
		},
		AllowHeaders: []string{
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
		},
		AllowCredentials: true,
		MaxAge:           24 * time.Hour,
	}))
	router.GET("/items", http.GetAll)
	router.GET("/items/:id", http.GetItemById)
	router.POST("/items", http.UpdateItem)
	router.DELETE("/items/:id", http.DeleteItem)
	router.Run()
}
