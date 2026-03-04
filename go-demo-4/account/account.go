package account

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"reflect"
	"time"

	"github.com/fatih/color"
)

type Account struct {
	Login     string    `json:"login" `
	Password  string    `json:"passwordn" `
	Url       string    `json:"url" `
	CreatedAt time.Time `json:"creatAt" `
	UpdateAt  time.Time `json:"updateAt" `
}

func (acc Account) OutputPassword() {
	color.Cyan(acc.Login)
}

func (acc Account) ToBytes() ([]byte, error) {
	file, err := json.Marshal(acc)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (acc *Account) generatePassword(n int) {
	res := make([]rune, n)
	for i := range res {
		res[i] = lettersRunes[rand.IntN(len(lettersRunes)-1)]
	}
	acc.Password = string(res)
}
func NewAccount(loginStr, password, urlString string) (*Account, error) {
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("Invalid URL")
	}
	if len(loginStr) == 0 {
		return nil, errors.New("Invalid Login")
	}

	newAcc := &Account{
		CreatedAt: time.Now(),
		UpdateAt:  time.Now(),
		Url:       urlString,
		Login:     loginStr,
		Password:  password,
	}

	field, _ := reflect.TypeOf(newAcc).Elem().FieldByName("Login")
	fmt.Println(string(field.Tag))

	if password == "" {
		newAcc.generatePassword(12)
	}
	return newAcc, nil

}

var lettersRunes = []rune("abcdefghijklmnopqrstvuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789-*!")
