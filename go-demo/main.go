package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	fmt.Println("__ Калькулятор индекса массы тела ___")

	for {

		userKg, userHeight := getUserInput()
		IMT, err := calculateIMT(userKg, userHeight)
		if err != nil {
			// fmt.Println("Не заданы параметры для расчета")
			// continue
			panic("Что-то пошло совсем не так")
		}
		outputResult(IMT)
		isRepeatCalculation := chekRepeatCalculation()
		if !isRepeatCalculation {
			break
		}
	}

}
func outputResult(imt float64) {

	switch {
	case imt < 16:
		fmt.Println("У вас сильный дефицит массы тела")
	case imt < 18.5:
		fmt.Println("У вас дефицит массы тела")
	case imt < 25:
		fmt.Println("У вас нормальный вес")
	case imt < 30:
		fmt.Println("У вас избыточный вес")
	default:
		fmt.Println("У вас оширение")

	}
	result := fmt.Sprintf("Ваш индекс массы тела: %.0f ", imt)
	fmt.Println(result)
}

func calculateIMT(userKg float64, userHeight float64) (float64, error) {
	const IMTPower = 2
	if userKg <= 0 || userHeight <= 0 {

		return 0, errors.New("NO_PARAMS_ERROR")
	}

	IMT := userKg / math.Pow(userHeight/100, IMTPower)
	return IMT, nil
}

func getUserInput() (float64, float64) {

	var userHeight float64
	var userKg float64
	fmt.Print("Введите свой рост в сантиметрах: ")
	fmt.Scan(&userHeight)
	fmt.Print("Введите свой вес: ")
	fmt.Scan(&userKg)

	return userKg, userHeight

}

func chekRepeatCalculation() bool {
	var userChoise string
	fmt.Print("Еще один расчет? (y/n): ")
	fmt.Scan(&userChoise)
	if string(userChoise) == "y" {
		return true
	}
	return false

}
