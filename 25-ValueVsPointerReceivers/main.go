package main

import "fmt"

type money uint
type moneyAmout []money

func (m *moneyAmout) add(amount money) {
	*m = append(*m, amount)
}

func (m *moneyAmout) sum() money {
	var total money

	for i := 0; i < len(*m); i++ {
		total += (*m)[i]
	}

	return total
}
func main() {
	amount := moneyAmout{1, 2, 3}

	amount.add(6)

	fmt.Println(amount.sum())
}
