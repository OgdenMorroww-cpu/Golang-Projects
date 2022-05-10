package structures

import "fmt"

type Contacts struct {
	first_name string
	last_name  string
	age        int
	sex        string
}

type Coordinates struct {
	x int
	y int
}

type Employee struct {
	name          string
	job_title     string
	annual_salary int
	gpa           float64
	age           int
	position      string
}

func My_structs() {
	contacts := Contacts{"Elon", "Musk", 45, "Male"}
	fmt.Println("First Name:", contacts.first_name)
	fmt.Println("Last Name:", contacts.last_name)
	fmt.Println("Age:", contacts.age)
	fmt.Println("Sex:", contacts.sex)
	fmt.Println()
	contacts_1 := Contacts{"Jack", "Dorsey", 35, "Male"}
	fmt.Println("First Name:", contacts_1.first_name)
	fmt.Println("Last Name:", contacts_1.last_name)
	fmt.Println("Age:", contacts_1.age)
	fmt.Println("Sex:", contacts_1.sex)
	fmt.Println()
	contacts_2 := Contacts{first_name: "Sarah", last_name: "Blakely", age: 38, sex: "Female"}
	fmt.Println("First Name:", contacts_2.first_name)
	fmt.Println("Last Name:", contacts_2.last_name)
	fmt.Println("Age:", contacts_2.age)
	fmt.Println("Sex:", contacts_2.sex)
	fmt.Println()
	a := Coordinates{x: 7, y: 5}
	fmt.Println(a.x - a.y)
	employee := Employee{
		name:          "Jake Paul",
		job_title:     "Boxer",
		annual_salary: 45000,
		age:           28,
		position:      "Chief Boxer",
		gpa:           45.8,
	}
	employee_ptr := &employee
	fmt.Println(employee_ptr.name)
	fmt.Println(employee_ptr.job_title)
	fmt.Println(employee_ptr.annual_salary)
	fmt.Println(employee_ptr.age)
	fmt.Println(employee_ptr.position)
	fmt.Println(employee_ptr.gpa)
	new_employee := &Employee{
		name:          "Benard Anault",
		job_title:     "Businessman",
		annual_salary: 450000000,
		age:           85,
		position:      "CEO",
		gpa:           68.9,
	}
	fmt.Println("Name:", new_employee.name)
	fmt.Println("Job Title:", new_employee.job_title)
	fmt.Println("Annual Salary:", new_employee.annual_salary)
	fmt.Println("Age:", new_employee.age)
	fmt.Println("Position:", new_employee.position)
	fmt.Println("GPA:", new_employee.gpa)
}
