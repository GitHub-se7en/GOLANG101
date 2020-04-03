package main

import (
	"GOLANG101/JWT04/config"
	"GOLANG101/JWT04/models"
	"GOLANG101/JWT04/routes"
	"fmt"
	_ "github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm"
	_ "gopkg.in/mgo.v2"
)

func main() {
	//db, err := gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")

	// Creating a connection to the database
	var err error
	//这里需要注意的是，在这里已经给config的DB赋值了，
	config.DB, err = gorm.Open("mysql", config.DbURL(config.BuildDBConfig()))

	if err != nil {
		fmt.Println("statuse: ", err)
	}

	defer config.DB.Close()

	// run the migrations: todo struct
	config.DB.AutoMigrate(&models.Todo{})

	//setup routes
	r := routes.SetupRouter()

	// running
	r.Run()

}
