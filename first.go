package main

import "fmt"

func main(){
	primeArray :=[5]int{2,3,5,7,11}
	primeSlice :=[]int{2,3,5}
	/*
	var primeSlice []int
	primeSlice =make([]int,3)
	primeSlice[0] =2
	primeSlice[1] =3
	primeSlice[2] =5 */
	fmt.Println(primeArray[2])
	fmt.Println(primeSlice[2])

	/*
	var dates [4]time.Time
	dates[0] =time.Unix(12357894000,0)
	dates[1] =time.Unix(1447920001,0)
	dates[2] =time.Unix(977750000,0)

	fmt.Println(dates[3])//초기화

	 */
	//var f float64
	//f =10
	/*
	primes := [5]int {2,3,5,7,11}//선언과 동시에 초기화  var 은 := 때문에 생략
	var sum int =0
	for _,v:=range primes{//value 출력
		sum+=v
	}
	fmt.Printf("%0.2f\n",float64(sum)/float64(len(primes)))
	*/

	/*
	for i,v:=range primes{//for each로 보면 됨
		fmt.Println(i,v)
	}
	for _,v:=range primes{//value 출력
		fmt.Println(v)
	}
	for i:=range primes{//인덱스 값 출력
		fmt.Println(i)
	}
	for k :=0;k< len(primes);k++{//panic runtime Error
		fmt.Println(primes[k])
	}*/
	//fmt.Println(primes[1],primes[4])//초기화 0으로
	//fmt.Println(f/float64(primes[1]))//f/prime[1] =>mismatched types float64 and int
	/*
	var notes [7] string
	notes[0] = "도"
	notes[1] = "레"
	notes[2] = "미"
	fmt.Println(notes[0],notes[1])
	*/

}