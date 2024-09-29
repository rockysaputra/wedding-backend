package model

import "gorm.io/gorm"

type WeddingTestimonial struct {
	gorm.Model
	Name                        string `json:"name"`
	Photo                       string `json:"photo"`
	WeddingPackageTestimonialID uint   `json:"wedding_package_id"`
	Message                     string `json:"message"`
	Occupation                  string `json:"occupation"`
}
