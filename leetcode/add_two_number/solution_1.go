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
        var (
            tempVal1 int
            tempVal2 int
        )

        if tempNode1 != nil {
           tempVal1 = tempNode1.Val
        }

        if tempNode2 != nil {
           tempVal2 = tempNode2.Val
        }

        // multiply the value
        tempAdd := tempVal1 + tempVal2
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
        if tempNode1.Next == nil && tempNode2.Next == nil {
            break;
        }

        // assign the next to temp val
        if tempNode1.Next != nil {
            tempNode1 = tempNode1.Next
        }
        
        if tempNode2.Next != nil {
            tempNode2 = tempNode2.Next
        }      
    }

    // create the result
    var (
        result *ListNode
    )
    for _, val := range resultArr {
        fmt.Println(val)
        tempNode := &ListNode{
            Val: val,
            Next: nil,
        }

        if result != nil {
            result.Next = tempNode
        }
        
        result = tempNode
        fmt.Println(result)
    }

    return result
}
