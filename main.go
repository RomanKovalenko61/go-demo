package main

import (
	"fmt"
	"math"
)

func main() {
	const IMTPower = 2
	var userHeight float64
	var userKg float64

	fmt.Print("___ Калькулятор индекса массы тела ___ \n")
	fmt.Print("Введите свой рост в метрах: ")
	fmt.Scan(&userHeight)

	fmt.Print("Введите свой вес: ")
	fmt.Scan(&userKg)
	fmt.Print("Ваш индекс массы тела: ")

	IMT := userKg / math.Pow(userHeight, IMTPower)
	fmt.Print(IMT)
}