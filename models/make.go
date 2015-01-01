package models

import (
	"time"
)

type VehicleMake struct {
	Id 	 		int64 		`json:"id"`
	Name 		string 		`sql:"not null;unique" json:"name" binding:"required"`
	CreatedAt 	time.Time	`json:"-"`
	UpdatedAt 	time.Time	`json:"-"`
}

func NewMake(name string) VehicleMake {
	return VehicleMake {
		Name: name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
