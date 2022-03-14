package main
import "fmt"
//関数
//クロージャ―の実装
func Later() func(string) string {
	var store string
	return func (next string) string {
		s := store
		store = next 
		return s
	}
}
func main(){
	f := Later()
	fmt.Println(f("Hello"))
	fmt.Println(f("My"))
	fmt.Println(f("name"))
	fmt.Println(f("is"))
}