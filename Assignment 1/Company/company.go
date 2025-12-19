package Company

import "fmt"

type Employee interface {
	GetDetails()
	GetID() uint64
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

func (fullTimeEmployee *FullTimeEmployee) GetID() uint64 {

	return fullTimeEmployee.ID

}

func (c *Company) CreateFullEmployee(name, surname, position string, monthSalary float64) *FullTimeEmployee {
	c.numberOfEmployees++
	return &FullTimeEmployee{
		ID:          c.numberOfEmployees,
		Name:        name,
		Surname:     surname,
		MonthSalary: monthSalary,
		Position:    position,
	}

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

func (partTimeEmployee *PartTimeEmployee) GetID() uint64 {

	return partTimeEmployee.ID

}
func (c *Company) CreatePartEmployee(name, surname, position string, MoneyPerHour float64, shifts int) *PartTimeEmployee {
	c.numberOfEmployees++
	return &PartTimeEmployee{
		ID:           c.numberOfEmployees,
		Name:         name,
		Surname:      surname,
		MoneyPerHour: MoneyPerHour,
		Shifts:       shifts,
		Position:     position,
	}

}

type Company struct {
	Employees         map[uint64]Employee
	numberOfEmployees uint64
}

//AddEmployee, ListEmployees
func (company *Company) AddEmployee(employee Employee) {
	id := employee.GetID()
	fmt.Println(id)
	if employee == nil {
		fmt.Printf("You send nil user.")
		return
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
