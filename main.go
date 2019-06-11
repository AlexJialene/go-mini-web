package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func main() {
	var router = gin.Default()
	router.LoadHTMLGlob("temp/*")
	router.Static("/static", "./static")

	bindApi(router)
	mysqlDB, _ := ConnectMySql()
	db = mysqlDB
	router.Run(":9000")
}

func bindApi(router *gin.Engine) {
	router.GET("/index", indexHtml)

	router.POST("/queryPoint", queryValue)
}

func indexHtml(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "this is title",
	})
}

func queryValue(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"ret": "success",
	})
}
