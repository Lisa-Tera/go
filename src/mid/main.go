package main

import (
	"fmt"
	"mathcalculator"
	"os"
)

func main()  {
	for x := 1; true; x++ {
		fmt.Print("1) 절대값 2) 팩토리얼 3) 피보나치 4) 거듭제곱 5)삼각수 6)오각수 7)팩트리온 8)소수 9)조합 10) 종료 : ")
		n := mathcalculator.Input()
		if n == 1 {
			fmt.Print("정수 입력(절대값 계산) : ")
			fmt.Println(mathcalculator.Absolute(mathcalculator.Input()))
		}else if n == 2 {
			fmt.Print("정수 입력(팩토리얼 계산) : ")
			fmt.Println(mathcalculator.Factorial(mathcalculator.Input()))
		} else if n == 3 {
			fmt.Print("정수 입력(피보나치 출력) : ")
			f := mathcalculator.Input()
			fmt.Println(mathcalculator.Fibonacci(f))
		} else if n == 4 {
			fmt.Print("Base 입력 : ")
			b := mathcalculator.Input()
			fmt.Print("Exponent 입력 : ")
			e := mathcalculator.Input()
			fmt.Println(mathcalculator.Power(b, e))
		}else if n == 5 {
			fmt.Print("정수 입력(삼각수 계산) : ")
			t:= mathcalculator.Input()
			fmt.Println(mathcalculator.Triangular(t))
		}else if n == 6 {
			fmt.Print("정수 입력(오각수 계산) : ")
			f:= mathcalculator.Input()
			fmt.Println(mathcalculator.Pentagonal(f))
		} else if n == 7 {
			var num string
			fmt.Print("정수 입력(팩토리온 확인) : ")
			fmt.Scanln(&num)
			fmt.Println(" is",mathcalculator.Factorion(num),".")

		} else if n == 8 {
			fmt.Print("소수 판정  입력 : ")
			p :=mathcalculator.Input()
			fmt.Print(p," is")
			if mathcalculator.Prime(p) ==true{
				fmt.Println(" prime number ")
			}else{
				fmt.Println(" not prime number ")
			}
		}else if n == 9 {
			fmt.Print("n  : ")
			n:=mathcalculator.Input()
			fmt.Print("r  : ")
			r:=mathcalculator.Input()
			fmt.Println(mathcalculator.Combination(n,r))
		}   else if n == 10 {
			os.Exit(3)
		} else {
			fmt.Print("잘못 된 입력 값입니다. 1~10사이의 수를 입력하세요!")
		}
	}
}