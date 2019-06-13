package main

import (
	"fmt"
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

	//fmt.Println(QueryRecordsLastTimeStr(db))
	router.Run(":9000")
}

func bindApi(router *gin.Engine) {
	router.GET("/index", indexHtml)

	router.GET("/queryValue", queryValue)
}

func indexHtml(c *gin.Context) {
	list := QueryItemsLimit100(db)
	fmt.Println(len(list))
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "your title",
		"list":  list,
	})
}

func queryValue(c *gin.Context) {
	nick := c.DefaultQuery("nick", "0")
	if nick == "0" {
		//return fail
		c.JSON(http.StatusOK, gin.H{
			"ret": "fail",
		})
	}
	value := QueryValueByName(db, nick)
	c.JSON(http.StatusOK, gin.H{
		"ret":   "success",
		"nick":  nick,
		"value": value.Value,
	})
}
