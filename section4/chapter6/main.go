package main
import "fmt"
//関数
//ジェネレーター
func integers() func()int {
	i := 0
	return func()int {
		i++
		return i
	}
}
func main(){
	ints := integers()
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())

	// 新たな領域が作成される
	otherints := integers()
	fmt.Println(otherints())
	fmt.Println(otherints())
	fmt.Println(otherints())
	fmt.Println(otherints())
}