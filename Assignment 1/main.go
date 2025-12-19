package main

import (
	"fmt"

	"github.com/DenisLi/Assignment1/Bank"
	"github.com/DenisLi/Assignment1/Company"
	"github.com/DenisLi/Assignment1/Library"
	"github.com/DenisLi/Assignment1/Shapes"
)

func main() {

	//====================================================================================================
	//library test
	//====================================================================================================

	library := new(Library.Library)
	library1 := new(Library.Library)
	library.CLI()
	library1.CLI()
	library.CLI()
	//====================================================================================================
	//shapes test
	//====================================================================================================
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

	//====================================================================================================
	//company test
	//====================================================================================================

	company := new(Company.Company)

	fullTimeEmployee1 := company.CreateFullEmployee("Denis", "Li", "Someone IDK", 3200000)
	fullTimeEmployee2 := company.CreateFullEmployee("Definetely not Denis", "Not Li", "Director", 3200001)
	partTimeEmployee1 := company.CreatePartEmployee("Aslan", "Asylkhanov", "Office worker", 5000, 15)
	company.AddEmployee(fullTimeEmployee1)
	company.AddEmployee(fullTimeEmployee2)
	company.AddEmployee(partTimeEmployee1)
	company.ListEmployees()

	//====================================================================================================
	//bank test
	//====================================================================================================

	bank := Bank.NewBankAccount("Denis", "Li", "4400 4400 1234 5678")
	bank.CLI()

}
