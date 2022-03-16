package main
import "fmt"
import "strconv"
//if
//条件分岐
//エラーハンドリング
func main(){
	var s string = "100"
	i,err := strconv.Atoi(s)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println("i = %T\n",i)
}