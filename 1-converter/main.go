package main

import "fmt"

func main() {
	const USD_TO_RUB = 100
	const amountUsd = 5

	const totalUsdToOutcome = USD_TO_RUB * amountUsd

	fmt.Println("USD_TO_RUB", totalUsdToOutcome)
}