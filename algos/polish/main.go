package main

import (
	"fmt"
	"strconv"
	"strings"
)

// v を順番に見ていく
//   数が登場したら stack に push
//   演算子が登場したら
//     stack から pop して取り出した値を a
//     続けて stack から pop して取り出した値を b
//     b, a の計算結果をスタックに push する (b が計算の最初にくる。下記プログラム参照)
// stack に残っている数を返す
func RPolishCalc(v []string) float64 {
	stack := []float64{}
	for _, item := range v {
		i, ok := strconv.Atoi(item)
		if ok == nil {
			stack = append(stack, float64(i))
		} else {
			poppedVal_a := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			poppedVal_b := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			switch item {
			case "+":
				stack = append(stack, poppedVal_b+poppedVal_a)
			case "-":
				stack = append(stack, poppedVal_b-poppedVal_a)
			case "*":
				stack = append(stack, poppedVal_b*poppedVal_a)
			case "/":
				stack = append(stack, poppedVal_b/poppedVal_a)
			}
		}
		fmt.Printf("%#v\n", stack)
	}
	return stack[0]
}

func main() {
	v := "6 1 2 + * 8 -"
	fmt.Printf("%#v\n", RPolishCalc(strings.Split(v, " ")))
}
