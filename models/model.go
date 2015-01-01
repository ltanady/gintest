package models

import (
	"time"
)

type VehicleModel struct {
	Id	 		int64 		`json:"id"`
	Name 		string 		`sql:"not null;unique" json:"name" binding:"required"`
	Year 		string 		`sql:"not_null" json:"year" binding:"required"`
	MakeId 		int64 		`sql:"not null" json:"make_id" binding:"required"`
	Make		VehicleMake `json:"make"`
	CreatedAt 	time.Time	`json:"-"`
	UpdatedAt 	time.Time	`json:"-"`
}

func NewModel(name string, year string, make VehicleMake) VehicleModel {
	return VehicleModel {
		Name: name,
		Year: year,
		//MakeId: make.Id,
		Make: make,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

//func CreateModel(model *VehicleModel, db *gorm.DB) {
//	if db.NewRecord(model) {
//		db.Create(model)
//	}
//}
//
//func SearchModelById(id int64, db *gorm.DB) *VehicleModel {
//	var result VehicleModel
//	err := db.Where("id = ?", id).First(&result)
//	if err.Error != nil {
//		fmt.Println(err.Error)
//		return nil
//	}
//	return &result
//}
