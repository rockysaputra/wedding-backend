package model

import "gorm.io/gorm"

type WeddingBonusPackage struct {
	gorm.Model
	WeddingPackageID uint `json:"wedding_package_id"`
	BonusPackageID   uint `json:"bonus_package_id"`
}
