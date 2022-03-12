package main

import "fmt"

//定数

//大文字にすることで他のパッケージから使用できる
const Pi = 3.14
const (
	URL = "http://aaa.co.jp"
	SiteName = "test"
)
const (

	c0 = iota
	c1
	c2
)

func main(){
	fmt.Println(Pi)
	fmt.Println	(URL,SiteName)
	fmt.Println(c0,c1,c2)
}