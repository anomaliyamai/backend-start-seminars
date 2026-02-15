package main

import "fmt"

type Person struct {
	Name string
	Age  int
	City string
}

func (p Person) Greet() string {
	return fmt.Sprintf("Привет, меня зовут %s, мне %d лет", p.Name, p.Age)
}

func (p *Person) HaveBirthday() {
	p.Age++
}

type Employee struct {
	Person
	Position string
	Salary   float64
}

func (e Employee) String() string {
	return fmt.Sprintf("%s работает как %s с зарплатой %.2f", e.Name, e.Position, e.Salary)
}

func main() {
	person := Person{
		Name: "Анна",
		Age:  25,
		City: "Москва",
	}

	fmt.Println(person.Greet())
	fmt.Printf("Возраст до дня рождения: %d\n", person.Age)

	person.HaveBirthday()
	fmt.Printf("Возраст после дня рождения: %d\n", person.Age)

	employee := Employee{
		Person: Person{
			Name: "Петр",
			Age:  30,
			City: "СПБ",
		},
		Position: "Разработчик",
		Salary:   150000.0,
	}

	fmt.Println(employee.String())
	fmt.Println(employee.Greet())
}
