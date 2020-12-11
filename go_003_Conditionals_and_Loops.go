package main

import (
	"fmt"
	"time"
)

func main()  {
	/*
	a:=3
	a,b :=2,7
	fmt.Println(a,b)
	//var append string = "shadow" //append 변수로 선언
	//fmt.Println(append)
	var test = append([]string{},"테스트") //append 내장 함수로 동작  -> 둘 다 사용 X로 이름을 잘 정하자
	fmt.Println(test)


	fmt.Println("반지름 입력 : ")
	reader := bufio.NewReader(os.Stdin)
	//ReadString \n 앞까지 내용을 in변수에 담음 , error가 났을 때 에러에 대한 공간
	in ,err:= reader.ReadString('\n')//\n 앞 글자까지 읽기 // err에러시 사용
	//in, _ := reader.ReadString('\n')
	if err!=nil{
		log.Fatal(err)
	}
	in =strings.TrimSpace(in)//줄바꿈 룬 제거 Trimspace입력받은 문자열 안의 문자열을 제거(\n)
	radius, err := strconv.ParseFloat(in,64)//실수형으로 변환
	fmt.Println("원의 넓이: ",radius*radius*3.141592)
	if err!=nil{
		log.Fatal(err)
	}
*/
	//wrongtexts :="Gx Gx Gx!"
	//replacer := strings.NewReplacer("x","o")//repalcer가 메소드 이름
	//correctTexts := replacer.Replace(wrongtexts)
	//fmt.Println(correctTexts)

	a := time.Now()

	fmt.Println(a)
	var now time.Time= time.Now()//now에 연월일 시분초가 다 있음
	var year int =now.Year()
	month := now.Month()
	fmt.Println(year)
	fmt.Println(int(month))
	//var length float64 = 5.7
	//var width int =5
	//fmt.Println("면적은 ",length*width) error
	//fmt.Println("면적은 ",int(length)*width)
}
