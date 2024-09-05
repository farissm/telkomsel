package product

import (
	"gorm.io/gorm"
)

type Product struct {
    gorm.Model
	ID 				uint	`gorm:"primary_key;auto_increment" json:"id"`
	Name 			string	`gorm:"not null" json:"name"`
	Description 	string	`gorm:"not null" json:"description"`
    Price   		float64	`gorm:"not null" json:"price"`
	Variety   		string	`gorm:"not null" json:"variety"`
	Rating   		float32	`gorm:"not null" json:"rating"`
	Stock   		int 	`gorm:"not null" json:"stock"`
    TotalSold    	int 	`gorm:"not null" json:"total_sold"`
}

type ProductResponse struct {
	ID				uint	`json:"id"`
	Name 			string	`json:"name"`
	Description 	string	`json:"desription"`
    Price   		float64	`json:"price"`
	Variety   		string	`json:"variety"`
	Rating   		float32	`json:"rating"`
	Stock   		int 	`json:"stock"`
    TotalSold    	int 	`json:"total_sold"`
}