package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type person_alias struct {
	person
	first string
}

func (p person_alias) fullNameNAlias() string { //any value of type person_alis may call this method
	return p.person.first + " " + p.person.last + " alias " + p.first
}
func main() {
	p1 := person{"daniel", "campillo", 25}
	fmt.Printf("%T\n", p1)
	fmt.Printf("name: %s -- age: %d\n\n", p1.first, p1.age)

	p2 := person_alias{
		person: person{
			first: "daniel",
			last:  "campillo",
			age:   25,
		},
		first: "el impredecible",
	}
	fmt.Println(p2.person.first, p2.first)
	fmt.Println("Calling the method:", p2.fullNameNAlias())
}
