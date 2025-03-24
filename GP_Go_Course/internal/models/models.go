package models

import "fmt"

type Person struct {
	Name     string
	LastName string
	Age      int
}

type TypeElement struct {
	Name string
}

var exampleElement = TypeElement{
	Name: "John Doe",
}

func MyModel(element TypeElement) {

	element.Name = "Jane Doe"

	fmt.Println(exampleElement.Name)

	var p *int
	i := 42
	p = &i

}
