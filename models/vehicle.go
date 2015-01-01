package models

import (
	"time"
	"github.com/jinzhu/gorm"
	"fmt"
)

type Vehicle struct {
	Id 			int64 		`db:"id"`
	Plate 		string 		`db:"plate`
	ModelId 	int64		`db:"model_id"`
	CreatedAt 	time.Time	`db:"created_at"`
	UpdatedAt 	time.Time	`db:"updated_at"`
}

func NewVehicle(plate string) Vehicle {
	return Vehicle {
		Plate: plate,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func CreateVehicle(vehicle *Vehicle, db *gorm.DB) {
	if db.NewRecord(vehicle) {
		db.Create(vehicle)
	}
}

func SearchVehicleById(id int64, db *gorm.DB) *Vehicle {
	var result Vehicle
	err := db.Where("id = ?", id).First(&result)
	if err.Error != nil {
		fmt.Println(err.Error)
		return nil
	}
	return &result
}
