package main

import "fmt"

type VowelsFinder interface {
	FindVowels() []rune
}

type SalaryCalculator interface {
	CalculateSalary() int
}

type Contract struct {
	empID    int
	basicpay int
}
type Permantent struct {
	empID    int
	basicpay int
	jj       int
}

func (c Contract) CalculateSalary() int {
	return c.basicpay
}

func (p Permantent) CalculateSalary() int {
	return p.basicpay + p.jj
	
}

func totalSum(s []SalaryCalculator) {
	expense := 0
	for _, v := range s {
		expense = expense + v.CalculateSalary()
	}
	fmt.Printf("SUM : %d\n", expense)
}

type Mystring string

func (ms Mystring) FindVowels() []rune {
	var vowels []rune
	for _, rune := range ms {
		if rune == 'a' || rune == 'e' || rune == 's' || rune == 'o' || rune == 'u' {
			vowels = append(vowels, rune)
		}
	}
	return vowels
}

func main() {
	//name := Mystring("Sam Anderson")
	//var v VowelsFinder
	//v = name
	//fmt.Printf("Vowels are %c", v.FindVowels())
	
	person1 := Permantent{1, 3000, 10000}
	person2 := Permantent{2, 3000, 15000}
	person3 := Contract{3, 3000}
	
	employees := [] SalaryCalculator{person1, person2, person3}
	totalSum(employees)
}





