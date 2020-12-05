package main

import (
	"fmt"
	"math"
)

type Shape interface {//도형이라는 인터페이스 생성
	area() float64 //도형은 면적과 둘레를 가지므로  생성
	perimeter() float64
}
type Rect struct {
	width, height float64
}
type Circle struct {
	radius float64
}

func (r Rect) area() float64 {
	return r.width * r.height
}
func (r Rect) perimeter() float64{
	return 2* r.width + 2* r.height
}
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c Circle) perimeter() float64{
	return 2*math.Pi + 2* c.radius
}
func measure(s Shape){ //shape 인터페이스에 해당되는게 다 올 수 있음
	fmt.Println(s)
	fmt.Println(s.area())
	fmt.Println(s.perimeter())
}
func showArea(shapes ...Shape)  {
	for _,shape := range shapes{ //인덱스 없이 값만 shape로 가져와서 씀
		fmt.Println(shape.area())
	}
}
func printlf(itf interface{})  {
	fmt.Println(itf)
}
func main()  {
	var x interface{} //요런식으로 empty interface 사용 가능
	r := Rect{width: 5,height: 2}
	c := Circle{radius: 11}
	x =55
	x = "Test"
	printlf(x)
	showArea(c,r)

	measure(r)
	measure(c)
}
//빈 인터페이스는 어떤 타입도 담을 수 있는 컨테이너 java에서 object,c/c++ void 타입의 포인터