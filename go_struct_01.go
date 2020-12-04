package main

/*
type part struct {
		description string
		count int
}
func showInfo(p part) {
	fmt.Println("Description : ", p.description)
	fmt.Println("Count : ", p.count)
}
//리턴이 있는 구조 부품에 대한 설명인 d 부분만 받음
func minimumOrder(description string) part{
	var p part //구조체 선언
	p.description = description
	p.count= 1000
	return p //구조체 리턴
}
func main()  {
	 p := minimumOrder("Hex bolts")//변수 선언 후 함수 호출
	 fmt.Println(p.description,p.count)
	/*var bolts part
	bolts.description = "Hex bolts"
	bolts.count =100
	showInfo(bolts)
*/
}

/*
//4.요렇게 위에 type ___ struct를 쓰면 코드낭비가 적음
type subscriber struct{
	name string
	rate int
	active bool
}
func main(){
	var s1 subscriber
	fmt.Printf("%#v\n",s1) //제로값 출력 why 초기화 X

	s1.name ="kim"
	s1.rate = 5000
	s1.active = false
	fmt.Printf("%s\n",s1.name)
	fmt.Println(s1.rate)
	fmt.Println(s1.active)

		//3. struct 사용하면 됨!! (구독자가 1명일 경우)
		/!var subscriber struct{//2명이상일 때는 이걸 다 반복해야되서
			name string
			rate int
			active bool
		}
		fmt.Printf("%#v\n",subscriber) //제로값 출력 why 초기화 X

		subscriber.name ="kim"
		subscriber.rate = 5000
		subscriber.active = false
		fmt.Printf("%s\n",subscriber.name)
		fmt.Println(subscriber.rate)
		fmt.Println(subscriber.active)!/

	//2. 그래서  map을 사용해서 해도 타입이 달라 에러
	subscriber:= map[string]bool{}
	subscriber["name"] ="kim"
	subscriber["rate"] = 5000
	subscriber["active"] =false

		//1.시작 타입이 다 달라서 에러
	   subscriber:= []string{}
	   subscriber =append(subscriber,"kim" )
	   subscriber =append(subscriber, 5000 )
	   subscriber =append(subscriber,false )
	   fmt.Println("test")

}*/
