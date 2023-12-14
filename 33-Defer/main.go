package main

import (
	"fmt"
	"strings"
)

func printer(lines []string) (out string) {
	defer func() {
		//Since defer it's executed in the end of this func the variable out is acessible
		if out == "" {
			fmt.Println("<nil>")
		}

		fmt.Println("OK")
	}()

	out = strings.Join(lines, "\n")

	if out != "" {
		out = out + "\n"
	}

	return out
}

func main() {
	message := printer([]string{"Foo", "Bar"})

	fmt.Print(message)
}
