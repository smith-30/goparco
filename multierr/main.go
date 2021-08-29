package main

import (
	"errors"
	"fmt"

	"github.com/hashicorp/go-multierror"
)

var (
	Err1  = errors.New("Err1")
	Err2  = errors.New("Err2")
	Err10 = errors.New("Err10")
)

func checkLinkGuides() error {
	var result error

	result = multierror.Append(result, Err1)
	result = multierror.Append(result, Err2)

	return result
}

func main() {
	err := checkLinkGuides()

	fmt.Printf("%#v\n", err.(*multierror.Error).Errors)
	fmt.Printf("%#v\n", errors.Is(err, Err1))
	fmt.Printf("%#v\n", errors.Is(err, Err10))

	for _, item := range err.(*multierror.Error).Errors {
		fmt.Printf("%#v\n", item)
	}
}
