/**
* @Description 工具类 - array,slice, map 的工具方法
* @Author: AlexZ33
* @Date: 2022-08-01 23:00:00
**/

package arrays

import (
	"bytes"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"sync"
	"unsafe"
)

// CopyShallowMap makes a shallow copy of a map
func CopyShallowMap(m map[string]string) map[string]string {
	if m == nil {
		return nil
	}
	copy := make(map[string]string, len(m))
	for key, value := range m {
		copy[key] = value
	}
	return copy
}

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

// Contains 判断某个元素是否在slice, array, map中
func Contains(search interface{}, target interface{}) bool {
	targetValue := reflect.ValueOf(target)
	switch reflect.TypeOf(target).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < targetValue.Len(); i++ {
			if targetValue.Index(i).Interface() == search {
				return true
			}
		}
	case reflect.Map:
		if targetValue.MapIndex(reflect.ValueOf(search)).IsValid() {
			return true
		}
	}

	return false
}

// ContainsIgnoreCase check if a string (search) is present in a slice of strings (target) regardless of the case.
//words := []string{"Apple", "Banana", "Cherry", "Date", "Fig", "Grape"}
//
//	fmt.Println(ContainsIgnoreCase("apple", words)) // true
//	fmt.Println(ContainsIgnoreCase("BANANA", words)) // true
//	fmt.Println(ContainsIgnoreCase("Cherry", words)) // true
//	fmt.Println(ContainsIgnoreCase("date", words)) // true
//	fmt.Println(ContainsIgnoreCase("fig", words)) // true
//	fmt.Println(ContainsIgnoreCase("grape", words)) // true
//	fmt.Println(ContainsIgnoreCase("kiwi", words)) // false
func ContainsIgnoreCase(search string, target []string) bool {
	if len(search) == 0 {
		return false
	}
	if len(target) == 0 {
		return false
	}
	search = strings.ToLower(search)
	for i := 0; i < len(target); i++ {
		if strings.ToLower(target[i]) == search {
			return true
		}
	}
	return false
}

// Equal 通过reflect.DeepEqual比较两个slice、struct、map是否相等
// 来自reflect.DeepEqual函数能够对两个值进行深度相等判断
// 注意： 它会对一个nil值map和非nil值但是空的map视作不相等
// 同样nil值slice和非nil值但是空的slice视作不相等
// 自己实现一个equal函数以解决上面问题
// Reference: 《GO语言圣经: 深度相等判断rrors.New("Not in array")》
// https://github.com/adonovan/gopl.io/blob/master/ch13/equal/equal.go
// Equal provides a deep equivalence relation for arbitrary values

type comparison struct {
	a, b unsafe.Pointer
	t    reflect.Type
}

func equal(a, b reflect.Value, seen map[comparison]bool) bool {
	if !a.IsValid() || !b.IsValid() {
		return a.IsValid() == b.IsValid()
	}
	if a.Type() != b.Type() {
		return false
	}
	if a.CanAddr() && b.CanAddr() {
		aptr := unsafe.Pointer(a.UnsafeAddr())
		bptr := unsafe.Pointer(b.UnsafeAddr())
		if aptr == bptr {
			return true // identical references
		}
		c := comparison{aptr, bptr, a.Type()}
		if seen[c] {
			return true // already seen
		}
		seen[c] = true
	}

	switch a.Kind() {
	case reflect.Bool:
		return a.Bool() == b.Bool()
	case reflect.String:
		return a.String() == b.String()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return a.Int() == b.Int()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return a.Uint() == b.Uint()
	case reflect.Float32, reflect.Float64:
		return a.Float() == b.Float()
	case reflect.Complex64, reflect.Complex128:
		return a.Complex() == b.Complex()
	case reflect.Chan, reflect.UnsafePointer, reflect.Func:
		return a.Pointer() == b.Pointer()
	case reflect.Ptr, reflect.Interface:
		return equal(a.Elem(), b.Elem(), seen)
	case reflect.Array, reflect.Slice:
		if a.Len() != b.Len() {
			return false
		}
		for i := 0; i < a.Len(); i++ {
			if !equal(a.Index(i), b.Index(i), seen) {
				return false
			}
		}
		return true
	case reflect.Struct:
		for i, n := 0, a.NumField(); i < n; i++ {
			if !equal(a.Field(i), b.Field(i), seen) {
				return false
			}
		}
		return true
	case reflect.Map:
		if a.Len() != b.Len() {
			return false
		}
		for _, k := range a.MapKeys() {
			if !equal(a.MapIndex(k), b.MapIndex(k), seen) {
				return false
			}
		}
		return true
	}
	panic("unreachable")
}

