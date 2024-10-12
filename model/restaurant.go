package model

import (
	"cinnanym/maps"
	"github.com/surrealdb/surrealdb.go"
)

type Restaurant struct {
	surrealdb.Basemodel `table:"restaurant"`
	Name                string   `json:"name"`
	Website             string   `json:"website"`
	Phone               string   `json:"phone"`
	Address             string   `json:"address"`
	City                string   `json:"city"`
	Country             string   `json:"country"`
	Tags                []string `json:"tags"`
	Type                string   `json:"type"`
	Rating              float64  `json:"rating"`
}

func (r *Restaurant) ToMap() maps.Map {
	return maps.Map{
		"name":    r.Name,
		"website": r.Website,
		"phone":   r.Phone,
		"address": r.Address,
		"city":    r.City,
		"country": r.Country,
		"tags":    r.Tags,
		"type":    r.Type,
		"rating":  r.Rating,
	}
}
