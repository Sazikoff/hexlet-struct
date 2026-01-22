package sandbox

import (
	"fmt"

	"github.com/Sazikoff/hexlet-struct/internal/structure"
)

func Run7() {

	cfg, _ := structure.NewReportConfig("Weekly sales", 0, true)
	// err == nil, cfg.Limit == 100

	fmt.Println(cfg.MaxItems()) // 100
	
		
	if _, err := structure.NewReportConfig("", 50, false); err != nil {
		fmt.Println(err)
	}
	// err != nil, потому что заголовок пустой

	
}
