package account

import (
	"errors"
	"math/rand/v2"
	"net/url"
	"github.com/fatih/color"
)

type account struct {
	login    string
	password string
	url      string
}

func (acc account) OutputPassword() {
	color.Cyan(acc.login)
}

func (acc *account) generatePassword(n int) {
	res := make([]rune, n)
	for i := range res {
		res[i] = lettersRunes[rand.IntN(len(lettersRunes)-1)]
	}
	acc.password = string(res)
}
func NewAccount(loginStr, password, urlString string) (*account, error) {
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("Invalid URL")
	}
	if len(loginStr) == 0 {
		return nil, errors.New("Invalid Login")
	}

	newAcc := &account{
		url:      urlString,
		login:    loginStr,
		password: password,
	}

	if password == "" {
		newAcc.generatePassword(12)
	}
	return newAcc, nil

}

var lettersRunes = []rune("abcdefghijklmnopqrstvuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789-*!")
