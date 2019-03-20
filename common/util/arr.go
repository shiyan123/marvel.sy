package util

import "github.com/cznic/mathutil"

func SplitArrayByCount(arr []string, count int) (result [][]string) {
	for i := 0; i < len(arr); i += count {
		result = append(result, arr[i:mathutil.Min(i+count, len(arr))])
	}
	return
}
