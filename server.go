package main

import (
	//"github.com/gin-gonic/gin"
	//"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"log"
	//"time"
	//"gintest/models"
	"gintest/services"
)

//func initDB() gorm.DB {
//	db, err := gorm.Open("mysql", "root:@/temp?charset=utf8&parseTime=True")
//	if err != nil {
//		println("Error opening mysql connection")
//	}
//	db.DB()
//
//	db.DB().Ping()
//	db.DB().SetMaxIdleConns(10)
//	db.DB().SetMaxOpenConns(100)
//
//	db.SingularTable(true)
//
//	db.AutoMigrate(&models.VehicleMake{}, &models.VehicleModel{})
//
//	return db
//}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

func main() {
//	db := initDB()
//	defer db.DB().Close()
//
//	router := gin.Default()
//	router.GET("/", func(c *gin.Context) {
////			newMake := models.NewMake("Mercedes")
////			models.CreateMake(&newMake, &db)
////			newModel := models.NewModel("E-300", "2011", &newMake)
////			models.CreateModel(&newModel, &db)
//			make := models.SearchMakeById(1, &db)
//			if make != nil {
//				println(make.Name)
//			}
//			c.String(200, "Hello World")
//		})
//
//	router.GET("/user/:name", func(c *gin.Context) {
//			name := c.Params.ByName("name")
//			message := "Hello " + name
//			c.String(200, message)
//		})

//	router.Run(":8080")
	svc := services.VehicleService{}
	svc.Run()
}
