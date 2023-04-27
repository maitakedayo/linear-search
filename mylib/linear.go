//E:線形探索(先頭から馬鹿正直に)

package mylib

func LinearSearch(arr []int, target int) int {
	/*
		線形探索を行う関数

		Args:
			arr (slice): 探索対象のスライス
			target (int): 探索する値

		Returns:
			int: ターゲットが見つかった場合はインデックス, 見つからない場合は-1
	*/
	for i := 0; i < len(arr); i++ {
		if arr[i] == target {
			return i // ターゲットが見つかった場合はインデックスを返す
		}
	}
	return -1 // ターゲットが見つからない場合は-1を返す
}
