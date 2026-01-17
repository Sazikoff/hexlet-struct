package structure

import "fmt"

// BEGIN (write your solution here)
type Candidate struct{
    Name string
    ExperienceYears int
    Skills []string
    Remote bool
    DesiredSalary int
}

func HasSkill(c Candidate, skill string) bool {
	for _,v := range c.Skills {
		if v == skill {
			return true
		}
	}
	return false
}

func DescribeCandidate(c Candidate) string {
	remote := "no"
	if c.Remote == true {remote = "yes"}
	return fmt.Sprintln(c.Name, ": ", c.ExperienceYears, " y exp, salary: ", c.DesiredSalary, ", remote: ", remote)
}
// END