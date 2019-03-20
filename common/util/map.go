package util

func MergeMap(first map[string]string, second map[string]string) map[string]string {
	result := make(map[string]string)

	if first != nil {
		for key, value := range first {
			result[key] = value
		}
	}
	if second != nil {
		for key, value := range second {
			result[key] = value
		}
	}

	return result
}
