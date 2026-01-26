package structure

import "fmt"

// BEGIN (write your solution here)
type Customer struct {
	Name string
	Phone string
}

type Address struct {
	City string
	Street string
	House string
}

type Delivery struct {
	ID string
	Customer Customer
	Address Address
	Items []string
}

func FormatDelivery(d Delivery) string {
	return fmt.Sprint(d.ID, d.Customer.Name, "(", ") ->", d.Address.City, d.Address.Street, d.Address.House)
}
// END