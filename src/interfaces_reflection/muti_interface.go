package main

import "fmt"

type SalaryCalculator interface {
	DisplaySalary()
}

type LeaveCalculator interface {
	CalculateLeavesLeft() int
}

type EmployeeOperations interface {
	SalaryCalculator
	LeaveCalculator
}

type Employee struct {
	firstName   string
	lastName    string
	basicPay    int
	pf          int
	totalLeaves int
	leavesTaken int
}

func (e Employee) DisplaySalary() {
	fmt.Printf("%s %s has salary $%d", e.firstName, e.lastName, (e.basicPay + e.pf))
}
func (e Employee) CalculateLeavesLeft() int {
	return e.totalLeaves - e.leavesTaken
}

func main() {
	e := Employee{
		firstName:"ttxs",
		lastName:"goto",
		basicPay: 3000,
		pf: 2000,
		totalLeaves: 4000,
		leavesTaken:3000,
		
	}
	//var s SalaryCalculator = e
	////s = e
	//s.DisplaySalary()
	//var l LeaveCalculator = e
	//fmt.Println("\nLeaves left =", l.CalculateLeavesLeft())
	
	var empOp EmployeeOperations = e
	empOp.DisplaySalary()
	fmt.Println("\nLeaves left =", empOp.CalculateLeavesLeft())
	var nul EmployeeOperations
	if nul == nil {
		// 使用nil的接口调用一个方法，则程序会panic，
		// 因为nil interface既没有底层的值也没有对应的具体类型.
		fmt.Printf("nul is nil and 类型 %T 值%v\n", nul, nul)
	}
}

