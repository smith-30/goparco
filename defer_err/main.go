package main

import (
	"errors"
	"fmt"
)

func main() {
	defer func() {
		if p := recover(); p != nil {
			fmt.Printf("%#v\n", "main panic")
			panic(p)
		}
	}()


	
	v, n, p := 0, 1, 2
	fmt.Printf("%#v\n", handle(v))
	fmt.Println()

	fmt.Printf("%#v\n", handle(n))
	fmt.Println()

	fmt.Printf("%#v\n", handle(p))
}

func handle(v int) error {
	var err error

	// この時点ではnil なので defer でポインタは使えない
	fmt.Printf("err %p\n", err)

	defer func() {
		fmt.Printf("defer %#v\n", err)
		if p := recover(); p != nil {
			// err = tx.Rollback()
			fmt.Printf("%#v\n", "handle panic")
			panic(p) // re-throw panic after Rollback
		} else if err != nil {
			// err = tx.Rollback()
			err = errors.New("rollback")
		} else {
			// err = tx.Commit()
			err = nil
		}
		fmt.Printf("%#v\n", err)
	}()

	err = do(v)
	return err
}

func do(v int) error {
	if v == 0 {
		return errors.New("0")
	}

	if v == 2 {
		panic(errors.New("2"))
	}
	return nil
}
