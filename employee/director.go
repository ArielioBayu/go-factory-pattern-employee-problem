package employee

type Director struct {
	Employee
}

func NewDirector() *Director {
	return &Director{
		Employee: Employee{
			Name:   "Director",
			Salary: 5000,
		},
	}
}

func (m *Director) GetBonus() float64 {
	return float64(m.Salary * 30 / 100)
}
