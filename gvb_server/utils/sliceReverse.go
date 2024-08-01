package utils

//切片倒排
/*
	例如 ["a","b","c","d"],倒排为["d","c","b","a"]
	会修改原切片
*/
func SliceReverse[T any](slice []T) {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}