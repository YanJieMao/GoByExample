package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	fmt.Println(person{"li", 20})
	fmt.Println(person{name: "liu", age: 30})
	fmt.Println(person{name: "wang"})
	fmt.Println(&person{name: "zhou", age: 40})

	s := person{name: "zhao", age: 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)

}

/*
{li 20}
{liu 30}
{wang 0}
&{zhou 40}
zhao
50
51


*/
