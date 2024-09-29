package model

import (
	"github.com/gosimple/slug"
	"gorm.io/gorm"
)

type WeddingOrganizer struct {
	gorm.Model
	Name            string           `json:"name"`
	Slug            string           `json:"slug"`
	Icon            string           `json:"icon"`
	Phone           string           `json:"phone"`
	WeddingPackages []WeddingPackage `gorm:"foreignKey:WeddingOrganizerID" json:"wedding_packages"`
}

func (wo *WeddingOrganizer) BeforeSave(tx *gorm.DB) (err error) {
	wo.Slug = slug.Make(wo.Name)
	return
}
