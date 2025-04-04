package main

import (
	"fmt"
	"math"
)

func main() {
	var userHeight float64
	var userKg float64

	fmt.Println("___ Калькулятор индекса массы тела ___")
	fmt.Print("Введите свой рост в сантиметрах: ")
	fmt.Scan(&userHeight)

	fmt.Print("Введите свой вес: ")
	fmt.Scan(&userKg)

	IMT := calculateIMT(userKg, userHeight)
	outputResult(IMT)
}

func outputResult(IMT float64) {
	result := fmt.Sprintf("Ваш индекс массы тела: %.0f", IMT)
	fmt.Print(result)
}

func calculateIMT(userKg float64, userHeight float64) float64 {
	const IMTPower = 2
	IMT := userKg / math.Pow(userHeight / 100, IMTPower)
	return IMT
}