package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func main() {
	gin.SetMode(gin.ReleaseMode)
	var router = gin.Default()
	router.LoadHTMLGlob("templates/*")

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

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
	//value := QueryValueByName(db, "木木")
	//fmt.Println(value)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "this is title",
	})
}

func queryValue(c *gin.Context) {
	//get value

}
