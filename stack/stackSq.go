package stack

import (
	"errors"
	"fmt"
)

/**
 * 栈的顺序存储结构实现
 */

type StackSq struct {
	top    int        //栈顶
	base   int        //栈底
	cap    int        //栈的大小
	elem   []ElemType //数据域
	length int        //元素个数
}

func NewStackSq(cap int) *StackSq {
	return &StackSq{top: 0, base: 0, cap: cap, elem: make([]ElemType, size), length: 0}
}

/**
 * 入栈
 * 判断栈是否满
 */
func (s *StackSq) Push(e ElemType) (int, error) {

	if s.IsFull() {
		return 1, errors.New("栈已溢出")
	}

	s.elem[s.top] = e
	s.top += 1
	s.length += 1
	return 0, nil

}

/**
 * 出栈
 * 判断栈是否为空
 */
func (s *StackSq) Pop(e *ElemType) (int, error) {

	if s.IsEmpty() {
		return 1, errors.New("栈已空")
	}

	s.top -= 1
	*e = s.elem[s.top]
	s.elem[s.top] = 0 //这里有问题，因为elem的元素是ElemType，ElemType的值是不定的，有可能是基础类型，也有可能是复合类型，所以不应该是0.
	s.length -= 1
	return 0, nil

}

/**
 * 获取栈顶元素
 */
func (s *StackSq) GetTop(e *ElemType) (int, error) {

	if s.IsEmpty() {
		return 1, errors.New("栈已空")
	}

	*e = s.elem[s.top-1]
	return 0, nil

}

/**
 * 栈清空
 */
func (s *StackSq) Clear() {

	var e ElemType
	fmt.Println("----------开始清空-----------")
	for !s.IsEmpty() {
		s.Pop(&e)
		fmt.Println("clear", e)
	}

	// 清空里的另一种实现
	// for i := 0; i < s.length; i++ {
	// 	fmt.Println("clear", s.elem[i])
	// 	s.elem[i] = 0 //这里有问题，因为elem的元素是ElemType，ElemType的值是不定的，有可能是基础类型，也有可能是复合类型，所以不应该是0.
	// }
	// s.length = 0
	// s.top = 0

	fmt.Println("----------清空结束-----------")
}

/**
 * 判断栈是否溢出
 */
func (s *StackSq) IsFull() bool {
	return s.top == s.size
}

/**
 * 判断栈是否为空
 */
func (s *StackSq) IsEmpty() bool {
	return s.top == s.base
}

/**
 * 获取栈里的元素个数
 */
func (s *StackSq) Length() int {
	return s.length
}

/**
 * 打印栈的信息
 */
func (s *StackSq) Print() {
	fmt.Println("--------开始打印栈信息--------------")
	fmt.Println("stack info:", s)
	fmt.Println("--------结束打印栈信息--------------")
}

/**
 * 高级语言不需要写销毁方法，因为有GC
 */
// func DestroyStackSq() {
// }
