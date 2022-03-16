package main
import "fmt"
//if
//条件分岐
func main(){
	a := 1
	if a == 2{
		fmt.Println("two")
	}else if a == 1{
		fmt.Println("one")
	}else{
		fmt.Println("I don't Know")
	}

	if b := 100; b == 100{
		fmt.Println("one hundred")
	}
	// if文の外と中で参照される変数が異なる
	x := 0
	if x:= 2; true{
		fmt.Println(x)
	}
	fmt.Println(x)
}