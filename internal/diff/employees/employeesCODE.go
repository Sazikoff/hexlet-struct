package employees

type Employee struct {
	Name       string
	Position   string
	baseSalary float64
	Skills     []string
}

func (e *Employee) SetBaseSalary(amount float64) bool {
	if amount > 0 {
		e.baseSalary = amount
		return true
	}
	return false
}

func (e *Employee) BaseSalary() float64 {
	return e.baseSalary
}
