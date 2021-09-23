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
	}
	return stack[0]
}

func Decode(v []string) string {
	result := []string{}
	for _, item := range v {
		_, ok := strconv.Atoi(item)
		if ok == nil {
			result = append(result, item)
		} else {
			poppedVal_a := result[len(result)-1]
			result = result[:len(result)-1]

			poppedVal_b := result[len(result)-1]
			result = result[:len(result)-1]

			switch item {
			case "+":
				result = append(result, fmt.Sprintf("%v + %v", poppedVal_b, poppedVal_a))
			case "-":
				result = append(result, fmt.Sprintf("%v - %v", poppedVal_b, poppedVal_a))
			case "*":
				if 1 < len(poppedVal_a) {
					poppedVal_a = fmt.Sprintf("( %v )", poppedVal_a)
				}
				if 1 < len(poppedVal_b) {
					poppedVal_b = fmt.Sprintf("( %v )", poppedVal_b)
				}
				result = append(result, fmt.Sprintf("%v * %v", poppedVal_b, poppedVal_a))
			case "/":
				if 1 < len(poppedVal_a) {
					poppedVal_a = fmt.Sprintf("( %v )", poppedVal_a)
				}
				if 1 < len(poppedVal_b) {
					poppedVal_b = fmt.Sprintf("( %v )", poppedVal_b)
				}
				result = append(result, fmt.Sprintf("%v / %v", poppedVal_b, poppedVal_a))
			}
		}
	}
	return strings.Join(result, " ")
}

func main() {
	v := "5 4 6 + 5 / *"
	fmt.Printf("%v = %#v\n", Decode(strings.Split(v, " ")), RPolishCalc(strings.Split(v, " ")))
}
