package main

import (
	"dataStructure/liner"
	"fmt"
)

func main() {

	// exampleSqList()
	// exampleLinkList()
	exampleLinkDList()

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
