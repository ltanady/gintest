package services

import (
	"github.com/jinzhu/gorm"
	"gintest/models"
	"github.com/gin-gonic/gin"
	"gintest/resources"
	_ "github.com/lib/pq"
	"os"
	"fmt"
)

type VehicleService struct {

}

func initDB() gorm.DB {
	db_config := "user=" + os.Getenv("HEROKU_POSTGRES_USERNAME") +
				 " password=" + os.Getenv("HEROKU_POSTGRES_PASSWORD") +
			     " host=" + os.Getenv("HEROKU_POSTGRES_HOST") +
				 " port=" + os.Getenv("HEROKU_POSTGRES_PORT") +
				 " dbname=" + os.Getenv("HEROKU_POSTGRES_DBNAME") +
				 " sslmode=require"
	db, err := gorm.Open("postgres", db_config)
	//db, err := gorm.Open("mysql", "root:@/temp?charset=utf8&parseTime=True")
	if err != nil {
		println("Error opening database connection")
		fmt.Println(err)
	}
	db.DB()

	db.DB().Ping()
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	db.SingularTable(true)

	db.AutoMigrate(&models.VehicleMake{}, &models.VehicleModel{})

	return db
}

func (s *VehicleService) Run() {
	port := ":"
	port += os.Getenv("PORT")
	db := initDB()
	defer db.DB().Close()
	makeResource := &resources.VehicleMakeResource{DB: db}
	modelResource := &resources.VehicleModelResource{DB: db}

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
			c.String(200, "Hello World")
		})

	// Make
	router.GET("/makes", makeResource.AllMakes)
	router.GET("/makes/search/:name", makeResource.FindMakeByName)
	router.POST("/makes", makeResource.CreateMake)
	router.PUT("/makes/:id", makeResource.UpdateMake)

	// Model
	router.GET("/models", modelResource.AllModels)
	router.GET("/models/:id", modelResource.FindModelById)
	router.POST("/models", modelResource.CreateModel)
	router.PUT("/models/:id", modelResource.UpdateModel)
	router.Run(port)
}
