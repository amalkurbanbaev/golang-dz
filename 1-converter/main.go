package main

import "fmt"

const USD_TO_EUR = 0.94 // 1 USD = 0.94 EUR
const USD_TO_RUB = 100  // 1 USD = 90 RUB
const EUR_TO_RUB = USD_TO_RUB / USD_TO_EUR

func convert(amount float64, currency1 string, currency2 string) float64 {
	// if currency1 == "USD" && currency2 == "EUR" {
	// 	return USD_TO_EUR / amount
	// }

	// return amount
}

func handleUserInput() (float64, string, string) {
	var amount float64
	var currency1 string
	var currency2 string

	fmt.Print("Введите сумму для обмена: ")
	fmt.Scan(&amount)

	fmt.Print("Какую валюту вы желаете обменять? ")
	fmt.Scan(&currency1)

	fmt.Print("Какую валюту вы желаете получить? ")
	fmt.Scan(&currency2)

	return amount, currency1, currency2
}

func main() {
	amount, currency1, currency2 :=  handleUserInput()

	fmt.Println("Вы получите = ", convert(amount, currency1, currency2))
}
