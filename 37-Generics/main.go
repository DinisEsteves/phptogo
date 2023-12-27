package main

import "fmt"

type collection[T int | float64] []T

func sum[T int | float64](s []T) T {
	var t T

	for i := 0; i < len(s); i++ {
		t += s[i]
	}

	return t
}
func main() {
	c := []float64{5.5, 10.5, 20.5}

	fmt.Println(sum(c))
}
