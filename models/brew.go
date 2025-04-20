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

func (b Brew) Update() error {
	db.DB.First(&b)
	tx := db.DB.Save(&Brew{
		CoffeeName: b.CoffeeName,
	})

	// query := `
	// UPDATE brews SET
	// 	coffeeName = ?,
	// 	coffeeType = ?,
	// 	coffeeGrams = ?,
	// 	grindSize = ?,
	// 	waterGrams = ?,
	// 	brewingMethod = ?,
	// 	brewTime = ?,
	// 	extractionTime = ?,
	// 	waterTemp = ?,
	// 	grinderType = ?
	//  WHERE id = ?;`

	// stmt, err := db.DB.Prepare(query)
	// if err != nil {
	// 	return err
	// }

	// defer stmt.Close()
	// _, err = stmt.Exec(b.CoffeeName, b.CoffeeType, b.CoffeeGrams,
	// 	b.GrindSize, b.WaterGrams, b.BrewingMethod, b.BrewTime,
	// 	b.ExtractionTime, b.WaterTemp, b.GrinderType, b.ID)

	return tx.Error
}

func (b Brew) Delete() error {
	tx := db.DB.Delete(&b)

	return tx.Error
}

// func GetAllBrews() ([]Brew, error) {
// 	query := `SELECT * FROM brews`

// 	rows, err := db.DB.Query(query)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var brews []Brew

// 	for rows.Next() {
// 		var brew Brew
// 		err := rows.Scan(&brew.ID, &brew.CoffeeName,
// 			&brew.CoffeeType, &brew.CoffeeGrams,
// 			&brew.GrindSize, &brew.WaterGrams,
// 			&brew.BrewingMethod, &brew.BrewTime,
// 			&brew.ExtractionTime, &brew.WaterTemp, &brew.GrinderType, &brew.CreatedAt)

// 		if err != nil {
// 			return nil, err
// 		}

// 		brews = append(brews, brew)
// 	}

// 	return brews, nil
// }

func GetBrewByID(id int64) (*Brew, error) {
	var brew Brew
	tx := db.DB.First(&brew, id)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return &brew, nil
}
