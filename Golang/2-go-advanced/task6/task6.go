package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	EmployeeID string
}

func (e Employee) info() {
	fmt.Printf("员工信息:\n")
	fmt.Printf("姓名: %s\n", e.Person.Name)
	fmt.Printf("姓名: %s\n", e.Name)
	fmt.Printf("年龄: %d\n", e.Person.Age)
	fmt.Printf("年龄: %d\n", e.Age)
	fmt.Printf("工号: %s\n", e.EmployeeID)
	fmt.Printf("所有信息: %v\n", e)
}

func main() {
	e := Employee{Person: Person{Name: "John", Age: 20}, EmployeeID: "20250609"}
	e.info()
}
