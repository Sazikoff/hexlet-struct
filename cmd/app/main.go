package main

import (
	// "fmt"

	"fmt"

	"github.com/Sazikoff/hexlet-struct/internal/structure"
	// "github.com/Sazikoff/hexlet-struct/internal/employees"
	// "github.com/Sazikoff/hexlet-struct/internal/payroll"
)

func main() {
	c := structure.Candidate{
		Name:            "Irina",
		ExperienceYears: 5,
		Skills:          []string{"Go", "Docker", "Kubernetes"},
		Remote:          true,
		DesiredSalary:   4200,
	}

	fmt.Print(structure.DescribeCandidate(c)) // Irina: 5y exp, salary: 4200, remote: yes
	fmt.Println(structure.HasSkill(c, "Go"))    // true
	structure.HasSkill(c, "Rust")  // false
}
