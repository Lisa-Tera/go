package main

import (
	"fmt"
	"reflect"
)

func main() {
	var id,z int //변수 선언
	var x,y = 3.14, 2.51 // float64 타입 생략 OK
	//var name string
	var human bool=true//선언과 동시에 초기화
	var a string
	var b bool
	
	id =66
	//id =true // error 선언한 변수의 타입을 바꿀 수 없다
	//x=0
	//y=0
	//x,y =0,10.9
	name := "trevor" //C++ auto,Python처럼 변수 선언과 동시에 타입 지정 없이 초기화 가능함.
	/*
	c++ 정적 타입으로 처리가 되기 때문에 컴파일 시점이 결정됨 go도
	auto name ="trevor"s;//string

	Python 동적 바인딩이기 때문에 인터프리터가 해석하는 시점에 값을 바꿈
	name ="trevor"; //string
	 */
	fmt.Println(a)// 문자열의 제로값: 빈 문자열
	fmt.Println(b)//bool 타입의 제로값 : false
	fmt.Println(z)//int 타입의 제로값 :0 		값을 할당하지 않고 선언한 변수는 제로 값을 갖게됨
	fmt.Println(human)
	fmt.Println(id)
	fmt.Println(x,y)
	//fmt.Println(x)
	//fmt.Println(y)
	fmt.Println(name)
	fmt.Println(reflect.TypeOf(name))




	//fmt.Println(math.Ceil("go"))
	//fmt.Println(strings.Title(2.71))
	/*
	fmt.Println(reflect.TypeOf(2.71))
	fmt.Println(reflect.TypeOf("go"))
	fmt.Println(reflect.TypeOf('B'))
	fmt.Println(reflect.TypeOf(false))
	fmt.Println(reflect.TypeOf(55))
	 */
}
