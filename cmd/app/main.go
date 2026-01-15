package main

import (
	"github.com/Sazikoff/hexlet-struct/internal/structure"
	"fmt"
)

func main() {
	var parcel1 = structure.Parcel{ID: "BR-7", Destination: "Berlin", Weight: 12, Insured: true}
	fmt.Println(structure.Describe(parcel1))
}
