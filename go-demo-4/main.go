package main

import (
	"demo/password/account"
	"demo/password/files"
	"fmt"
)

func main() {
	fmt.Println("__Менеджер паролей__")

Menu:
	for {
		variant := getMenu()
		switch variant {
		case 1:
			createAccount()
		case 2:
			findAccount()
		case 3:
			deleteAccount()
		default:
			break Menu
		}
	}
}

func getMenu() int {
	var variant int
	fmt.Println("Выберите вариант: ")
	fmt.Println("1. Создать аккаунт.")
	fmt.Println("2. Найти аккаунт.")
	fmt.Println("3. Удалить аккаунт.")
	fmt.Println("4. Выйти из меню.")
	fmt.Scanln(&variant)
	return variant
}

func findAccount() {
}

func deleteAccount() {
}

func createAccount() {

	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите URL")
	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		fmt.Println("Неверный формат URL или Login")
		return
	}

	vault := account.NewVault()
	vault.AddAccount(*myAccount)
	if err != nil {
		fmt.Println("Не удалось преобразовать в JSON")
		return
	}
	data, err := vault.ToBytes()
	files.WriteFile(data, "data.json")

	//files.WriteFile("C:\\Users\\sergeev_dy\\Desktop\\КодSQL для выдачи прав при создании таблицы.txt", "Какой-то текст в файл.")

}
func promptData(prompt string) string {

	fmt.Print(prompt + ": ")
	var res string
	fmt.Scanln(&res)
	return res
}
