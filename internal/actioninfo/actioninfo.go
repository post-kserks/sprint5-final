package actioninfo

import (
	"fmt"
	"log"
	"strings"
)

type DataParser interface {
	Parse(datastring string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	for i, data := range dataset {
		err := dp.Parse(data)
		if err != nil {
			log.Printf("Error parsing data at index %d: %v", i, err)
			continue
		}

		info, err := dp.ActionInfo()
		if err != nil {
			log.Printf("Error getting action info at index %d: %v", i, err)
			continue
		}

		if !strings.HasSuffix(info, "\n") {
			info += "\n"
		}
		fmt.Print(info)
	}
}
