package string

import (
	"strconv"
	"strings"
	"unicode"
)

// IsPalindrome 判断字符串是否是回文(reports whether s reads the same forward and backward)
func IsPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}
func IsPalindrome2(s string) bool {
	if len(s) == 0 {
		return true
	}
	for i := range s {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

// IsPalindrome3 判断字符串是否是回文(reports whether s reads the same forward and backward)
// 忽略大小写和非字母 Letter case is ignored, as are non-letters.
func IsPalindrome3(s string) bool {
	var letters []rune
	for _, r := range s {
		if unicode.IsLetter(r) {
			letters = append(letters, unicode.ToLower(r))
		}
	}
	for i := range letters {
		if letters[i] != letters[len(letters)-1-i] {
			return false
		}
	}
	return true
}

// LongestPalindromes Longest Palindromic Substring 最长回文子串
// Given a string s, find the longest palindromic substring in s.
// You may assume that the maximum length of s is 1000.
// Example 1:
// Input: "babad"
// Output: "bab"
// Note: "aba" is also a valid answer.
// Example 2:
// Input: "cbbd"
// Output: "bb"
func LongestPalindromes(s string) string {
	ll := len(s)
	if ll == 0 {
		return ""
	}
	var l, r, pl, pr int
	for r < ll {
		// gobble up dup chars
		for r+1 < ll && s[r] == s[r+1] {
			r++
		}
		// find size of this palindrome
		for l-1 >= 0 && r+1 < ll && s[l-1] == s[r+1] {
			l--
			r++
		}

		if r-l > pr-pl {
			pl, pr = l, r
		}
		// rest to next mid point
		l = (l+r)/2 + 1
		r = l
	}
	return s[pl : pr+1]

}

func ToInt64(str string) int64 {
	return ToInt64ByDefault(str, 0)
}

// ToInt64ByDefault 将字符串转换为int64
// str: 字符串
// def: 转化失败时候使用默认值
func ToInt64ByDefault(str string, def int64) int64 {
	val, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		val = def
	}
	return val
}

// ToInt str to int, if error return 0
// str: 字符串
func ToInt(str string) int {
	return ToIntByDefault(str, 0)
}

// ToIntByDefault 将字符串转换为int
// str: 字符串
// def: 转化失败时候使用默认值
func ToIntByDefault(str string, def int) int {
	// Atoi is equivalent to ParseInt(s, 10, 0), converted to type int.
	val, err := strconv.Atoi(str)
	if err != nil {
		val = def
	}
	return val
}

/*
IsBlank checks if a string is whitespace or empty (""). Observe the following behavior:
    goutils.IsBlank("")        = true
    goutils.IsBlank(" ")       = true
    goutils.IsBlank("bob")     = false
    goutils.IsBlank("  bob  ") = false
Parameter:
    str - the string to check
Returns:
    true - if the string is whitespace or empty ("")
*/
func IsBlank(str string) bool {
	strLen := len(str)
	if str == "" || strLen == 0 {
		return true
	}
	for i := 0; i < strLen; i++ {
		if unicode.IsSpace(rune(str[i])) == false {
			return false
		}
	}
	return true
}

func IsNotBlank(str string) bool {
	return !IsBlank(str)
}

func IsAnyBlank(strs ...string) bool {
	for _, str := range strs {
		if IsBlank(str) {
			return true
		}
	}
	return false
}

func DefaultIfBlank(str, def string) string {
	if IsBlank(str) {
		return def
	} else {
		return str
	}
}

// IsEmpty checks if a string is empty (""). Returns true if empty, and false otherwise.
func IsEmpty(str string) bool {
	return len(str) == 0
}

// IsNotEmpty checks if a string is not empty (""). Returns false if not empty, and false otherwise
func IsNotEmpty(str string) bool {
	return !IsEmpty(str)
}

// SubStr 截取字符串
// 截取下标[start,end]的字符串
func SubStr(s string, start, length int) string {
	// rune is an alias for int32 and is equivalent to int32 in all ways. It is used, by convention, to distinguish character values from integer values.
	bt := []rune(s)
	if start < 0 {
		start = 0
	}
	if start > len(bt) {
		start = start % len(bt)
	}
	var end int
	if (start + length) > (len(bt) - 1) {
		end = len(bt)
	} else {
		end = start + length
	}
	return string(bt[start:end])
}

// SubStrLen
// 截取下标start开始length长度的字符串
func SubStrLen(s string, start, length int) string {
	return SubStr(s, start, start+length)
}

// Equals 比较两个字符串是否相等
func Equals(a, b string) bool {
	return a == b
}

// EqualsIgnoreCase 比较两个字符串是否相等，忽略大小写
func EqualsIgnoreCase(a, b string) bool {
	return a == b || strings.ToUpper(a) == strings.ToUpper(b)
}

// RuneLen 字符成长度
func RuneLen(s string) int {
	bt := []rune(s)
	return len(bt)
}

// RuneAt 字符串第n个字符
func RuneAt(s string, n int) rune {
	bt := []rune(s)
	return bt[n]
}

// RuneLastAt 字符串最后一个字符
func RuneLastAt(s string) rune {
	bt := []rune(s)
	return bt[len(bt)-1]
}

// RuneFirstAt 字符串第一个字符
func RuneFirstAt(s string) rune {
	bt := []rune(s)
	return bt[0]
}

// ToUnderline 驼峰转下划线
func ToUnderline(str string) string {
	return strings.ToLower(strings.Replace(str, " ", "_", -1))
}

// ToCamel 下划线转驼峰
func ToCamel(str string) string {
	return strings.ToUpper(strings.Replace(str, "_", " ", -1))
}

// ToCamelLower 下划线转驼峰，首字母小写
func ToCamelLower(str string) string {
	return strings.ToLower(strings.Replace(str, "_", " ", -1))
}

// ToCamelUpper 下划线转驼峰，首字母大写
func ToCamelUpper(str string) string {
	return strings.ToUpper(strings.Replace(str, "_", " ", -1))
}
