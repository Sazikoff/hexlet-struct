package structure

import (
    "testing"
	"fmt"
)

func TestExperiment(t *testing.T) {
	shipments := []Shipment{
		{ID: "A1", City: "Berlin", Delivered: true},
		{ID: "B2", City: "Paris", Delivered: false},
		{ID: "C3", City: "Berlin", Delivered: true},
	}

	fmt.Println(FilterByCity(shipments, "Berlin")) // []Shipment{{"A1","Berlin"}, {"C3","Berlin"}}
	fmt.Println(DeliveredIDs(shipments))           // []string{"A1", "C3"}
	fmt.Println(IndexByID(shipments)["B2"])        // Shipment{ID: "B2", City: "Paris", Delivered: false}
}