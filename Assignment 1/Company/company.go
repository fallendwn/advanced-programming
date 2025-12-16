package Company

import "fmt"

type Employee interface {
	GetDetails()
}

type FullTimeEmployee struct {
	ID          uint64
	Name        string
	Surname     string
	MonthSalary float64
	Position    string
}

func (fullTimeEmployee *FullTimeEmployee) GetDetails() {

	fmt.Printf("Fulltime employee named : %s %s, Position %s, Salary %.2f\n", fullTimeEmployee.Name, fullTimeEmployee.Surname, fullTimeEmployee.Position, fullTimeEmployee.MonthSalary)

}

type PartTimeEmployee struct {
	ID           uint64
	Name         string
	Surname      string
	MoneyPerHour float64
	Shifts       int
	Position     string
}

func (partTimeEmployee *PartTimeEmployee) GetDetails() {

	fmt.Printf("Parttime employee named : %s %s, Position %s, Salary for shift %.2f, Shifts %d\n", partTimeEmployee.Name, partTimeEmployee.Surname, partTimeEmployee.Position, partTimeEmployee.MoneyPerHour, partTimeEmployee.Shifts)

}

type Company struct {
	Employees map[uint64]Employee
}

//AddEmployee, ListEmployees
func (company *Company) AddEmployee(employee Employee) {
	var id uint64
	switch e := employee.(type) {
	case *FullTimeEmployee:
		id = e.ID
	case *PartTimeEmployee:
		id = e.ID
	}

	if company.Employees == nil {

		company.Employees = make(map[uint64]Employee)

	}

	if _, exists := company.Employees[id]; exists {
		fmt.Println("Employee with such ID already exists")
		return
	}

	company.Employees[id] = employee
	fmt.Println("Employee has been successfully added")

}

func (company *Company) ListEmployees() {

	for _, employee := range company.Employees {

		employee.GetDetails()

	}

}
