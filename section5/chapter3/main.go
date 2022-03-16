package main
import "fmt"
//if
//繰り返し処理
func main(){
	// i := 0
	// for {
	// 	i++
	// 	if i == 3{
	// 		break
	// 	}
	// 	fmt.Println("Loop")
	// }

// while文のようなもの
	// point := 0
	// for point < 10{
	// 	fmt.Println(point)
	// 	point++
	// }
// C++のfor文みたいなもの
	// for i := 0; i < 10; i++ {
	// 	if i == 3{
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }

// 配列でfor文を使う
	// arr := [3]int{1,2,3}
	// for i := 0; i < len(arr); i++ {
	// 	fmt.Println(arr[i])
	// }

// 範囲式のfor
    // arr := [3]int{1,2,3}
	// for i,v := range arr{
	// 	fmt.Println(i,v)
	// }
// スライス
	// sl := []string{"Python","PHP","GO"}
	// for i,v := range sl{
	// 	fmt.Println(i,v)
	// }

// マップ
	m := map[string]int{"apple":100,"banana":200}
	for k,v := range m {
		fmt.Println(k,v)
	}
}