package main

import (
	"fmt"
	"time"
)

type person struct {
	name   string
	age    int
	bornOn int
}

func main() {
	me := person{name: "Andre", age: 23, bornOn: 1}
	monthsOfAge(me)
	personToCalc := newPersona()
	monthsOfAge(personToCalc)
}
func months(x int, y int) int {
	actualMonthI := actualMonth()
	if y == actualMonthI {
		return 0
	}
	if y < actualMonthI {
		return actualMonthI - y + (12 * x)
	} else {
		return (12 * x) + (12 - (y - actualMonthI))
	}

}
func actualMonth() int {
	now := time.Now()
	_, currentMonth, _ := now.Date()
	i := int(currentMonth)
	return int(i)
}

func monthsOfAge(x person) {
	fmt.Println(x.name+"'s Months of age:", months(x.age, x.bornOn))
}
func newPersona() person {
	var xName string
	var xAge, xMOB int
	fmt.Print("Enter Name: ")
	fmt.Scanln(&xName)
	fmt.Print("Enter Age: ")
	fmt.Scanln(&xAge)
	for xAge < 0 {
		fmt.Print("Please enter a valid age: ")
		fmt.Scanln(&xAge)
	}
	fmt.Print("Enter Month of birth: ")
	fmt.Scanln(&xMOB)
	for xMOB < 1 || xMOB > 12 {
		fmt.Print("Please enter a month of birth between 1-12: ")
		fmt.Scanln(&xMOB)
	}
	newPerson := person{name: xName, age: xAge, bornOn: xMOB}
	return newPerson
}
