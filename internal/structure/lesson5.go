package structure

type Shipment struct {
	ID        string
	City      string
	Items     []string
	Delivered bool
}

// BEGIN (write your solution here)
func FilterByCity(shipments []Shipment, city string) []Shipment {
	var Slice []Shipment
	for _,j := range shipments {
		if j.City == city {
			Slice = append(Slice, j)
		}
	}

	return Slice
}

func DeliveredIDs(shipments []Shipment) []string {
	var Slice []string
	for _,j := range shipments {
		if j.Delivered == true {
			Slice = append(Slice, j.ID)
		}
	}

	return Slice
}

func IndexByID(shipments []Shipment) map[string]Shipment {
	M := make(map[string]Shipment)
	for _,j := range shipments{
		M[j.ID] = j
	}

	return M
}
// END