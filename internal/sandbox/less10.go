package sandbox

import (
	"fmt"
	"reflect"
)

type Product struct {
	ID    int
	Price int
	Tags  []string
}

// Equal проверяет равенство товаров

func (p Product) Equal(other Product) bool {
	if p.ID != other.ID || p.Price != other.Price {
		return false
	}

	if len(p.Tags) != len(other.Tags) {
		return false
	}

	for i := range p.Tags {
		if p.Tags[i] != other.Tags[i] {
			return false
		}
	}

	return true
}
func Run10() {
	p1 := Product{ID: 1, Price: 100, Tags: []string{"sale", "popular"}}
	p2 := Product{ID: 1, Price: 100, Tags: []string{"sale", "popular"}}

	fmt.Println(p1.Equal(p2)) // true

	fmt.Println(reflect.DeepEqual(p1, p2)) // true
}

type Profile struct {
	ID   int
	Tags []string
	Meta map[string]string
}

func (p Profile) Equal(other Profile) bool {
	if p.ID != other.ID {
		return false
	}

	if len(p.Tags) != len(other.Tags) {
		return false
	}

	for i := range p.Tags {
		if p.Tags[i] != other.Tags[i] {
			return false
		}
	}

	if len(p.Meta) != len(other.Meta) {
		return false
	}

	for k, v := range p.Meta {
		if other.Meta[k] != v {
			return false
		}
	}

	return true
}

func (p Profile) Clone() Profile {

	var tags []string

	if p.Tags != nil {
		tags = make([]string, len(p.Tags))
		copy(tags, p.Tags)
	}

	var meta map[string]string

	if p.Meta != nil {
		meta = make(map[string]string, len(p.Meta))
		for k, v := range p.Meta {
			meta[k] = v
		}
	}

	return Profile{
		ID:   p.ID,
		Tags: tags,
		Meta: meta,
	}

}

func Run11() {
	p := Profile{
		ID:   7,
		Tags: []string{},
		Meta: map[string]string{"city": "Moscow"},
	}

	copy := p.Clone()
	fmt.Println(p.Equal(copy)) // true

	copy.Tags[0] = "std"
	fmt.Println(p.Equal(copy)) // false, оригинал не меняется
}

type Health struct {
	Status    string
	Incidents []string
}

func (h *Health) AddIncident(incident string) {
	if incident != "" {
		h.Incidents = append(h.Incidents, incident)
	}
	if h.Status == "" {
		h.Status = "degraded"
	}
}

func (h Health) IsStable() bool {
	if h.Status == "ok" && len(h.Incidents) == 0 {
		return true
	}

	return false
}

type Service struct {
	Name string
	Health
}

func NewService(name string) Service {
	return Service{
		Name: name,
		Health: Health{
			Status:    "ok",
			Incidents: []string{},
		},
	}
}

func Run12() {
	svc := NewService("billing")
	fmt.Println(svc.IsStable()) // true

	svc.Status = ""
	svc.AddIncident("")

	fmt.Println(svc.Incidents)  // []string{"database latency"}
	fmt.Println(svc.Status)     // "degraded"
	fmt.Println(svc.IsStable()) // false

	// Lol := Service{
	// 	Name: "name",
	// 	Health: Health{
	// 		Status:    "ok",
	// 		Incidents: []string{""},
	// 	},
	// }
	fmt.Println(len(svc.Incidents))
}
