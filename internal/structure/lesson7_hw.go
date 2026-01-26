package structure

import (
	"fmt"
)

type ReportConfig struct {
	Title           string
	Limit           int
	IncludeArchived bool
}

func NewReportConfig(title string, limit int, includeArchived bool) (ReportConfig, error) {
	if title == "" {
		return ReportConfig{}, fmt.Errorf("non of title!")
	}

	if limit <= 0 {
		limit = 100
	}

	return ReportConfig{
		Title:           title,
		Limit:           limit,
		IncludeArchived: includeArchived,
	}, nil
}

func (cfg ReportConfig) MaxItems() int {
	return cfg.Limit
}
