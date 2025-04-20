package models

import (
	"fmt"
	"time"

	"github.com/siojl13/brewbrain/db"
)

type Brew struct {
	ID             int64 `gorm:"primaryKey"`
	CoffeeName     string
	CoffeeType     string
	CoffeeGrams    float64
	GrindSize      int
	WaterGrams     float64
	BrewingMethod  string
	BrewTime       int
	ExtractionTime int
	WaterTemp      float64
	GrinderType    string
	CreatedAt      time.Time `gorm:"autoCreateTime"`
}

func (b Brew) Save() error {
	tx := db.DB.Create(&b)

	if tx.Error != nil {
		fmt.Println("here")
		return tx.Error
	}

	return nil
}

func (b *Brew) Update(fields Brew) error {
	return db.DB.
		Model(&Brew{}).
		Where("id = ?", b.ID).
		Updates(fields).
		Error
}

func (b Brew) Delete() error {
	tx := db.DB.Delete(&b)

	return tx.Error
}

func GetAllBrews() ([]Brew, error) {
	var brews []Brew

	rows := db.DB.Find(&brews)

	return brews, rows.Error
}

func GetBrewByID(id int64) (*Brew, error) {
	var brew Brew
	tx := db.DB.First(&brew, id)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return &brew, nil
}
