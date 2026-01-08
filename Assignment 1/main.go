package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/DenisLi/Assignment1/Bank"
	"github.com/DenisLi/Assignment1/Company"
	"github.com/DenisLi/Assignment1/Library"
	"github.com/DenisLi/Assignment1/Shapes"
	"github.com/DenisLi/Assignment1/tickets"
)

func main() {

	//====================================================================================================
	//library test
	//====================================================================================================

	library := new(Library.Library)
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
	CLI()

}

func CLI() {
	id := 1
	id_agents := 1
	store := tickets.NewTicketStore()
	agents := make(map[string]tickets.Agent)
	scanner := bufio.NewScanner(os.Stdin)
	var cursor string
	for cursor != "9" {
		fmt.Println("Welcome to console menu\n 1. create ticket\n 2. Add agent\n 3. Assign ticket to agent\n 4. Resolve ticket\n 5. List all tickets\n 6. List open tickets\n 7. List done tickets\n 8.list unassigned tickets \n9. exit")
		scanner.Scan()
		cursor = scanner.Text()
		switch cursor {
		case "1":
			fmt.Println("Enter ticket name")
			scanner.Scan()
			title := scanner.Text()
			fmt.Println("Enter ticket description")
			scanner.Scan()
			description := scanner.Text()
			fmt.Println("Enter priority")
			scanner.Scan()
			priority, error := strconv.Atoi(scanner.Text())
			if error == nil {
				store.Create(tickets.Ticket{

					ID:          strconv.Itoa(id),
					Title:       title,
					Description: description,
					Priority:    priority,
					AssigneeID:  "Noone",
					Status:      "OPEN",
				})
				id++
			}
		case "2":
			fmt.Println("Enter human or bot")
			scanner.Scan()
			checker := scanner.Text()
			if checker == "Human" {
				fmt.Println("Enter name")
				scanner.Scan()
				name := scanner.Text()
				agents[strconv.Itoa(id_agents)] = tickets.HumanAgent{

					ID:   strconv.Itoa(id_agents),
					Name: name,
				}
				id_agents++
			} else if checker == "Bot" {
				fmt.Println("Enter name")
				scanner.Scan()
				name := scanner.Text()
				fmt.Println("Enter version")
				scanner.Scan()
				ver := scanner.Text()
				agents[strconv.Itoa(id_agents)] = tickets.BotAgent{

					ID:      strconv.Itoa(id_agents),
					Name:    name,
					Version: ver,
				}
				id_agents++
			}
		case "3":
			fmt.Println("Enter agent id")
			scanner.Scan()
			someid := scanner.Text()
			if _, exists := agents[someid]; !exists {
				fmt.Println("No such agent id")
			} else {
				fmt.Println("Enter ticket id")
				scanner.Scan()
				ticketid := scanner.Text()
				store.Assign(ticketid, someid)
			}
		case "4":
			fmt.Println("Enter ticket Id you want to resolve")
			scanner.Scan()
			someid := scanner.Text()
			store.Resolve(someid)

		case "5":
			fmt.Println(store.ListAll())
		case "6":
			fmt.Println(store.ListByStatus("OPEN"))
		case "7":
			fmt.Println(store.ListByStatus("DONE"))
		case "8":
			fmt.Println(store.ListUnassigned())
		case "9":
			return
		}

	}

}
