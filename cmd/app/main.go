package main

import (
	"fmt"
	"github.com/Sazikoff/hexlet-struct/internal/employees"
	"github.com/Sazikoff/hexlet-struct/internal/payroll"
)

func main() {
	alex := employees.Employee{

		Name:     "Alex",
		Position: "Junior",
		// baseSalary: 100,
		Skills: []string{"hard worker", "smart", "Honest"},
	}

	var test bool
	for {
		test = alex.SetBaseSalary(100)
		if test == true {
			break
		}
	}
	fmt.Println(alex.BaseSalary())
	dirtyCash := payroll.CalcGross(alex, 50)
	fmt.Println(dirtyCash)
	fmt.Println(payroll.CalcNet(dirtyCash, 0.13))

}
