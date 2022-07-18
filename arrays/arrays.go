package arrays

/**
删除[]string 中的重复元素
*/

func RemoveDuplicate(slice []string) []string {
	m := make(map[string]int)
	for _, v := range slice {
		m[v] = 0
	}
	r := make([]string, 0)
	for k := range m {
		r = append(r, k)
	}

	return r
}