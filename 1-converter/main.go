package main

import (
	"errors"
	"fmt"
)

// 1. Сделать меню с шагами
// 2. Выделить функции
// 3. После получения всех данных с помощью if / switch вычислить итог и вывести результат.

const USD_TO_EUR = 0.94                    // 1 USD = 0.94 EUR
const USD_TO_RUB = 100                     // 1 USD = 90 RUB
const EUR_TO_RUB = USD_TO_RUB / USD_TO_EUR // 106
const EUR_TO_USD = 1 / USD_TO_EUR

func main() {
	currency1, amount, currency2 := handleUserInput()

	fmt.Println("Вы получите = ", convert(currency1, amount, currency2))
}

func convert(currency1 string, amount float64, currency2 string) float64 {
	switch {
	case currency1 == "USD" && currency2 == "EUR":
		return amount * USD_TO_EUR
	case currency1 == "USD" && currency2 == "RUB":
		return amount * USD_TO_RUB
	case currency1 == "EUR" && currency2 == "USD":
		return amount * EUR_TO_USD
	case currency1 == "EUR" && currency2 == "RUB":
		return amount * EUR_TO_RUB
	case currency1 == "RUB" && currency2 == "USD":
		return amount / USD_TO_RUB
	case currency1 == "RUB" && currency2 == "EUR":
		return amount / EUR_TO_RUB
	}

	return amount
}

// 2. Выделить функцию ввода / проверки валюты и числа
func handleInputCurrency() (string, error) {
	var currency string

	fmt.Scan(&currency)

	if currency != "USD" && currency != "EUR" && currency != "RUB" {
		return "", errors.New("Такой валюты нет")
	}

	return currency, nil
}

// 2. Выделить функцию ввода / проверки валюты и числа
func handleInputAmount() (float64, error) {
	var amountInput float64

	fmt.Scan(&amountInput)

	if amountInput <= 0 {
		return 0, errors.New("Введено некорректное значение")
	}

	return amountInput, nil

}

// Обрабатываем пользовательский ввод
func handleUserInput() (string, float64, string) {
	var currency1 string
	var amount float64
	var currency2 string

	// Ввод исходной валюты (подсказываем варианты) - если ошибка, заново вводим

	for true {
		fmt.Print("Какую валюту вы желаете обменять (USD, EUR, RUB) ? ")
		currencyInput, error := handleInputCurrency()

		if error != nil {
			continue
		} else {
			currency1 = currencyInput
			break
		}

	}

	// Ввод числа - если ошибка, заново вводим
	for true {
		fmt.Print("Введите сумму обмена: ")
		amountInput, error := handleInputAmount()
		if error != nil {
			continue
		} else {
			amount = amountInput
			break
		}
	}

	// Ввод целевой валюты (подсказываем варианты, кроме исходной валюты) - если ошибка, заново вводим
	for true {
		fmt.Print("Какую валюту вы желаете получить? ")
		currencyInput, error := handleInputCurrency()

		if error != nil {
			continue
		} else if currencyInput == currency1 {
			fmt.Println("Целевая валюта не может быть равна исходной")
			continue
		} else {
			currency2 = currencyInput
			break
		}

	}

	return currency1, amount, currency2
}
