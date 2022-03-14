package main
import "fmt"
//関数
//無名関数
func main(){
	// 無名関数の定義
	f := func(x,y int) int {
		return x + y
	}
	i := f(1,2)
	fmt.Println(i)

	// 簡易的な関数定義をしたい時
	i2 := func(x,y int) int {
		return (x + y) * 2
	}(1,2)
	fmt.Println(i2)
}