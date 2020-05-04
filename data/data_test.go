package data

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApp(t *testing.T) {
	app := NewApp()

	assert.Equal(t, int64(0), app.ParcelID)
	assert.Equal(t, int64(0), app.TruckID)
	assert.Equal(t, 0, len(app.Trucks))
}

func TestCreateParcel(t *testing.T) {
	parcelID := int64(1)
	weight := float64(10.5)
	parcel := Parcel{parcelID, weight}

	assert.NotNil(t, parcel)
	assert.Equal(t, parcelID, parcel.ID)
	assert.Equal(t, weight, parcel.WeightKG)
}

func TestCreateTruck(t *testing.T) {
	ID := int64(1)
	capacityKg := float64(50.50)
	truck := createTruck(ID, capacityKg)

	assert.NotNil(t, truck)
	assert.Equal(t, ID, truck.ID)
}

func TestAddParcelToTruck(t *testing.T) {
	ID := int64(1)
	capacityKg := float64(50.50)
	truck := createTruck(ID, capacityKg)

	parcelID1 := int64(1)
	weight1 := float64(20.5)
	parcel1 := createParcel(parcelID1, weight1)

	truck.AddParcel(*parcel1)

	assert.NotNil(t, truck.Parcels)
	assert.Equal(t, 1, len(truck.Parcels))
	assert.Equal(t, parcelID1, truck.Parcels[0].ID)
	assert.Equal(t, weight1, truck.Parcels[0].WeightKG)

	parcelID2 := int64(2)
	weight2 := float64(10.5)
	parcel2 := createParcel(parcelID2, weight2)

	truck.AddParcel(*parcel2)

	assert.NotNil(t, truck.Parcels)
	assert.Equal(t, 2, len(truck.Parcels))
	assert.Equal(t, int64(2), truck.GetTotalParcels())
	assert.Equal(t, parcelID2, truck.Parcels[1].ID)
	assert.Equal(t, weight2, truck.Parcels[1].WeightKG)
}

func TestAddParcelToTruckOverload(t *testing.T) {
	ID := int64(1)
	capacityKg := float64(10.50)
	truck := createTruck(ID, capacityKg)

	parcelID1 := int64(1)
	weight1 := float64(10.5)
	parcel1 := createParcel(parcelID1, weight1)

	truck.AddParcel(*parcel1)

	parcelID2 := int64(2)
	weight2 := float64(5)
	parcel2 := createParcel(parcelID2, weight2)

	result := truck.AddParcel(*parcel2)

	assert.NotNil(t, truck.Parcels)
	assert.Equal(t, 1, len(truck.Parcels))
	assert.Equal(t, false, result)
}

func TestRemoveParcelFromTruck(t *testing.T) {
	ID := int64(1)
	capacityKg := float64(100.50)
	truck := createTruck(ID, capacityKg)

	parcelID1 := int64(1)
	weight1 := float64(10.5)
	parcel1 := createParcel(parcelID1, weight1)

	truck.AddParcel(*parcel1)

	parcelID2 := int64(2)
	weight2 := float64(20.5)
	parcel2 := createParcel(parcelID2, weight2)

	truck.AddParcel(*parcel2)

	parcelID3 := int64(3)
	weight3 := float64(30.5)
	parcel3 := createParcel(parcelID3, weight3)

	truck.AddParcel(*parcel3)

	assert.NotNil(t, truck.Parcels)
	assert.Equal(t, 3, len(truck.Parcels))

	result := truck.RemoveParcel(parcelID2)

	assert.Equal(t, true, result)
	assert.Equal(t, 2, len(truck.Parcels))
}

func TestRemoveNonLoadedParcelFromTruck(t *testing.T) {
	ID := int64(1)
	capacityKg := float64(50.50)
	truck := createTruck(ID, capacityKg)

	parcelID1 := int64(1)
	weight1 := float64(10.5)
	parcel1 := createParcel(parcelID1, weight1)

	truck.AddParcel(*parcel1)

	parcelID2 := int64(10)

	result := truck.RemoveParcel(parcelID2)

	assert.Equal(t, false, result)
	assert.Equal(t, 1, len(truck.Parcels))
}

func TestGetTruckTotalWeigth(t *testing.T) {
	ID := int64(1)
	capacityKg := float64(50.50)
	truck := createTruck(ID, capacityKg)

	parcelID1 := int64(1)
	weight1 := float64(10.5)
	parcel1 := createParcel(parcelID1, weight1)

	truck.AddParcel(*parcel1)

	parcelID2 := int64(2)
	weight2 := float64(20.5)
	parcel2 := createParcel(parcelID2, weight2)

	truck.AddParcel(*parcel2)

	totalWeight := truck.GetTotalWeight()

	assert.NotNil(t, totalWeight)
	assert.Equal(t, weight1+weight2, totalWeight)
}

func TestGetTruckJSON(t *testing.T) {
	ID := int64(1)
	capacityKg := float64(50.50)
	truck := createTruck(ID, capacityKg)

	parcelID1 := int64(1)
	weight1 := float64(20.5)
	parcel1 := createParcel(parcelID1, weight1)

	truck.AddParcel(*parcel1)

	jsonData, err := truck.MarshalJSON()
	jsonExpected := `{"id":1,"model":"Model  1","capacityKg":50.5,"parcels":[{"id":1,"weightKg":20.5}],"totalWeight":20.5,"totalParcels":1}`

	assert.Nil(t, err)
	assert.NotNil(t, truck.Parcels)
	assert.Equal(t, jsonExpected, string(jsonData))
}

func createTruck(ID int64, capacityKg float64) *Truck {
	return &Truck{ID: ID, Model: fmt.Sprintf("%v %v", "Model ", ID), CapacityKG: capacityKg}
}

func createParcel(ID int64, weightKg float64) *Parcel {
	return &Parcel{ID: ID, WeightKG: weightKg}
}
