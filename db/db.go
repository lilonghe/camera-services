package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/xormplus/xorm"
	"log"
)

var Store *xorm.Engine

func Init() {
	engine, err := xorm.NewEngine("mysql", "root:lilonghe@/services?charset=utf8")
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	Store = engine
	Store.Ping()
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	Store.ShowSQL(true)

}
