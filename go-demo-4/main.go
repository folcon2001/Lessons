package main

import (
	"demo/password/account"
	"demo/password/files"
	"fmt"
)

func main() {
	files.ReadFile()
	files.WriteFile("Привет! Я файл", "file.txt")
	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите URL")

	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		fmt.Println("Неверный формат URL или Login")
		return
	}

	myAccount.OutputPassword()

	//files.WriteFile("C:\\Users\\sergeev_dy\\Desktop\\КодSQL для выдачи прав при создании таблицы.txt", "Какой-то текст в файл.")

}
func promptData(prompt string) string {

	fmt.Print(prompt + ": ")
	var res string
	fmt.Scanln(&res)
	return res
}
