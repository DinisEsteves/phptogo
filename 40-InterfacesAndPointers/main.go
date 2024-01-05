package main

import "fmt"

type developer interface {
	message() string
}

type phpDeveloper struct {
}

type goDeveloper struct {
}

func (d phpDeveloper) message() string {
	return "PHP is not dead"
}

func (d *goDeveloper) message() string {
	return "I love gopher"
}

func main() {
	d := []developer{phpDeveloper{}, &goDeveloper{}}

	for _, p := range d {
		fmt.Println(p.message())
	}

}
