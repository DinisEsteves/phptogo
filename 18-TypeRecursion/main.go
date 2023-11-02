package main

func main() {
	//Cannot define partner as a person, it would cause an infinite memory count (go mesures always the custom types size)
	//In this case we need to set partner as a pointer to a person struct in memory
	type person struct {
		name    string
		age     int
		partner *person
	}

	dinis := person{}
	partner1 := person{}

	//& is how we get the memory address of a variable
	dinis.partner = &partner1
}
