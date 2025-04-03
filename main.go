package main

import (
	"fmt"
	"math"
)

func main() {
	userHeight, userKg := 1.8, 100.0
	IMT := userKg / math.Pow(userHeight, 2)
	fmt.Print(IMT)
}