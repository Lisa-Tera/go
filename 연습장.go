package main

import (
	"fmt"
	"strconv"
)

func main(){
	//var m map[int] string
	var li[] string;
	var a string
	var rows int
	var temp int = 1
	fmt.Print("Enter number of rows : ")
	fmt.Scan(&rows)

	for i := 0; i < rows; i++ {
		for k := 0; k <= i; k++ {
			if (k==0 || i==0) {
				temp = 1
				//fmt.Print(" %d",temp)
			}else{
				temp = temp*(i-k+1)/k
			}
			a+= strconv.Itoa(temp) +" "
		}
		//fmt.Println(a)
		li =append(li, a)

		a=""
	}
	fmt.Println(li[rows-1])
}