package main

import (
	"fmt"

	"github.com/smith-30/goparco/secure_by/domain"
)

func main() {
	p := domain.NewPassword("password")
	fmt.Printf("%#v\n", p)
	fmt.Printf("%v\n", p)
	fmt.Printf("%v\n", fmt.Sprintf("%#v", p))
}
