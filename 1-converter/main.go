package main

import "fmt"

func main() {
	const USD_TO_EUR = 0.94 // 1 USD = 0.94 EUR
	const USD_TO_RUB = 100 // 1 USD = 90 RUB
	const EUR_TO_RUB = USD_TO_RUB / USD_TO_EUR

	fmt.Println("Евро к рублю = ", EUR_TO_RUB)
}