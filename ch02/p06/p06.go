package p06

import (
	"ctci/container/list"
	"ctci/container/stack"
)

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
	fast, slow := l.Front(), l.Front()
	stack := stack.New()
	for ; !l.IsConsumed(fast) && !l.IsConsumed(fast.Next); fast = fast.Next.Next {
		stack.Push(slow.Value)
		slow = slow.Next
	}
	if l.IsConsumed(fast.Next) {
		slow = slow.Next
	}
	for ; !l.IsConsumed(slow); slow = slow.Next {
		if slow.Value != stack.Pop() {
			return false
		}
	}
	return true
}

// IsPalindrome3 checks if a list l is a palindrome
func IsPalindrome3(l *list.List) bool {
	ret := isPalindromeRecurse(l.Front(), l.Len())
	return ret.isPalindrome
}

func isPalindromeRecurse(e *list.Element, len int) *result {
	switch len {
	case 0: //even
		return &result{true, e}
	case 1: //odd
		return &result{true, e.Next}
	}

	ret := isPalindromeRecurse(e.Next, len-2)
	if !ret.isPalindrome {
		return ret
	}
	if e.Value != ret.e.Value {
		ret.isPalindrome = false
	}
	ret.e = ret.e.Next
	return ret
}

type result struct {
	isPalindrome bool
	e            *list.Element
}
