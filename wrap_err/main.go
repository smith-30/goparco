package main

import (
	"errors"
	"fmt"
)

var AnyHttpErr = errors.New("http err")
var ErrNotFound = errors.New("not found")
var ErrUnexpected = errors.New("ErrUnexpected")

type NotFoundError struct {
	Name string
}

type QueryError struct {
	Query string
	Err   error
}

func (e *QueryError) Error() string {
	return e.Query + ": " + e.Err.Error()
}

// err から別の err を出力したいときに使う
// err がある程度の種類の型を持つときに使う感じかな
// PathError には組み込まれている
func (e *QueryError) Unwrap() error {
	return e.Err
}

func main() {
	e := ErrNotFound
	otherErr := errors.New("other err")
	fmt.Printf("%#v\n", errors.Unwrap(e))

	// Is は err の値を検証する
	fmt.Printf("%#v\n", errors.Is(e, ErrNotFound))
	fmt.Printf("%#v\n", errors.Is(otherErr, ErrNotFound))

	qerr := &QueryError{
		Query: "select * from ",
		Err:   ErrUnexpected,
	}

	fmt.Printf("%#v\n", qerr)

	fmt.Printf("%#v\n", qerr.Error())
	fmt.Printf("%#v\n", errors.Unwrap(qerr))
	fmt.Println()

	var qe *QueryError
	fmt.Printf("qe %#v\n", qe)
	fmt.Printf("*qe %#v\n", &qe)
	fmt.Printf("%#v\n", errors.As(e, &qe))
	fmt.Printf("%#v\n", errors.As(qe, &qe))
	wrappedqe := wrap(qerr)
	// As は特定のtype かどうか調べる。type assertion みたいな感じ。
	// Wrap されたチェインをすべて検証して結果を返す
	fmt.Printf("wrappedqe %#v\n", errors.As(wrappedqe, &qe))
	fmt.Println()

	httpErr := wrap(AnyHttpErr)
	fmt.Printf("%#v\n", httpErr)
	if errors.Is(httpErr, AnyHttpErr) {
		fmt.Printf("** hit err is ** %#v\n", httpErr)
	}
}

// %w で wrap したものを返せる。
// &fmt.wrapError{msg:"wrap: http err", err:(*errors.errorString)(0xc0000101f0)} になる
func wrap(err error) error {
	return fmt.Errorf("wrap: %w", err)
}
