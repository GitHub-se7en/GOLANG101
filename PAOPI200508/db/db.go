package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var (
	//这个变量还能跨包使用呢！！！！！！
	dbConn *sql.DB
	err    error
)

func init() {
	log.Println("init数据库")
	dbConn, err = sql.Open("mysql", "root:root@tcp(localhost:3306)/golang?charset=utf8&parseTime=True&loc=Local")
	/**
	1. 数据库的限制，最大就是150，这个是开启一个sql.Open的最大限制，如果是集群的话，可能会增加，但是就是不知道数据库能不能撑得住
	2. 这个open和idle的关系是什么？idle是池子里面的连接数，open是所有的开启的
	*/
	dbConn.SetMaxOpenConns(150)
	dbConn.SetMaxIdleConns(150)
	if err != nil {
		panic(err.Error())
	}
}
