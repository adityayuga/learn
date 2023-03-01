/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    isOverflow := false
    resultArr := []int{}

    tempNode1 := l1
    tempNode2 := l2

    for {
        // multiply the value
        tempAdd := tempNode1.Val + tempNode2.Val
        if isOverflow {
             tempAdd += 1
        }
        
        // append the result with the modulo
        // the probability is the multiplier can be more than 10 only
        resultArr = append(resultArr, tempAdd % 10)

        // is overflow then mark is, if not then reset
        if tempAdd >= 10 {
            isOverflow = true
        } else {
            isOverflow = false            
        }

        // if no more next value
        if tempNode1.Next == nil || tempNode2.Next == nil {
            break;
        }
        
        // assign the next to temp val
        tempNode1 = tempNode1.Next
        tempNode2 = tempNode2.Next        
    }

    // create the result
    var (
        result *ListNode
    )
    for _, val := range resultArr {
        tempNode := &ListNode{
            Val: val,
            Next: nil,
        }

        if result != nil {
            result.Next = tempNode
        }
        
        result = tempNode
    }

    return result
}
