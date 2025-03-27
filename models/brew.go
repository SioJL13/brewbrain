package models

import (
	"time"

	"github.com/siojl13/brewbrain/db"
)

type Brew struct {
	ID             int64
	CoffeeName     string
	CoffeeType     string  // honey, natural, lavado
	CoffeeGrams    float64 `binding:"required"`
	GrindSize      int     `binding:"required"` // numeric grind size
	WaterGrams     float64 `binding:"required"`
	BrewingMethod  string  `binding:"required"` // kalita, breville, aeropress, etc
	BrewTime       int     `binding:"required"` //in seconds
	ExtractionTime int     // in seconds
	WaterTemp      float64 // Celcius
	GrinderType    string  // breville, timemore
	CreatedAt      time.Time
}

func (b Brew) Save() error {
	query := `
	INSERT INTO brews(
		coffeeName, 
		coffeeType, 
		coffeeGrams, 
		grindSize,
		waterGrams,
		brewingMethod,
		brewTime,
		extractionTime,
		waterTemp,
		grinderType,
		createdAt
	) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, CURRENT_TIMESTAMP);`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()
	res, err := stmt.Exec(b.CoffeeName, b.CoffeeType, b.CoffeeGrams,
		b.GrindSize, b.WaterGrams, b.BrewingMethod, b.BrewTime,
		b.ExtractionTime, b.WaterTemp, b.GrinderType)

	if err != nil {
		return err
	}

	_, err = res.LastInsertId()

	return err
}

func GetAllBrews() ([]Brew, error) {
	query := `SELECT * FROM brews`

	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var brews []Brew

	for rows.Next() {
		var brew Brew
		err := rows.Scan(&brew.ID, &brew.CoffeeName,
			&brew.CoffeeType, &brew.CoffeeGrams,
			&brew.GrindSize, &brew.WaterGrams,
			&brew.BrewingMethod, &brew.BrewTime,
			&brew.ExtractionTime, &brew.WaterTemp, &brew.GrinderType, &brew.CreatedAt)

		if err != nil {
			return nil, err
		}

		brews = append(brews, brew)
	}

	return brews, nil
}

func GetBrewByID(id int64) (*Brew, error) {
	query := "SELECT * FROM brews WHERE id = ?"
	row := db.DB.QueryRow(query, id)

	//TODO: check sqlx for direct mapping
	var brew Brew
	err := row.Scan(&brew.ID, &brew.CoffeeName,
		&brew.CoffeeType, &brew.CoffeeGrams,
		&brew.GrindSize, &brew.WaterGrams,
		&brew.BrewingMethod, &brew.BrewTime,
		&brew.ExtractionTime, &brew.WaterTemp, &brew.GrinderType, &brew.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &brew, nil
}
