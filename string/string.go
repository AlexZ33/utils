package string

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
