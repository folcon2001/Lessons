package main

import "fmt"

func main() {

	transactions := []float64{}
	for {

		transaction := scanTransaction()
		if transaction == 0 {
			break
		}

		transactions = append(transactions, transaction)
	}

	balance := calculateBalance(transactions)
	fmt.Printf("Ваш баланс:  %.2f", balance)

}
func scanTransaction() float64 {
	var transaction float64
	fmt.Println("Введите транзакцию (n для выхода):")
	fmt.Scan(&transaction)
	return transaction
}

func calculateBalance(transactions []float64) float64 {
	var balance float64
	for _, value := range transactions {
		balance += value
	}
	return balance
