package main

import "fmt"

type money uint
type moneyAmout []money

func (m moneyAmout) sum() money {
	var total money

	for i := 0; i < len(m); i++ {
		total += m[i]
	}

	return total
}
func main() {
	amount := moneyAmout{600, 400, 500}

	fmt.Println(amount.sum())
}
