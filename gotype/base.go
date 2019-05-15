package gotype

//链表定义
type LNode struct {
	Data interface{}
	Next *LNode
}
func NewLNode()*LNode{
	return &LNode{}
}