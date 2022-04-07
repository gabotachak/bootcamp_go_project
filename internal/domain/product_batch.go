package domain

import "time"

type ProductBatch struct {
	ID                 int       `json:"id"`
	BatchNumber        string    `json:"batch_number"`
	CurrentQuantity    int       `json:"current_capacity"`
	CurrentTemperature float64   `json:"current_temperature"`
	DueDate            time.Time `json:"due_date"`
	InitialQuantity    int       `json:"initial_capacity"`
	ManufacturingDate  time.Time `json:"manufacturing_date"`
	ManufacturingHour  time.Time `json:"manufacturing_hour"`
	MinimumTemperature float64   `json:"minimum_temperature"`
	ProductId          int       `json:"product_id"`
	SectionId          int       `json:"section_id"`
}

type ReportProductBatch struct {
	SectionId     int `form:"id"`
	SectionNumber string
	Quantity      int
}
