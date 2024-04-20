package employee

// Staff is a struct for staff
// It embeds Employee
// This means that Staff has all the fields and methods of Employee
type Staff struct {
	Employee
}

// NewStaff creates a new staff
// It returns a pointer to the staff
// Creational method
func NewStaff() *Staff {
	return &Staff{
		Employee: Employee{
			Name:   "Staff",
			Salary: 500,
		},
	}
}

func (m *Staff) GetBonus() float64 {
	return float64(m.Salary * 10 / 100)
}
