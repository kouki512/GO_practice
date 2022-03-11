// 変数
package main
// goは定義した変数を全て使う必要がある。
import (
    "strconv"
    "fmt"
)

func main(){
 var i int = 100
 fmt.Println(i)

 var s string = "Hello Go"
 fmt.Println(s)

 var t,f bool = true,false
 fmt.Println(t,f)

 var i3 int 
 var s3 string
 fmt.Println(i3,s3)

 i3 = 300
 s3 = "go"
 fmt.Println(i3,s3)

 // 暗黙的な定義
 // 変数名 := 値
 // 関数内でしか定義できない
 i4 := 400
 fmt.Println(i4)
//  i4 = "Hello"
//  fmt.Println(i4)

// 配列型
// 後から変更できない
// 要素数が3のint型の配列を宣言
 var arr1 [3]int 
 fmt.Println(arr1)
 fmt.Printf("%T\n",arr1)

 var arr2 [3]string = [3]string{"A","B"}
 fmt.Println(arr2)

 // 暗黙的な配列の定義
 arr3 := [3]int{1,2,3}
 fmt.Println(arr3)
 
 arr4 := [...]string{"C"}
 fmt.Println(arr4)
 
 fmt.Println(arr2[0])
 // 行番号の計算
 fmt.Println(arr2[2-1])

 arr2[2] = "C"
 fmt.Println(arr2)

 //var arr5 [4]int
 //arr5 = arr1
 // 3[int]に4[int]を入れようとしているためエラー
 // 要素数が異なると別の変数として扱われる
 //fmt.Println(arr5)

//interface & nil

//あらゆる型と互換性がある
// データ特有の演算は出来ない
	var x interface{}
	fmt.Println(x)
	x = 1
	fmt.Println(x)
	x = 3.14
	fmt.Println(x)
	x = "aaa"
	fmt.Println(x)
	x = [3]int{1,2,3}
	fmt.Println(x)

// 型変換
 var a int = 1
 fl64 := float64(i)
 fmt.Printf("a = %T\n",a)
 fmt.Printf("a = %T\n",fl64)

 i2 := int(fl64)
 fmt.Printf("i2 = %T\n",i2)

var d string = "100"
 b, _ := strconv.Atoi(d)
 fmt.Println(b)
}
