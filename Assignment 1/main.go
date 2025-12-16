package main

import (
	"fmt"

	"github.com/DenisLi/Assignment1/Company"
	"github.com/DenisLi/Assignment1/Library"
	"github.com/DenisLi/Assignment1/Shapes"
)

func main() {
	//library test
	library := new(Library.Library)
	book := Library.Book{

		ID:         "3",
		Title:      "no",
		Author:     "no",
		IsBorrowed: false,
	}
	library.AddBook(book)

	//shapes test

	circle := &Shapes.Circle{

		Radius: 3,
	}

	rectangle := &Shapes.Rectangle{

		Width:  3,
		Height: 3,
	}

	square := &Shapes.Square{

		Width: 3,
	}

	triangle := &Shapes.Triangle{

		FirstSide:  3,
		SecondSide: 3,
		ThirdSide:  3,
	}

	slice := []Shapes.Shape{circle, rectangle, square, triangle}

	for _, shape := range slice {
		fmt.Printf("Area of %T is %.2f and perimeter is %.2f\n", shape, shape.CalculateArea(), shape.CalculatePerimeter())

	}

	company := new(Company.Company)

	fullTimeEmployee1 := Company.FullTimeEmployee{

		ID:          1,
		Name:        "Denis",
		Surname:     "Li",
		MonthSalary: 320000,
		Position:    "Someone IDK",
	}
	fullTimeEmployee2 := Company.FullTimeEmployee{

		ID:          1,
		Name:        "Definetely not Denis",
		Surname:     "Not Li",
		MonthSalary: 320001,
		Position:    "Director maybe",
	}
	partTimeEmployee1 := Company.PartTimeEmployee{

		ID:           2,
		Name:         "Aslan",
		Surname:      "Asylkhanov",
		MoneyPerHour: 5000,
		Shifts:       15,
		Position:     "Office worker",
	}
	company.AddEmployee(&fullTimeEmployee1)
	company.AddEmployee(&fullTimeEmployee2)
	company.AddEmployee(&partTimeEmployee1)
	company.ListEmployees()
}
