package main

import (
	"dataStructure/liner"
	"dataStructure/queue"
	"dataStructure/stack"
	"fmt"
)

func main() {

	// exampleSqList()
	// exampleLinkList()
	// exampleLinkDList()
	// exampleStackSq()
	// exampleStackLink()
	exampleQueueSq()

}

/**
 * 队列的顺序存储结构案例
 */
func exampleQueueSq() {

	q := queue.NewQueueSq(10)
	q.Print()
	fmt.Println("---------开始入队-----------")
	for i := 0; i < 10; i++ {
		status, err := q.EnQueueSq(queue.ElemType(i + 1))
		fmt.Println("error: ", status, err)
	}
	fmt.Println("---------结束入队-----------")
	q.Print()

	fmt.Println("---------开始出队-----------")
	var e queue.ElemType
	for i := 0; i < 5; i++ {
		status, err := q.DeQueueSq(&e)
		fmt.Println("error: ", status, err, e)
	}
	fmt.Println("---------结束出队-----------")
	q.Print()
	fmt.Println("---------开始入队-----------")
	for i := 0; i < 10; i++ {
		status, err := q.EnQueueSq(queue.ElemType(i + 1))
		fmt.Println("error: ", status, err)
	}
	fmt.Println("---------结束入队-----------")
	q.Print()

}

/**
 * 栈的顺序存储结构案例
 */
func exampleStackLink() {
	s := stack.NewStackLink()

	for i := 0; i < 10; i++ {
		s.Push(stack.ElemType(i + 1))
	}
	s.Print()

	s.Clear()
	s.Print()

	var e stack.ElemType
	status, err := s.Pop(&e)
	fmt.Println("error :", status, err, e)
	s.Print()

}

/**
 * 栈的顺序存储结构案例
 */
func exampleStackSq() {

	s := stack.NewStackSq(10)
	s.Print()
	for i := 0; i < 10; i++ {
		status, err := s.Push(stack.ElemType(i + 1))
		fmt.Println("error :", status, err)
	}

	status, err := s.Push(stack.ElemType(1043))
	fmt.Println("error :", status, err)
	s.Print()

	//清空栈
	s.Clear()
	s.Print()

	var e stack.ElemType
	for i := 0; i < 10; i++ {
		status, err := s.Pop(&e)
		fmt.Println("error :", status, err)
		fmt.Println(e)
	}

	status, err = s.Pop(&e)
	fmt.Println("error :", status, err)
	s.Print()

}

/**
 * 链式存储结构案例 - 双向链表
 */
func exampleLinkDList() {

	//创建链表
	l := liner.NewLinkDList()
	fmt.Println(l.Length())

	fmt.Println("---------start----------")
	for i := 1; i <= 10; i++ {
		l.Insert(liner.ElemType(i), i)
	}
	l.Insert(liner.ElemType(100), 1)
	l.Insert(liner.ElemType(120), 2)
	l.Insert(liner.ElemType(130), 1)
	l.Insert(liner.ElemType(1043), 7)
	l.Print()

	// //根据位置获取元素
	// var e liner.ElemType
	// l.GetAtPosition(10, &e)
	// fmt.Println(e)

	// //根据元素获取位置
	// var i int
	// l.GetAtElem(10, &i)
	// fmt.Println(i)

	l.Delete(7)
	l.Print()

	// for i := 0; i < 12; i++ {
	// 	l.Delete(1)
	// }
	// fmt.Println("删除后的链表")
	// l.Print()

	fmt.Println("-----------end-----------")

	// // //销毁链表
	liner.DestroyLinkDList(&l)
	if l != nil {
		l.Print()
	}

}

/**
 * 链式存储结构案例
 */
func exampleLinkList() {

	//创建链表
	l := liner.NewLinkList()
	fmt.Println(l.IsEmpty(), l.Length())

	fmt.Println("---------start----------")
	for i := 1; i <= 10; i++ {
		l.Insert(liner.ElemType(i), i)
	}
	l.Insert(liner.ElemType(100), 1)
	l.Insert(liner.ElemType(120), 2)
	l.Insert(liner.ElemType(130), 1)
	l.Print()

	//根据位置获取元素
	var e liner.ElemType
	l.GetAtPosition(10, &e)
	fmt.Println(e)

	//根据元素获取位置
	var i int
	l.GetAtElem(10, &i)
	fmt.Println(i)

	l.Delete(1)
	l.Print()

	for i := 0; i < 12; i++ {
		l.Delete(1)
	}
	fmt.Println("删除后的链表")
	l.Print()

	fmt.Println("-----------end-----------")

	// // //销毁链表
	liner.DestroyLinkList(&l)
	if l != nil {
		l.Print()
	}

}

/**
 * 顺序存储结构案例
 */
func exampleSqList() {

	s := liner.NewSqList()
	fmt.Println(s)

	liner.DestroySqList(&s)
	fmt.Println(s)

	// for i := 1; i <= 10; i++ {
	// 	s.Insert(liner.ElemType(i), i)

	// }

	// fmt.Println(s)

	// _, err := s.Insert(liner.ElemType(100), 5)
	// fmt.Println(s, err)

	// _, err = s.Insert(liner.ElemType(100), 12)
	// fmt.Println(s, err)

	// _, err = s.Insert(liner.ElemType(100), 15)
	// fmt.Println(s, err)

	// //以下有一个bug
	// p, err := s.Get(12)
	// fmt.Println(p, err)

	// _, err = s.DeleteAtElem(p)
	// fmt.Println(s, err)

	// _, err = s.Delete(11)
	// fmt.Println(s, err)

	// _, err = s.Delete(12)
	// fmt.Println(s, err)
}
