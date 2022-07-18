package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func isPalindrome(arr []int) bool {
    // var arr []int;
    var condition bool;
    
    // for head.Next != nil {
    //     arr = append(arr, head.Val)
    //     head = head.Next
    // }
    
    j := len(arr) - 1;
    
    for i := 0; i < len(arr) ; i++ {
        if arr[i] == arr[j]{
            condition = true
        } else {
            return false;
        }
        j--;
    }
    
    return condition
}

func main(){
	fmt.Println(isPalindrome([]int{1,2,1}))
}