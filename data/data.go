package data

import "encoding/json"

// Parcel -
type Parcel struct {
	ID       int64   `json:"id"`
	WeightKG float64 `json:"weightKg"`
}

// Truck -
type Truck struct {
	ID           int64    `json:"id"`
	Model        string   `json:"model"`
	CapacityKG   float64  `json:"capacityKg"`
	Parcels      []Parcel `json:"parcels,omitempty"`
	TotalWeight  float64  `json:"totalWeight"`
	TotalParcels int64    `json:"totalParcels"`
}

// App -
type App struct {
	TruckID  int64
	ParcelID int64
	Trucks   map[int64]*Truck
}

// NewApp -
func NewApp() *App {
	trucks := make(map[int64]*Truck)
	return &App{TruckID: 0, ParcelID: 0, Trucks: trucks}
}

// GetTotalParcels -
func (t *Truck) GetTotalParcels() int64 {
	return int64(len(t.Parcels))
}

// GetTotalWeight -
func (t *Truck) GetTotalWeight() float64 {
	sum := float64(0)
	for _, p := range t.Parcels {
		sum += p.WeightKG
	}
	return sum
}

// AddParcel -
func (t *Truck) AddParcel(parcel Parcel) bool {
	if parcel.WeightKG+t.GetTotalWeight() > t.CapacityKG {
		return false
	}

	t.Parcels = append(t.Parcels, parcel)
	return true
}

// RemoveParcel -
func (t *Truck) RemoveParcel(parcelID int64) bool {
	parcelIndex := -1
	for index := range t.Parcels {
		parcel := t.Parcels[index]
		if parcel.ID == parcelID {
			parcelIndex = index
			break
		}
	}

	if parcelIndex == -1 {
		return false
	}

	t.Parcels = append(t.Parcels[:parcelIndex], t.Parcels[parcelIndex+1:]...)

	return true
}

// MarshalJSON -
func (t *Truck) MarshalJSON() ([]byte, error) {
	type truck Truck
	t.TotalWeight = t.GetTotalWeight()
	t.TotalParcels = t.GetTotalParcels()
	newTruck := truck(*t)
	return json.Marshal(newTruck)
}
