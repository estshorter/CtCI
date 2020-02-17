package p06

import "ctci/container/list"

// IsPalindrome checks if a list l is a palindrome
func IsPalindrome(l *list.List) bool {
	ef, eb := l.Front(), l.Back()
	for i := 0; i < l.Len()/2; i++ {
		if ef.Value != eb.Value {
			return false
		}
		ef, eb = ef.Next, eb.Prev
	}
	return true
}

// IsPalindrome2 checks if a list l is a palindrome
func IsPalindrome2(l *list.List) bool {
	slice := make([]int, l.Len())
	i := 0
	for e := l.Front(); !l.IsConsumed(e); e = e.Next {
		slice[i] = e.Value
		i++
	}
	for i := 0; i < l.Len()/2; i++ {
		if slice[i] != slice[len(slice)-1-i] {
			return false
		}
	}
	return true
}
