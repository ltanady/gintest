package resources

import (
	"github.com/jinzhu/gorm"
	"gintest/models"
	"github.com/gin-gonic/gin"
	"time"
	"encoding/json"
	"fmt"
	"strconv"
)

type VehicleModelResource struct {
	DB gorm.DB
}

func (vmr *VehicleModelResource) CreateModel(c *gin.Context) {
	var make models.VehicleMake
	var model models.VehicleModel
	c.Bind(&model)
	model.CreatedAt = time.Now()
	model.UpdatedAt = model.UpdatedAt
	decoder := json.NewDecoder(c.Request.Body)
	fmt.Println(decoder)

	if vmr.DB.Where("id = ?", 1).First(&make).RecordNotFound() {
		c.JSON(404, "Invalid make id")
	}
	model.Make = make

	vmr.DB.Create(&model)
	if !vmr.DB.NewRecord(model) {
		c.JSON(201, "Created")
	} else {
		c.JSON(200, "Failed to create")
	}
}

func (vmr *VehicleModelResource) UpdateModel(c *gin.Context) {
	var model models.VehicleModel
	if id, err := strconv.Atoi(c.Params.ByName("id")); err == nil {
		if vmr.DB.Where("id = ?", id).First(&model).RecordNotFound() {
			c.JSON(404, "Invalid model id")
		} else {
			c.Bind(&model)
			vmr.DB.Save(&model)
			c.JSON(200, model)
		}
	} else {
		c.JSON(40, "Invalid id type")
	}
}

func (vmr *VehicleModelResource) AllModels(c *gin.Context) {
	var models []models.VehicleModel
	vmr.DB.Find(&models)
	if len(models) > 0 {
		for index, model := range models {
			println(model.Id)
			vmr.DB.Model(&models[index]).Related(&models[index].Make, "MakeId")
		}
		c.JSON(200, models)
	} else {
		c.JSON(200, "{}")
	}
}

func (vmr *VehicleModelResource) FindModelById(c *gin.Context) {
	var model models.VehicleModel

	if id, err := strconv.Atoi(c.Params.ByName("id")); err == nil {
		if vmr.DB.Where("id = ?", id).First(&model).RecordNotFound() {
			c.JSON(404, "Invalid make id")
		}
		vmr.DB.Model(&model).Related(&model.Make, "MakeId")
		c.JSON(200, model)
	}

}
