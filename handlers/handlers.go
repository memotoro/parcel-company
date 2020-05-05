package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"

	"zavamed.com/parcel-company/data"
)

// CreateTruck -
func CreateTruck(myApp *data.App) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var truck data.Truck

		if err := json.NewDecoder(r.Body).Decode(&truck); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		truckID := myApp.TruckID + 1
		myApp.TruckID = truckID

		truck.ID = truckID
		myApp.Trucks[truckID] = &truck

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(truck)
	}
}

// GetTrucks -
func GetTrucks(myApp *data.App) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		trucks := make([]*data.Truck, 0)

		for _, t := range myApp.Trucks {
			trucks = append(trucks, t)
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(trucks)
	}
}

// GetTruck -
func GetTruck(myApp *data.App) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		truckIDValue := vars["truckID"]

		truckID, err := strconv.ParseInt(truckIDValue, 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		truck := myApp.Trucks[truckID]
		if truck == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(truck)
	}
}

// LoadParcelToTruck -
func LoadParcelToTruck(myApp *data.App) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		truckIDValue := vars["truckID"]

		truckID, err := strconv.ParseInt(truckIDValue, 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		var parcel data.Parcel

		if err := json.NewDecoder(r.Body).Decode(&parcel); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		truck := myApp.Trucks[truckID]
		if truck == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		parcelID := myApp.ParcelID + 1
		myApp.ParcelID = parcelID

		parcel.ID = parcelID

		now := time.Now()

		result := truck.AddParcel(parcel, now)
		if !result {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(truck)
	}
}

// UnloadParcelFromTruck -
func UnloadParcelFromTruck(myApp *data.App) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		truckIDValue := vars["truckID"]
		parcelIDValue := vars["parcelID"]

		truckID, err := strconv.ParseInt(truckIDValue, 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		parcelID, err := strconv.ParseInt(parcelIDValue, 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		truck := myApp.Trucks[truckID]
		if truck == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		now := time.Now()

		result := truck.RemoveParcel(parcelID, now)
		if !result {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(truck)
	}
}

// GetTruckState -
func GetTruckState(myApp *data.App) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		truckIDValue := vars["truckID"]
		timeStampValue := vars["timeStamp"]

		truckID, err := strconv.ParseInt(truckIDValue, 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		timeStamp, err := time.Parse(time.RFC3339, timeStampValue)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		truck := myApp.Trucks[truckID]
		if truck == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		/*
			state := truck.History[timeStampValue]
			if state == nil {
				w.WriteHeader(http.StatusNotFound)
				return
			}
		*/

		var selectedStated *data.TruckState
		for index := range truck.History {
			history := truck.History[index]
			timeStampHistory, err := time.Parse(time.RFC3339, history.TimeStamp)
			if err != nil {
				log.Printf("err %v", err)
			}
			if timeStampHistory.After(timeStamp) {
				selectedStated = truck.History[index-1]
				break
			}
		}

		// 2020-01-01 9am
		// 2020-01-01 11am
		// 2020-01-01 1pm

		// 2020-01-01 10pm

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(selectedStated)
	}
}