// Equal reports whether a and b are deeply equal.
// Map keys are always compared with ==, not deeply.
// (This matters for keys containing pointers or interfaces)
func Equal(a, b interface{}) bool {
	seen := make(map[comparison]bool)
	return equal(reflect.ValueOf(a), reflect.ValueOf(b), seen)
}

/**
 * @name: IntArrayEqual
 * @descripttion: 判断IntArray是否相等
 * @param {[]int} arr1
 * @param {[]int} arr2
 * @return {bool}
 */
func StringArrayEqual(arr1, arr2 []int) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	sort.Ints(arr1)
	sort.Ints(arr2)
	for i, arr := range arr1 {
		if arr != arr2[i] {
			return false
		}
	}
	return true
}

/**
 * @name: InStringArray
 * @descripttion: 是否在字符串列表中
 * @param {string} target
 * @param {[]string} strArray
 * @return {bool}
 */
func InStringArray(target string, strArray []string) bool {
	sort.Strings(strArray)
	index := sort.SearchStrings(strArray, target)
	if index < len(strArray) && strArray[index] == target {
		return true
	}
	return false
}

/**
 * @name: InInt64Array
 * @descripttion: 是否在Int64列表中
 * @param {string} target
 * @param {[]string} strArray
 * @return {bool}
 */
func InInt64Array(target int64, int64Array []int64) bool {
	targetStr := strconv.FormatInt(target, 10)
	var strArray []string
	for _, s := range int64Array {
		strArray = append(strArray, strconv.FormatInt(s, 10))
	}
	return InStringArray(targetStr, strArray)
}

/**
 * @name: InIntArray
 * @descripttion: 是否在Int列表中
 * @param {string} target
 * @param {[]string} strArray
 * @return {bool}
 */
func InIntArray(target int, intArray []int) bool {
	sort.Ints(intArray)
	index := sort.SearchInts(intArray, target)
	if index < len(intArray) && intArray[index] == target {
		return true
	}
	return false
}

//ToSlice 转换为数组
func ToSlice(arr interface{}) []interface{} {
	v := reflect.ValueOf(arr)
	if v.Kind() != reflect.Slice {
		panic("toslice() given a non-slice type")
	}
	l := v.Len()
	ret := make([]interface{}, l)
	for i := 0; i < l; i++ {
		ret[i] = v.Index(i).Interface()
	}
	return ret
}

// UniqueString returns a unique string from a slice
func UniqueString(stringSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range stringSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

// Difference returns the difference  between two string slices
func Difference(a, b []string) (diff []string) {
	m := make(map[string]bool)

	for _, item := range b {
		m[item] = true
	}

	for _, item := range a {
		if _, ok := m[item]; !ok {
			diff = append(diff, item)
		}
	}

	return
}

var (
	bufferPool = sync.Pool{
		New: func() interface{} {
			return bytes.NewBuffer([]byte{})
		},
	}
)

// JoinInts format int64 to slice like: n1, n2, n3
func JoinInts(list []int64) string {
	if len(list) == 0 {
		return ""
	}

	if len(list) == 1 {
		return strconv.FormatInt(list[0], 10)
	}

	buf := bufferPool.Get().(*bytes.Buffer)
	for _, i := range list {
		buf.WriteString(strconv.FormatInt(i, 10))
		buf.WriteByte(',')
	}
	if buf.Len() > 0 {
		buf.Truncate(buf.Len() - 1)
	}
	s := buf.String()
	buf.Reset()
	bufferPool.Put(buf)
	return s
}

// SpliteInts split string into int64 slice.
func SpliteInts(s string) ([]int64, error) {
	if s == "" {
		return nil, nil
	}
	sArr := strings.Split(s, ",")
	res := make([]int64, 0, len(sArr))
	for _, sc := range sArr {
		i, err := strconv.ParseInt(sc, 10, 64)
		if err != nil {
			return nil, err
		}
		res = append(res, i)
	}
	return res, nil
}
