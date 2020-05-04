package handlers

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"

	"github.com/stretchr/testify/assert"
	"zavamed.com/parcel-company/data"
)

func TestCreateTruck(t *testing.T) {
	app := data.NewApp()
	router := CreateRouter(app)

	payload := `{"model":"m1","capacityKg":100.5}`

	request, err := http.NewRequest(http.MethodPost, "/trucks", strings.NewReader(payload))
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	defer recorder.Result().Body.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(recorder.Result().Body)
	result := buf.String()

	assert.Equal(t, 200, recorder.Result().StatusCode)
	assert.Equal(t, true, strings.Contains(result, `"model":"m1"`))
	assert.Equal(t, true, strings.Contains(result, `"capacityKg":100.5`))
}

func TestCreateInvalidTruck(t *testing.T) {
	app := data.NewApp()
	router := CreateRouter(app)

	payload := ``

	request, err := http.NewRequest(http.MethodPost, "/trucks", strings.NewReader(payload))
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	defer recorder.Result().Body.Close()

	assert.Equal(t, 400, recorder.Result().StatusCode)
}

func TestGetTrucks(t *testing.T) {
	app := data.NewApp()
	router := CreateRouter(app)

	createTrucks(t, router)

	request, err := http.NewRequest(http.MethodGet, "/trucks", nil)
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	buf := new(bytes.Buffer)
	buf.ReadFrom(recorder.Result().Body)
	result := buf.String()

	assert.Equal(t, 200, recorder.Result().StatusCode)
	assert.Equal(t, true, strings.Contains(result, `"model":"m1"`))
	assert.Equal(t, true, strings.Contains(result, `"capacityKg":10`))
}

func TestGetTruckByID(t *testing.T) {
	app := data.NewApp()
	router := CreateRouter(app)

	createTrucks(t, router)

	ID := 3
	request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/trucks/%d", ID), nil)
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	buf := new(bytes.Buffer)
	buf.ReadFrom(recorder.Result().Body)
	result := buf.String()

	assert.Equal(t, 200, recorder.Result().StatusCode)
	assert.Equal(t, true, strings.Contains(result, `"model":"m3"`))
	assert.Equal(t, true, strings.Contains(result, `"capacityKg":30`))
}

func TestGetTruckByNotFoundID(t *testing.T) {
	app := data.NewApp()
	router := CreateRouter(app)

	createTrucks(t, router)

	ID := -1
	request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/trucks/%d", ID), nil)
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	assert.Equal(t, 404, recorder.Result().StatusCode)
}

func TestGetTruckByInvalidID(t *testing.T) {
	app := data.NewApp()
	router := CreateRouter(app)

	createTrucks(t, router)

	wrongID := "A"
	request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/trucks/%s", wrongID), nil)
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	assert.Equal(t, 400, recorder.Result().StatusCode)
}

func TestLoadParcel(t *testing.T) {
	app := data.NewApp()
	router := CreateRouter(app)

	createTrucks(t, router)

	truckID := 2
	payload := `{"weightKg":5}`
	request, err := http.NewRequest(http.MethodPost, fmt.Sprintf("/trucks/%d/parcels", truckID), strings.NewReader(payload))
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	buf := new(bytes.Buffer)
	buf.ReadFrom(recorder.Result().Body)
	result := buf.String()

	assert.Equal(t, 200, recorder.Result().StatusCode)
	assert.Equal(t, true, strings.Contains(result, `"weightKg":5`))
}

func TestLoadInvalidParcel(t *testing.T) {
	app := data.NewApp()
	router := CreateRouter(app)

	createTrucks(t, router)

	truckID := 2
	payload := ``
	request, err := http.NewRequest(http.MethodPost, fmt.Sprintf("/trucks/%d/parcels", truckID), strings.NewReader(payload))
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	assert.Equal(t, 400, recorder.Result().StatusCode)
}

func TestLoadParcelWrongTruckID(t *testing.T) {
	app := data.NewApp()
	router := CreateRouter(app)

	createTrucks(t, router)

	wrongID := "A"
	payload := ``
	request, err := http.NewRequest(http.MethodPost, fmt.Sprintf("/trucks/%s/parcels", wrongID), strings.NewReader(payload))
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	assert.Equal(t, 400, recorder.Result().StatusCode)
}

func TestLoadParcelInvalidTruck(t *testing.T) {
	app := data.NewApp()
	router := CreateRouter(app)

	createTrucks(t, router)

	truckID := -2
	payload := `{"weightKg":5}`
	request, err := http.NewRequest(http.MethodPost, fmt.Sprintf("/trucks/%d/parcels", truckID), strings.NewReader(payload))
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	assert.Equal(t, 404, recorder.Result().StatusCode)
}

