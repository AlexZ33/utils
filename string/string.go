package string

import "unicode"

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
