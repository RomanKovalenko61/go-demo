package main

import (
	"fmt"
	"math"
)

const IMTPower = 2

func main() {
	fmt.Println("___ Калькулятор индекса массы тела ___")
	userKg, userHeight := getUserInput()
	IMT := calculateIMT(userKg, userHeight)
	if IMT < 16 {
		fmt.Println("У вас сильный дефицит массы тела")
	}
	outputResult(IMT)
}

func outputResult(IMT float64) {
	result := fmt.Sprintf("Ваш индекс массы тела: %.0f", IMT)
	fmt.Print(result)
}

func calculateIMT(userKg float64, userHeight float64) (IMT float64) {
	IMT = userKg / math.Pow(userHeight / 100, IMTPower)
	return
}

func getUserInput() (float64, float64) {
	var userHeight float64
	var userKg float64
	fmt.Print("Введите свой рост в сантиметрах: ")
	fmt.Scan(&userHeight)
	fmt.Print("Введите свой вес: ")
	fmt.Scan(&userKg)
	return userKg, userHeight
}