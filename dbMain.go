package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func getTableLoc() *sqlx.DB {
	db, err := sqlx.Connect("mysql", locDbLocation)
	if err != nil {
		panic(err)
	}
	return db
}

func dbLocCreate() {
	db, err := sqlx.Connect("mysql", locDbAdmin)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	_, err = db.Exec("CREATE DATABASE MAIN ")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("USE MAIN")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("CREATE TABLE LOCATION ( id VARCHAR(30), location json, dateAdd datetime)")

	if err != nil {
		panic(err)
	}
}

func dbLocDropAll() {
	db, err := sqlx.Connect("mysql", locDbAdmin)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("DROP SCHEMA IF EXISTS MAIN")
	if err != nil {
		panic(err)
	}
}
