package main

import "fmt"

func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}

	hash := map[byte]byte{')': '(', ']': '[', '}': '{'}
	stack := make([]byte, 0)

	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack = append(stack, s[i])
		} else if len(stack) > 0 && stack[len(stack) - 1] == hash[s[i]] {
			stack = stack[:(len(stack) - 1)]
		} else {
			return false
		}
	}

	return len(stack) == 0
}

func main() {
	fmt.Println(isValid("("))
}
