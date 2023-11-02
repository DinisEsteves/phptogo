package main

func main() {
	var person struct {
		name    string
		age     int
		cars    [1]string
		devices map[string]string
	}

	person.name = "Dinis"
	person.age = 25
	person.cars = [1]string{"Audi"}
	person.devices = map[string]string{
		"Mac book pro": "PC",
		"Airpods pro":  "PC",
	}
}
