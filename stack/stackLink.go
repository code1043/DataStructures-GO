package stack

import (
	"errors"
	"fmt"
)

/**
 * 栈的链式存储结构实现
 */

type node struct {
	elem ElemType //元素
	next *node    //指针域
}

type StackLink struct {
	top    *node //栈顶
	length int   //栈元素个数
}

/**
 * 创建链式栈
 */
func NewStackLink() *StackLink {

	s := StackLink{top: &node{elem: 1043, next: nil}, length: 0}
	return &s
}

/**
 * 入栈
 * 判断栈是否满
 */
func (s *StackLink) Push(e ElemType) (int, error) {

	s.top.next = &node{elem: e, next: s.top.next}
	s.length += 1
	return 0, nil

}

/**
 * 出栈
 * 判断栈是否为空
 */
func (s *StackLink) Pop(e *ElemType) (int, error) {

	if s.IsEmpty() {
		return 1, errors.New("错误信息: 空栈")
	}

	*e = s.top.next.elem
	s.top.next = s.top.next.next
	s.length -= 1
	return 0, nil

}

/**
 * 栈清空
 */
func (s *StackLink) Clear() {

	var e ElemType
	fmt.Println("----------开始清空-----------")
	for !s.IsEmpty() {
		s.Pop(&e)
		fmt.Println("clear", e)
	}
	fmt.Println("----------清空结束-----------")
}

/**
 * 判断栈是否为空
 */
func (s *StackLink) IsEmpty() bool {
	return s.top.next == nil
}

/**
 * 获取栈里的元素个数
 */
func (s *StackLink) Length() int {
	return s.length
}

/**
 * 打印栈的信息
 */
func (s *StackLink) Print() {
	fmt.Println("--------开始打印栈信息--------------")
	fmt.Println("stack info:", s)
	//获取首元节点
	node := s.top.next

	for node != nil {
		fmt.Printf("%v ", node.elem)
		node = node.next
	}
	fmt.Println()
	fmt.Println("--------结束打印栈信息--------------")
}
