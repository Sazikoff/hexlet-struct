package main

import (
	// "fmt"

	"fmt"

	. "github.com/Sazikoff/hexlet-struct/internal/structure"
	// "github.com/Sazikoff/hexlet-struct/internal/employees"
	// "github.com/Sazikoff/hexlet-struct/internal/payroll"
)

func main() {
	shipments := []Shipment{
		{ID: "A1", City: "Berlin", Delivered: true},
		{ID: "B2", City: "Paris", Delivered: false},
		{ID: "C3", City: "Berlin", Delivered: true},
	}

	fmt.Println(FilterByCity(shipments, "Berlin"))         // []Shipment{{"A1","Berlin"}, {"C3","Berlin"}}
	fmt.Println(DeliveredIDs(shipments))                   // []string{"A1", "C3"}
	fmt.Println(IndexByID(shipments)["B2"])             // Shipment{ID: "B2", City: "Paris", Delivered: false}
	
	
	// fmt.Print(structure.DescribeCandidate(c)) // Irina: 5y exp, salary: 4200, remote: yes
	// fmt.Println(structure.HasSkill(c, "Go"))  // true
	// structure.HasSkill(c, "Rust")             // false
}
