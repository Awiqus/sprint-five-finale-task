package actioninfo

import (
	"log"
)

type DataParser interface {
	Parse(string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	for _, v := range dataset {
		err := dp.Parse(v)
		if err != nil {
			log.Println("actioninfo: Parese() err != nil")
		}
		dp.ActionInfo()
		if err != nil {
			log.Println("actioninfo: ActionInfo() err != nil")
		}
	}

}
