package domain

import "fmt"

type SecureString interface {
	String() string
	GoString() string
	MarshalJSON() ([]byte, error)
	UnmarshalJSON(bs []byte) error
}

type secureString string

func newSecureString(v string) *secureString {
	r := secureString(v)
	return &r
}

func (a *secureString) String() string {
	return "*****"
}

func (a *secureString) GoString() string {
	return "*****"
}

func (a *secureString) MarshalJSON() ([]byte, error) {
	return nil, fmt.Errorf("can't marshal this value")
}

func (a *secureString) UnmarshalJSON(bs []byte) error {
	return fmt.Errorf("can't marshal this value")
}
