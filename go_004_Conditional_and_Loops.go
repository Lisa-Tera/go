package main

import "fmt"

func fact(n int)  int{
	result :=1
	for i :=2;i<=n;i++{
		result = result*i
	}
	return result
}
func main()  {
	fmt.Println(fact(6))
	//second := time.Now().Unix()
	//rand.Seed(second)
	//target := rand.Intn(100)+1// 컴퓨터 난수가 실제로 난수가 아니여서 같은 값이 나옴
	//fmt.Println(target)
}