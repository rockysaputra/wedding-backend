package model

import (
	"github.com/gosimple/slug"
	"gorm.io/gorm"
)

type City struct {
	gorm.Model
	Name            string           `json:"name"`
	Slug            string           `json:"slug"`
	Icon            string           `json:"icon"`
	WeddingPackages []WeddingPackage `gorm:"foreignKey:CityID" json:"wedding_packages"`
}

func (City) TableName() string {
	return "cities"
}

func (city *City) BeforeSave(tx *gorm.DB) (err error) {
	city.Slug = slug.Make(city.Name)
	return
}
