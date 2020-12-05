package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main(){
	resp, e := http.Get("https://www.inhatc.ac.kr")
	if e != nil{
		log.Fatal(e)
	}
	defer resp.Body.Close()//메모리의 누수 현상이 발생하여 메인함수가 끝나는 시점에는 에러가 있든 없든 프로그램연결 해제하는 작업
	body ,e := ioutil.ReadAll(resp.Body)
	if e !=nil{
		log.Fatal(e)
	}
	//fmt.Println(body) //요렇게 가져오면 슬라이스로 담겨 오는데 아스키 코드 값만 나오므로 글자 형태로 바꾸기 위해
	//fmt.Println(string(body)) string을 쓰면 html같은 소스 코드가 나옴
	fmt.Println(len(body))
}
