package model

type Catalog struct {
	ID       uint   `json:"id" gorm:"unique;primaryKey;autoIncrement"`
	Name     string `json:"prod_name" gorm:"default:null"`
	Price    int64  `json:"prod_price" gorm:"default:null"`
	Category string `json:"prod_category" gorm:"default:null"`
}
