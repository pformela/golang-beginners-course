package main

import (
	"fmt"
)

func main() {
	statePopulations := make(map[string]int)
	statePopulations = map[string]int{
		"California":  39250017,
		"Texas":       27862596,
		"Florida":     20612439,
		"New York":    19745289,
		"Ohio":        11614373,
		"Pensylvania": 12802503,
		"Illinois":    12802503,
	}

	statePopulations["Georgia"] = 10310371

	fmt.Println(statePopulations)
	fmt.Println(statePopulations["Ohio"])
	fmt.Println(statePopulations["Georgia"])

	delete(statePopulations, "Georgia")

	fmt.Println(statePopulations)
}
