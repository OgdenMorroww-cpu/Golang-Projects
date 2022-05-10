package structures

import "fmt"

type Billionaires struct {
	name             string
	net_worth        int
	company          string
	company_industry string
	country          string
	age              int
	is_married       bool
}

type Test struct {
	a int
	b int
}

func Example_structs() {
	billionaire_1 := &Billionaires{
		name:             "Bill Gate",
		net_worth:        123000000000000,
		company:          "Microsoft",
		company_industry: "Software",
		country:          "USA",
		age:              57,
		is_married:       false,
	}
	fmt.Println("Industry:", billionaire_1.company_industry)
	billionaire_1.details(2022)
	billionaire_1.check_full_details(1983)
	test := Test{8, 3}
	fmt.Println(test.do())
	billionaire_1.increase_age(10)
	fmt.Println(billionaire_1.age)

}

func (billionaire Billionaires) details(current_year int) int {
	fmt.Println("My name is:", billionaire.name)
	fmt.Println("My networth as at", current_year, "is:", "$:", billionaire.net_worth)
	return current_year
}

func (marital_status Billionaires) check_full_details(year_founded int) int {
	fmt.Println(marital_status.name, "Marital Status is:", marital_status.is_married)
	fmt.Println(marital_status.name, "Is from", marital_status.country)
	fmt.Println(marital_status.name, "Started:", marital_status.company, "in 4th July:", year_founded)
	return year_founded
}

func (x Test) do() int {
	return x.a - x.b
}

func (ptr *Billionaires) increase_age(new_age int) {
	ptr.age += new_age
}
