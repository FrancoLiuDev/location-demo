package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type httpHandlerFunc func(w http.ResponseWriter, r *http.Request)

func checkSecurity(next httpHandlerFunc) httpHandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		next(res, req)
	}
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/location", checkSecurity(getLocations)).Methods("GET")
	router.HandleFunc("/location", checkSecurity(putLocation)).Methods("POST")
	log.Fatal(http.ListenAndServe(":3010", router))
}

func getLocations(res http.ResponseWriter, req *http.Request) {
	locs := dbLocSelect()
	locsB, _ := json.Marshal(locs)
	res.Write([]byte(string(locsB)))
}

func putLocation(res http.ResponseWriter, req *http.Request) {
	loc := new(Data)
	err := json.NewDecoder(req.Body).Decode(loc)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	dbLocInsert(*loc)
}
