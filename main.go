//E:線形探索(先頭から馬鹿正直に)

package main

import (
	"fmt"

	"github.com/maitakedayo/linear-search/mylib"
)

func main() {
	arr := []int{3, 9, 2, 7, 1, 5, 6, 8, 4} //初期化
	target := 7
	result := mylib.LinearSearch(arr, target)
	if result != -1 {
		fmt.Println("ターゲットが見つかりました。インデックス: ", result)
	} else {
		fmt.Println("ターゲットが見つかりませんでした。")
	}
}
