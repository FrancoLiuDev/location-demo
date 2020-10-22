package main

import (
	"time"
)

type Data struct {
	ID       string `db:"id"`
	Location struct {
		Lat  float32
		Long float32
	} `db:"location"`
	DateAdded time.Time `db:"dateAdd"`
}
