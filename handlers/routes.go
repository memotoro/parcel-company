package handlers

import (
	"net/http"

	"github.com/gorilla/mux"

	"zavamed.com/parcel-company/data"
)

// CreateRouter -
func CreateRouter(myApp *data.App) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/trucks", CreateTruck(myApp)).Methods(http.MethodPost)
	router.HandleFunc("/trucks", GetTrucks(myApp)).Methods(http.MethodGet)
	router.HandleFunc("/trucks/{truckID}", GetTruck(myApp)).Methods(http.MethodGet)
	router.HandleFunc("/trucks/{truckID}/state/{timeStamp}", GetTruckState(myApp)).Methods(http.MethodGet)
	router.HandleFunc("/trucks/{truckID}/parcels", LoadParcelToTruck(myApp)).Methods(http.MethodPost)
	router.HandleFunc("/trucks/{truckID}/parcels/{parcelID}", UnloadParcelFromTruck(myApp)).Methods(http.MethodDelete)

	return router
}
