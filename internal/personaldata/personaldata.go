package personaldata

import "fmt"

type Personal struct {
	Name   string
	Weight float64
	Height float64
}

func (p Personal) Print() {
	fmt.Printf("Имя: %s\nВес: %.02f кг.\nРост: %.02f м.\n", p.Name, p.Weight, p.Height)
}
