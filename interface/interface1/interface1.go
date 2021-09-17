package interface1

import "fmt"

type Person struct {
	Name string
	Sexual string
	Age int
}

func (p *Person) Print() {
	fmt.Printf("Name: %s, Sexual: %s, Age: %d\n", p.Name, p.Sexual, p.Age)
}

func PrintPerson(p *Person) {
	fmt.Printf("Name: %s, Sexual: %s, Age: %d\n", p.Name, p.Sexual, p.Age)
}

func Interface1() {
	var p = Person{
		Name: "richard",
		Sexual: "Male",
		Age: 33,
	}

	PrintPerson(&p)
	p.Print()
}