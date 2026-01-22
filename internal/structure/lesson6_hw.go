package structure

import (
	// "fmt"
)

type Draft struct {
	Title string
	Items []string
	published bool
}

func (d *Draft) AddItem(item string) {
	if item != "" {
		d.Items = append(d.Items, item)
	}
}

func (d *Draft) Publish() {
	d.published = true
}

func (d Draft) IsPublished() bool {
	return d.published
}