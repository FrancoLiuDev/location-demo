package main

import (
	"time"
)

type Location struct {
	Lat  float32
	Long float32
}

type Data struct {
	ID        string    `db:"id"`
	Location  Location  `db:"location"`
	DateAdded time.Time `db:"dateAdd"`
}
