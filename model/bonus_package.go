package model

import (
	"github.com/gosimple/slug"
	"gorm.io/gorm"
)

type BonusPackage struct {
	gorm.Model
	Name                 string                `json:"name"`
	Slug                 string                `json:"slug"`
	Thumbnail            string                `json:"thumbnail"`
	About                string                `json:"about"`
	Price                uint                  `json:"price"`
	WeddingBonusPackages []WeddingBonusPackage `gorm:"foreignKey:BonusPackageID" json:"wedding_bonus_package"`
}

func (bonus *BonusPackage) BeforeSave(tx *gorm.DB) (err error) {
	bonus.Slug = slug.Make(bonus.Name)
	return
}