func TestLoadParcelBiggerCapacity(t *testing.T) {
	app := data.NewApp()
	router := CreateRouter(app)

	createTrucks(t, router)

	truckID := 2
	payload := `{"weightKg":500}`
	request, err := http.NewRequest(http.MethodPost, fmt.Sprintf("/trucks/%d/parcels", truckID), strings.NewReader(payload))
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	assert.Equal(t, 500, recorder.Result().StatusCode)
}

func TestUnloadParcel(t *testing.T) {
	app := data.NewApp()
	router := CreateRouter(app)

	createTrucks(t, router)

	truckID := 2
	payload := `{"weightKg":5}`
	request, err := http.NewRequest(http.MethodPost, fmt.Sprintf("/trucks/%d/parcels", truckID), strings.NewReader(payload))
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	parcelID := 1
	request, err = http.NewRequest(http.MethodDelete, fmt.Sprintf("/trucks/%d/parcels/%d", truckID, parcelID), nil)
	if err != nil {
		t.Fatal(err)
	}
	recorder = httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	buf := new(bytes.Buffer)
	buf.ReadFrom(recorder.Result().Body)
	result := buf.String()

	assert.Equal(t, 200, recorder.Result().StatusCode)
	assert.Equal(t, true, strings.Contains(result, `"model":"m2"`))
	assert.Equal(t, true, strings.Contains(result, `"totalWeight":0`))
	assert.Equal(t, true, strings.Contains(result, `"totalParcels":0`))
}

func TestUnloadParcelInvalidTruckID(t *testing.T) {
	app := data.NewApp()
	router := CreateRouter(app)

	createTrucks(t, router)

	truckID := 1
	payload := `{"weightKg":5}`
	request, err := http.NewRequest(http.MethodPost, fmt.Sprintf("/trucks/%d/parcels", truckID), strings.NewReader(payload))
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	wrongTruckID := "A"
	parcelID := 1
	request, err = http.NewRequest(http.MethodDelete, fmt.Sprintf("/trucks/%s/parcels/%d", wrongTruckID, parcelID), nil)
	if err != nil {
		t.Fatal(err)
	}
	recorder = httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	assert.Equal(t, 400, recorder.Result().StatusCode)
}

func TestUnloadParcelInvalidParcelID(t *testing.T) {
	app := data.NewApp()
	router := CreateRouter(app)

	createTrucks(t, router)

	truckID := 1
	payload := `{"weightKg":5}`
	request, err := http.NewRequest(http.MethodPost, fmt.Sprintf("/trucks/%d/parcels", truckID), strings.NewReader(payload))
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	parcelID := "A"
	request, err = http.NewRequest(http.MethodDelete, fmt.Sprintf("/trucks/%d/parcels/%s", truckID, parcelID), nil)
	if err != nil {
		t.Fatal(err)
	}
	recorder = httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	assert.Equal(t, 400, recorder.Result().StatusCode)
}

func TestUnloadParcelTruckIDNotFound(t *testing.T) {
	app := data.NewApp()
	router := CreateRouter(app)

	createTrucks(t, router)

	truckID := 1
	payload := `{"weightKg":5}`
	request, err := http.NewRequest(http.MethodPost, fmt.Sprintf("/trucks/%d/parcels", truckID), strings.NewReader(payload))
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	truckID = -1
	parcelID := 1
	request, err = http.NewRequest(http.MethodDelete, fmt.Sprintf("/trucks/%d/parcels/%d", truckID, parcelID), nil)
	if err != nil {
		t.Fatal(err)
	}
	recorder = httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	assert.Equal(t, 404, recorder.Result().StatusCode)
}

func TestUnloadParcelIDNotFound(t *testing.T) {
	app := data.NewApp()
	router := CreateRouter(app)

	createTrucks(t, router)

	truckID := 1
	payload := `{"weightKg":5}`
	request, err := http.NewRequest(http.MethodPost, fmt.Sprintf("/trucks/%d/parcels", truckID), strings.NewReader(payload))
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	truckID = 1
	parcelID := -1
	request, err = http.NewRequest(http.MethodDelete, fmt.Sprintf("/trucks/%d/parcels/%d", truckID, parcelID), nil)
	if err != nil {
		t.Fatal(err)
	}
	recorder = httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	assert.Equal(t, 404, recorder.Result().StatusCode)
}

// Private methods

func createApp() *data.App {
	return &data.App{}
}

func createTrucks(t *testing.T, router *mux.Router) {
	i := 1
	for i < 10 {
		payload := fmt.Sprintf("{\"model\":\"m%d\",\"capacityKg\":%d}", i, i*10)

		request, err := http.NewRequest(http.MethodPost, "/trucks", strings.NewReader(payload))
		if err != nil {
			t.Fatal(err)
		}
		recorder := httptest.NewRecorder()

		router.ServeHTTP(recorder, request)

		i++
	}
}
