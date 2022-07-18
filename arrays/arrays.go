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

// Flatten 返回一个新的一维平面数组 Returns a new array that is one-dimensional flat.
// Example:
//	var arr1 = []interface{}{1, 2, 3, 4}       // [1, 2, 3, 4]
//	var arr2 = []interface{}{5, 6, 7, arr1}    // [5, 6, 7, [1, 2, 3, 4]]
//	result := arrays.Flatten(arr2)          // [5, 6, 7, 1, 2, 3, 4]
func Flatten(arr []interface{}) []interface{} {
	if arr == nil {
		return arr
	}

	result := make([]interface{}, 0, len(arr))
	for _, v := range arr {
		switch v.(type) {
		case []interface{}:
			f1 := Flatten(v.([]interface{}))
			for _, v1 := range f1 {
				result = append(result, v1)
			}
		default:
			result = append(result, v)

		}
	}
	return result
}

//为数组中的每个元素调用“MapFunc”。创建一个新数组，包含 `MapFunc` 返回的值
// Example:
//	var arr = []interface{}{1, 2, 3, 4}
// func MyMapFunc(v interface{}) interface{} {
//		return v.(int) * 3
//	}
//  result := arrays.Map(arr, MyMapFunc)  // [3, 6, 9, 12]

type MapFunc func(v interface{}) interface{}

func Map(arr []interface{}, f MapFunc) []interface{} {
	if arr == nil || f == nil {
		return arr
	}

	result := make([]interface{}, 0, len(arr))
	for _, v := range arr {
		result = append(result, f(v))
	}
	return result
}
