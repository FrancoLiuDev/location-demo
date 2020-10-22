package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/lithammer/shortuuid"
	"log"
	"net/http"
)

type httpHandlerFunc func(w http.ResponseWriter, r *http.Request)

func checkSecurity(next httpHandlerFunc) httpHandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		// header := req.Header.Get("Super-Duper-Safe-Security")
		next(res, req)
	}
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/location/", checkSecurity(getLocations)).Methods("Get")
	router.HandleFunc("/location/", checkSecurity(putLocation)).Methods("Put")
	log.Fatal(http.ListenAndServe(":3010", router))
}

func getLocations(res http.ResponseWriter, req *http.Request) {
	locs := dbLocSelect()
	locsB, _ := json.Marshal(locs)
	fmt.Println(string(locsB))
	res.Write([]byte(string(locsB)))
}

func putLocation(res http.ResponseWriter, req *http.Request) {

	loc := new(Data)
	err := json.NewDecoder(req.Body).Decode(loc)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	loc.ID = shortuuid.New()
	fmt.Println((loc))
	dbLocInsert(*loc)
}
