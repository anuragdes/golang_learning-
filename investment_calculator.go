package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate float64 = 2.5
	var investmeentAmount float64 = 1000
	years := 10.0
	exptecdReturnRate := 5.5

	futureValue := investmeentAmount * math.Pow(1+exptecdReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}
