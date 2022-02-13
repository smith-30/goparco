package domain

import (
	"errors"
)

type Password interface {
	Value() (string, error)
	SecureString
}

type password struct {
	*secureString
}

func NewPassword(val string) Password {
	return &password{
		secureString: newSecureString(val),
	}
}

func (a *password) Value() (string, error) {
	if a.secureString == nil {
		// こういうときに例外投げたくなるなぁ
		return "", errors.New("can't get value. you")
	}
	str := a.String()
	a.secureString = nil
	return str, nil
}
