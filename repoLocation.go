package main

import (
	"encoding/json"
	_ "fmt"
	"github.com/lithammer/shortuuid"
	"log"
	"time"
)

func dbLocInsert(data Data) {
	db := getTableLoc()
	query := "INSERT INTO LOCATION(id,location, dateAdd) VALUES(?,?,?)"
	str, _ := json.Marshal(data.Location)
	_, err := db.Exec(query, shortuuid.New(), str, time.Now().UTC())

	if err != nil {
		log.Fatalln(err)
	}
}

func dbLocSelect() []Data {
	db := getTableLoc()
	defer db.Close()
	locs := []Data{}

	rows, err := db.Queryx("SELECT * FROM LOCATION")
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		results := make(map[string]interface{})
		err = rows.MapScan(results)
		if err != nil {
			panic(err)
		}
		id := results["id"].([]byte)
		dateAdd := results["dateAdd"].(time.Time)
		loc := new(Location)
		json.Unmarshal([]byte(string(results["location"].([]byte))), loc)
		d := Data{ID: string(id), Location: *loc, DateAdded: dateAdd}
		locs = append(locs, d)
	}
	return locs
}
