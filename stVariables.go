package main

import (
	"fmt"
	"os"
)

var dbLocHost string = os.Getenv("SCENARIO_DBHOST")
var dbLocName string = os.Getenv("SCENARIO_DBNAME")
var dbLocPwd string = os.Getenv("SCENARIO_DBPWD")
var dbLocPort string = os.Getenv("SCENARIO_DBPPORT")
var locDbAdmin string = fmt.Sprintf("root:%s@(%s:%s)/?charset=utf8&parseTime=true&net_write_timeout=9000", dbLocPwd, dbLocHost, dbLocPort)
var locDbLocation string = fmt.Sprintf("root:%s@(%s:%s)/%s?charset=utf8&parseTime=true&net_write_timeout=9000", dbLocPwd, dbLocHost, dbLocPort, dbLocName)

func initVariables() {
	os.Setenv("SCENARIO_DBHOST", "localhost")
	os.Setenv("SCENARIO_DBNAME", "MAIN")
	os.Setenv("SCENARIO_DBPWD", "123456")
	os.Setenv("SCENARIO_DBPPORT", "3306")

	dbLocHost = os.Getenv("DBHOST")
	dbLocName = os.Getenv("SCENARIO_DBNAME")
	dbLocPwd = os.Getenv("SCENARIO_DBPWD")
	dbLocPort = os.Getenv("SCENARIO_DBPPORT")

	locDbAdmin = fmt.Sprintf("root:%s@(%s:%s)/?charset=utf8&parseTime=True&net_write_timeout=9000", dbLocPwd, dbLocHost, dbLocPort)
	locDbLocation = fmt.Sprintf("root:%s@(%s:%s)/%s?charset=utf8&parseTime=True&net_write_timeout=9000", dbLocPwd, dbLocHost, dbLocPort, dbLocName)
}
func printVariables() {
	fmt.Println(locDbLocation)
}
