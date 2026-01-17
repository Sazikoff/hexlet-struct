package payroll

import "github.com/Sazikoff/hexlet-struct/internal/employees"

func CalcGross(e employees.Employee, bonusPct float64) float64 {
	return e.BaseSalary() + bonusPct
}

func CalcNet(gross float64, taxPct float64) float64 {
	return gross*(1-taxPct)
}
