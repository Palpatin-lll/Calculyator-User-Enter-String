/*package main

import (
	"fmt"
	"os"
)

var firstNum int
var secondNum int

func main() {

	fmt.Println("Введите перове число: ")
	fmt.Fscanln(os.Stdin, &firstNum)
	fmt.Println("Введите второе число: ")
	fmt.Fscanln(os.Stdin, &secondNum)
	Arifmetik()
}

func Arifmetik() {

	var variant string
	fmt.Println("Выберите действие: Сложенеи '+', Вычитание '-', Умножение '*', Деленеи '/'.")
	fmt.Fscan(os.Stdin, &variant)

	if variant == "+" {
		fmt.Println("Сумма: ", firstNum+secondNum)
	}
	if variant == "-" {
		fmt.Println("Разница: ", firstNum-secondNum)
	}
	if variant == "*" {
		fmt.Println("Произведение: ", firstNum*secondNum)
	}
	if variant == "/" {
		fmt.Println("Частное: ", firstNum/secondNum)
	}
}
*/
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
/*
func Convert() string {
	var stroka string
	fmt.Println("Введите прмиер: ")
	fmt.Fscan(os.Stdin, &stroka)
	return stroka
}


	func Distvie() float64 {
		var stroka string
		var firstNum float64
		var secondNum float64
		var otvet float64
		var Primer int
		fmt.Println("Введите прмиер: ")
		fmt.Fscan(os.Stdin, &stroka)
		Primer, _ := strconv.Atoi(stroka)
		fmt.Println("Сумма: ", otvet)

		for stroka := 0; Primer < 100 && Primer[i] != 0; stroka++ {
			if stroka == '+' {
				fmt.Println("Сумма: ", otvet)
				otvet = firstNum + secondNum
			}

			if Convert() == '-' {
				fmt.Println("Сумма: ", otvet)
				otvet = firstNum - secondNum
			}
			if Convert() == '*' {
				fmt.Println("Сумма: ", otvet)
				otvet = firstNum * secondNum
			}
			if Convert() == '/' {
				fmt.Println("Сумма: ", otvet)
				otvet = firstNum / secondNum
			}
			if Convert() == ' ' {
				fmt.Println("Вывод ошибки, так как строка не является математической операцией.")
			}
		}
		return otvet
	}
*/
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
/*
package main

import (

	"fmt"
	"strings"
	"strconv"

)

const operation = "+-*"

	func main() {
		var x,y,z string
		fmt.Scan(&x,&y,&z)
	}

	func result(x,y,z string) string {
		ferstNum, _ := strconv.Atoi(x)
		secondtNum, _ := strconv.Atoi(x)

}

func Deistvie(x,y,z string){
}*/
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"strconv"
	"strings"
)

const op = "+-*/"

func main() {
	var x, a, y string
	fmt.Scan(&x, &a, &y)
	fmt.Println(result(x, a, y))
}

func result(x, a, y string) string {
	DontFound(a)

	first, err1 := strconv.Atoi(x)
	second, err2 := strconv.Atoi(y)
	if err1 == nil && err2 == nil {
		return strconv.Itoa(parsNumber(first, second, a))
	} else {
		return parsRome(x, a, y)
	}
}

func DontFound(a string) string {
	if strings.ContainsAny(a, op) {
		return a
	}
	panic("На вход поступил неверный арифметический знак")
}

func parsNumber(x, y int, a string) int {
	var otvet int
	if x < 11 && y < 11 {
		switch a {
		case "+":
			otvet = x + y
		case "-":
			otvet = x - y
		case "*":
			otvet = x * y
		case "/":
			otvet = x / y
		}
		return otvet
	} else {
		panic("Ошибка! Ожидаются только арбские либо только рисмские числа которые не больше 10")
	}

}

func parsRome(x, a, y string) string {
	numRome := map[string]int{
		"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5,
		"VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
	}
	firstRome, ok1 := numRome[x]
	secondRome, ok2 := numRome[y]
	if ok1 && ok2 {
		num := parsNumber(firstRome, secondRome, a)
		if num < 0 {
			panic("Ошибка!В римской системе нет отрицательных чисел.")
		}
		symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
		values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
		var result strings.Builder
		for i := 0; i < len(symbols); i++ {
			for num >= values[i] {
				result.WriteString(symbols[i])
				num -= values[i]
			}
		}
		return result.String()
	} else {
		panic("Ошибка! Ожидаются только арбские либо только рисмские числа которые не больше 10")
	}
}
