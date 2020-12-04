/*
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)
func absolute(n int ) int{ //절대값 계산 함수로
	if n<0 {//n이 음수이므로
		return n * -1 //양수로 리턴해준다
	}
	return n
}
func factorial(n int) int {//팩토리얼 계산 함수로
	if n==0{ //n이 0 일 때 1을 리턴해준다
		return 1
	}
	return n * factorial(n -1) // n * (n-1)!으로 n*(n-1)*(n-2)...가 되도록 해준다
}
func fibonacci(n int)int{//피보나치 계산 함수로
	if n <=2{ //n이 1이거나 2이면
		return 1 //1을 리턴해주고
	}else { //n이 2보다 크면
		n1 :=1
		n2 :=1 //n1,n2를 설정해서

		for i:=2;i<=n;i++{ //Fn = Fn-1 + Fn-2이므로 i를 2로 설정하고 n일 때 까지 실행한다
		n1,n2 = n2,n1+n2 // n1에 n2 값을 넣어주고, n2에 n1+n2값을 넣어 Fn = Fn-1 + Fn-2이 성립되도록 해준다
		}
		return n1
	}
}
func exponentiation(b int,e int) int {//거듭제곱 계산 함수로
	n :=1
	for i:=1;i<=e;i++{ //밑을 곱하는 횟수를 지수의 수 만큼 반복해준다
		n *=b
	}
	return n
}

func main(){
	for true {
		fmt.Printf("1)절대값 2)팩토리얼 3)파보나치 4)거듭제곱 5)종료 : ")
		//reader를 이용해서 입력값을 읽어온다.
		reader := bufio.NewReader(os.Stdin)
		s, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		s = strings.TrimSpace(s)
		//변수 2개 설정해서 1~5번에서 사용해 줄 것을 설정해준다.
		var num int
		var num2 int
		//fmt.Printf("%T ",in1)
		if s =="1"{
			fmt.Printf("정수 입력 (절대값 계산) : ")
			fmt.Scanln(&num)
			fmt.Println(absolute(num))
		}else if s =="2"{
			fmt.Printf("정수 입력 (팩토리얼 계산) : ")
			fmt.Scanln(&num)
			if num >0 {
				fmt.Println(factorial(num))
			}else { //num이 1보다 작을 경우
				fmt.Println("잘못입력했습니다")
			}
		}else if s =="3"{
			fmt.Printf("정수 입력 (피보나치 출력) : ")
			fmt.Scanln(&num)
			if num >0 {
				fmt.Println(fibonacci(num))
			}else { //num이 1보다 작을 경우
				fmt.Println("잘못입력했습니다")
			}
		}else if s =="4"{
			fmt.Printf("Base 입력 : ")//밑
			fmt.Scanln(&num)
			fmt.Printf("Exponent 입력 : ")//지수
			fmt.Scanln(&num2)
			fmt.Println(exponentiation(num,num2))
		}else if s =="5"{
			os.Exit(3) //os.Exit를 사용해서 종료해준다
		}else {
			fmt.Println("잘못 입력하셨습니다. 다시 입력하세요")
		}
	}

}*/