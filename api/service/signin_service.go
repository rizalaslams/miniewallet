package service

import "miniewallet/api/auth"

type sigInInterface interface {
	SignIn(auth.AuthDetails) (string, error)
}

type signInStruct struct{}

var (
	Authorize sigInInterface = &signInStruct{}
)

func (si *signInStruct) SignIn(authD auth.AuthDetails) (string, error) {
	token, err := auth.CreateToken(authD)
	if err != nil {
		return "", err
	}
	return token, nil
}
