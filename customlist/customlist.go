package customlist

type linkedNodeCallback func(node *LinkedNode)

func New() *CustomList {
    return new(CustomList)
}

/*
 * Count
 * Push
 * Remove
 * IndexOf
 * LastIndexOf
 * ForEach
 * Clear
 */
type CustomList struct {
    root *LinkedNode
    tail *LinkedNode
    count int
}
//
func (customList *CustomList) Count() int {
    return customList.count
}
//
func (customList *CustomList) Push(obj interface{}) {
    if customList.root == nil {
        nextNode := &LinkedNode{value: obj}
        customList.root =  nextNode
        customList.tail = nextNode
        customList.count = 1
    } else if customList.tail != nil {
        nextNode := &LinkedNode{value: obj}
        customList.tail.next = nextNode
        customList.count++
    }
}
//
func (customList *CustomList) ForEach(fn linkedNodeCallback) {
    if customList.root != nil {
        customList.root.ForEach(fn)
    }
}
//
/**/
type LinkedNode struct {
    value interface{}
    next *LinkedNode
}

func (node *LinkedNode) ForEach(fn linkedNodeCallback) {
    var current *LinkedNode
    current = node
    fn(current)
    if current.next != nil {
        current.next.ForEach(fn)
    }
}