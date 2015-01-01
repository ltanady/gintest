package resources

import (
	"github.com/jinzhu/gorm"
	"gintest/models"
	"github.com/gin-gonic/gin"
	"time"
	"bytes"
	"strconv"
)

type VehicleMakeResource struct {
	DB gorm.DB
}

func (vmr *VehicleMakeResource) CreateMake(c *gin.Context) {
	var make models.VehicleMake
	c.Bind(&make)
	make.CreatedAt = time.Now()
	make.UpdatedAt = make.CreatedAt

	vmr.DB.Create(&make)
	if !vmr.DB.NewRecord(make) {
		c.JSON(201, "Created")
	} else {
		c.JSON(200, "Failed to create")
	}
}

func (vmr *VehicleMakeResource) UpdateMake(c *gin.Context) {
	var make models.VehicleMake
	if id, err := strconv.Atoi(c.Params.ByName("id")); err == nil {
		if vmr.DB.Where("id = ?", id).First(&make).RecordNotFound() {
			c.JSON(404, "Invalid make id")
		} else {
			c.Bind(&make)
			vmr.DB.Save(&make)
			c.JSON(200, make)
		}
	} else {
		c.JSON(404, "Invalid id type")
	}
}

func (vmr *VehicleMakeResource) AllMakes(c *gin.Context) {
	var makes []models.VehicleMake
	vmr.DB.Find(&makes)
	if len(makes) > 0 {
		c.JSON(200, makes)
	} else {
		c.JSON(200, "{}")
	}
}

func (vmr *VehicleMakeResource) FindMakeByName(c *gin.Context) {
	var make models.VehicleMake
	var buffer bytes.Buffer
	name := c.Params.ByName("name")
	buffer.WriteString("%")
	buffer.WriteString(name)
	buffer.WriteString("%")
	if vmr.DB.Where("name LIKE ?", buffer.String()).First(&make).RecordNotFound() {
		c.JSON(200, "{}")
	} else {
		c.JSON(200, make)
	}
}
