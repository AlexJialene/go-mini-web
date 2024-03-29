package main

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const (
	USERNAME = "root"
	PASSWORD = ""
	NETWORK  = "tcp"
	SERVER   = "47.106.205.186"
	PORT     = 3306
	DATABASE = "records?charset=utf8&parseTime=True&loc=Local"
)

func ConnectMySql() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
	db, err := gorm.Open("mysql", dsn)
	//defer db.Close()
	if err != nil {
		fmt.Println("[ERROR] | connect mysql error:%v", err)
		return nil, err
	}
	db.DB().SetConnMaxLifetime(100 * time.Second) //最大连接周期，超过时间的连接就close
	db.DB().SetMaxOpenConns(100)                  //设置最大连接数
	db.DB().SetMaxIdleConns(16)                   //设置闲置连接数
	db.Debug()
	db.LogMode(true)

	return db, nil

}

func QueryValueByName(db *gorm.DB, name string) Tbl_Item_Records {
	var result Tbl_Item_Records
	db.Table("tbl_item_records").Select("value").Where("nick = ?", name).Scan(&result)
	//db.Model(&Tbl_Item_Records{}).Select("value").Where("nick = ?", name).Scan(&result)
	//db.Raw("SELECT value FROM tbl_item_records WHERE nick = ?", name).Scan(&result)
	return result
}

func QueryItemsLimit100(db *gorm.DB) []Tbl_Item_Records {
	var array []Tbl_Item_Records
	db.Model(&Tbl_Item_Records{}).Order("value desc").Limit(100).Find(&array)

	return array
}

func QueryRecordsLastTimeStr(db *gorm.DB) interface{} {
	//var vo Table
	//db.Table("information_schema.TABLES").Select("UPDATE_TIME").Where("information_schema.TABLES.TABLE_SCHEMA=? AND information_schema.TABLES.TABLE_NAME=?", "records", "tbl_item_records").Scan(&vo)

	//todo Gorm does not seem to support cross-database queries ? And does anyone know the solution?

	//var updateTime time.Time
	//db.Raw("select UPDATE_TIME from information_schema.TABLES where information_schema.TABLES.TABLE_SCHEMA=? AND information_schema.TABLES.TABLE_NAME=?", "records", "tbl_item_records").Scan(&updateTime)
	//fmt.Println(updateTime)
	//return updateTime

	return nil
}

type TimeVo struct {
	id          int
	UPDATE_TIME time.Time
}

type Table struct {
	Update_time time.Time
}
