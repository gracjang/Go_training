package main

import (
	"fmt"
)

func main() {
	fmt.Printf(
		"TwoSum: %v\n",
		twoSum([]int{2, 7, 11, 15}, 9),
	)

	fmt.Printf(
		"Add two numbers: %v\n",
		addTwoNumbers(&ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: nil}}}, &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4, Next: nil}}}),
	)
}

func twoSum(nums []int, target int) []int {
	for i, num := range nums {
		for j, num2 := range nums {
			if i == j {
				continue
			}
			if num+num2 == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	l3 := &ListNode{}
	if l1 == nil && l2 == nil {
		return nil
	}

	if l1 == nil {
		return &ListNode{l2.Val, addTwoNumbers(nil, l2.Next)}
	}

	if l2 == nil {
		return &ListNode{l1.Val, addTwoNumbers(l1.Next, nil)}
	}

	value := l1.Val + l2.Val
	if value >= 10 {
		value = value - 10
		l3.Val = value
		l3.Next = addTwoNumbers(&ListNode{Val: 1, Next: nil}, addTwoNumbers(l1.Next, l2.Next))
		return l3
	}

	return &ListNode{value, addTwoNumbers(l1.Next, l2.Next)}
}

func lengthOfLongestSubstring(s string) int {
	charMaps := make(map[int]string)

	output := 0
	for word, character := range s {
		
		charMaps[output] = string(character)
		if(charMaps)
	}
}
