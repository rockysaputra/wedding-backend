package model

import (
	"github.com/gosimple/slug"
	"gorm.io/gorm"
)

type WeddingPackage struct {
	gorm.Model
	Name                 string                `json:"name"`
	Slug                 string                `json:"slug"`
	Thumbnail            string                `json:"thumbnail"`
	About                string                `json:"about"`
	Price                float64               `json:"price"`
	IsPopular            bool                  `json:"is_popular"`
	CityID               uint                  `json:"city_id"`
	WeddingOrganizerID   uint                  `json:"wedding_organizer_id"`
	City                 City                  `gorm:"foreignKey:CityID" json:"city"`
	WeddingOrganizer     WeddingOrganizer      `gorm:"foreignKey:WeddingOrganizerID" json:"wedding_organizer"`
	WeddingBonusPackages []WeddingBonusPackage `gorm:"foreignKey:WeddingPackageID" json:"wedding_bonus_package"`
	WeddingPhotos        []WeddingPhoto        `gorm:"foreignKey:WeddingPackagePhotoID" json:"wedding_bonus_package_photo"`
	WeddingTestimonials  []WeddingTestimonial  `gorm:"foreignKey:WeddingPackageTestimonialID" json:"wedding_testimonial"`
	WeddingTransactions  []BookingTransaction  `gorm:"foreignKey:WeddingPackageTransactionID" json:"wedding_package"`
}

func (wp *WeddingPackage) BeforeSave(tx *gorm.DB) (err error) {
	wp.Slug = slug.Make(wp.Name)
	return
}
