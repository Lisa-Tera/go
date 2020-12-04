package mathcalculator

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)
func Absolute(n int) int {
	if n < 0 {
		return n * -1
	}
	return n
}
func Factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * Factorial(n-1)
}
func Prime(value int) bool {
	for i := 2; i <= int(math.Floor(float64(value) / 2)); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}
func  Combination(n int,r int)int{
	result :=Factorial(n)/Factorial(r)
	return result
}
func Input() int {
	reader := bufio.NewReader(os.Stdin)
	in, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	in = strings.TrimSpace(in)
	number, err := strconv.Atoi(in)
	if err != nil {
		log.Fatal(err)
	}
	return number
}
func Triangular(n int) int {
	n = (n*n+n)/2
	return n
}
func Pentagonal(n int) int {
	n =Triangular(3*n-1)/3
	return n
}

func Factorion(num string) string {
	ma:= map[string]int{"0":1,"1":1,"2":2,"3":6,"4":24,"5":120,"6":720,"7":5040,"8":40320,"9":362880}
	var s int
	for i :=0;i<len(num);i++{
		n := num[i:i+1]
		s +=ma[n]
	}
	intn := strconv.Itoa(s)
	fmt.Print(num,"(",s,")")
	if intn == num {
		return "factorion"
	}
	return "not factorion"
}


func Fibonacci(n int) int {
	if n == 1 || n == 2 {
		return 1
	} else {
		return Fibonacci(n-1) + Fibonacci(n-2)
	}
}
func Power(b int, e int) int{
	r := 1
	for i := 1; i<=e ; i++{
		r = r * b
	}
	return r
}