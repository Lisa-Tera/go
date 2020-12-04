package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	fmt.Println('A')// 출력 :65 unicode로 출력  룬
	fmt.Println('\n')//2byte
	fmt.Println("\n")
	fmt.Println(true)
	fmt.Println(5!=5)
	fmt.Println("Hello, Go!") //;찍어도 실행은 되지만 없어도 돌아감
	fmt.Println(math.Ceil(3.14))
	fmt.Println(strings.Title("\"hello\" \t go~"))
}
