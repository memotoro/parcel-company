package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

func main() {
	url := "http://localhost:8080"
	// Trucks
	i := 1
	for i < 10 {
		payload := fmt.Sprintf("{\"model\":\"m%d\",\"capacityKg\":%d}", i, i*10)

		req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/%s", url, "trucks"), strings.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		client := http.Client{}

		resp, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Truck %d response %d", i, resp.StatusCode)
		i++
	}

	rand.Seed(time.Now().UnixNano())

	// Parcels
	j := 1
	min := 1
	max := 10
	for j < 20 {
		w := float64(j) * float64(0.5)
		truckID := rand.Intn(max-min) + min
		payload := fmt.Sprintf("{\"weightKg\":%f}", w)

		req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/%s/%d/%s", url, "trucks", truckID, "parcels"), strings.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		client := http.Client{}

		resp, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Parcel %d response %d", j, resp.StatusCode)
		j++
	}
}
