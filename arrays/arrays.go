/**
* 工具类 - array,slice, map 的工具方法
**/

package arrays

// RemoveDuplicate 删除[]string 中的重复元素
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

// 遍历数组

// ReverseStringSlice 反转字符串切片
// [site user info 0 ] -> [0 info user site]
func ReverseStringSlice(slice []string) []string {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
	return slice
}

// 从数组中删除 nil 值 Removes nil values from an array
// Example:
//	var arr = []interface{}{1, 2, 3, 4, nil, 5}
//	result := Compact(arr)  // [1, 2, 3, 4, 5]

func Compact(arr []interface{}) []interface{} {
	if arr == nil {
		return nil
	}

	result := make([]interface{}, 0, len(arr))
	for _, v := range arr {
		if v != nil {
			result = append(result, v)
		}
	}
	return result
}
