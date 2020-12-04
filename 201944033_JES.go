package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main()  {

	//1.첫 번째 정수 입력
	fmt.Println("첫 번째 정수 입력 : ")
	reader := bufio.NewReader(os.Stdin) //reader :읽기
	//ReadString \n 앞까지 내용을 in변수에 담음 , error가 났을 때 에러에 대한 공간
	in1 ,err:= reader.ReadString('\n')//첫번째 정수 입력한 것을 in1에 입력값을 넣음
	//error가 nil일 시 처리
	if err!=nil{
		log.Fatal(err)
	}
	in1 =strings.TrimSpace(in1)//줄바꿈 룬 제거
	i1, err := strconv.ParseInt(in1,10,64) //in1을 int64로 변환한 값을 i1에 넣어준다.

	//2.2번재 정수 입력
	fmt.Println("두 번째 정수 입력 : ")

	in2 ,err:= reader.ReadString('\n')//두 번째 정수 입력한 것을 in2에 입력값을 넣음
	if err!=nil{
		log.Fatal(err)
	}
	in2 =strings.TrimSpace(in2)
	i2, err := strconv.ParseInt(in2,10,64)//in2을 int64로 변환한 값을 i2에 넣어준다.

	//fmt.Println(i2,i1) 값 확인
	//1~50까지 for문을 사용해서 확인해준다.
	for i :=1; i<=50;i++ {
		//i가 i1의 배수일 경우
		if i%int(i1) == 0{
			//i가 i1,i2의 공배수일 경우
			if i%int(i2) == 0{
				fmt.Println(i1, i2,"의 공배수",i)
			}else {
				fmt.Println(i1, "의 배수")
			}
		}else if i%int(i2) == 0{ //i가 i2의 배수일 경우
			fmt.Println(i2,"의 배수")
		}else { //i가 i1 or i2의 배수가 아닌 경우
			fmt.Println(i)
		}
	}
}

