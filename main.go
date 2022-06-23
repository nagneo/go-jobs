package main

import (
	"fmt"
	"strings"
)

// func, defer
func lenAndUpper(name string) (length int, uppercase string) {
	defer fmt.Println("I'm done")
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

// for
func superAdd(numbers ...int) int {
	sum := 0
	for number := range numbers {
		sum += number
	}

	return sum
}

// if, else
func canIDrink(age int) bool {
	if koreanAge := age + 2; koreanAge < 18 {
		return false
	}

	return true
}

// Pointer
func exercisePointer() {
	a := 2
	b := &a
	c := a
	*b = 2022
	c = 20
	fmt.Println(&a, b, &c)
	fmt.Println(a, *b, c)
}

// Arrays and Slices
func exerciseArray() {
	names := []string{"nico", "lynn", "dal"}
	names = append(names, "flynn")
	fmt.Println(names)
}

// map
func exerciseMap() {
	nico := map[string]string{"name": "nico", "age": "12"}
	//fmt.Println(nico)
	for _, value := range nico {
		fmt.Println(value)
	}
}

type person struct {
	name    string
	age     int
	favFood []string
}

func exerciseStruct() {
	favFood := []string{"kimchi", "ramen"}
	nico := person{name: "nico", age: 18, favFood: favFood}
	fmt.Println(nico)
}

func main() {
	
}
