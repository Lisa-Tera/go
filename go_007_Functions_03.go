package main

import (
	"fmt"
	"reflect"
)

func main(){
	//var id int = 55
	//var pid *int
	//pid = &id
	id :=55
	grade := 3.91
	pid := &id
	pgrade := &grade
	fmt.Println(id,&id,reflect.TypeOf(id))
	fmt.Println(*pid,pid,&pid, reflect.TypeOf(pid))
	fmt.Println(grade,&grade, reflect.TypeOf(grade))
	fmt.Println(*pgrade,grade,&pgrade, reflect.TypeOf(pgrade))

}

/*
package main
import "fmt"
func main(){
	number := 3
	cube(number)
	fmt.Println(number)
}
func cube(n int){
	n = n * n * n
	//fmt.Println(n)
}
*/