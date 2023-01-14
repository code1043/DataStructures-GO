package liner

import (
	"errors"
	"fmt"
)

// 声明节点的类型
type Lnode struct {
	data ElemType // 数据域
	next *Lnode   // 指针域
}

// 声明链表的类型
type LinkList struct {
	length int    // 链表的长度
	head   *Lnode // 指向头节点
}

/**
 * 创建一个单链表
 * 长度初始化未1
 * head初始化未节点
 */
func NewLinkList() *LinkList {
	return &LinkList{length: 0, head: &Lnode{data: 1043}}
}

/**
 * 销毁指针
 */
func DestroyLinkList(l **LinkList) {

	/**
	 * 链表：头指针、头节点、首元节点
	 */
	//调用链表清空方法 清空所有节点
	(*l).Clear()

	//头节点销毁
	(*l).head = nil

	//链表销毁
	(*l) = nil
}

/**
 * 判断链表是否为空
 * 判断逻辑，头节点的指针域是否为nil
 * 若是则链表为空
 * 否则链表不为空
 */
func (l *LinkList) IsEmpty() bool {
	return l.head.next == nil
}

/**
 * 清空列表
 * 即保留头指针和头节点
 */
func (l *LinkList) Clear() {
	/**
	 * 链表：头指针、头节点
	 */

	// 指向头节点
	p := l.head
	//循环链表
	for p.next != nil {

		//创建一个变量指向头节点指向的节点
		node := p.next

		fmt.Println("删除的节点+", node.data)

		// 更新头节点指向下一个节点
		p.next = node.next

		//节点销毁
		node = nil
	}

}

/**
 * 求列表的长度
 * 实现逻辑，循环节点+ 1
 */

func (l *LinkList) Length() int {

	// //累加变量
	// count := 0
	// // 指向首元节点
	// node := l.head.next
	// //循环链表
	// for node != nil {
	// 	//节点不为空 累加
	// 	count += 1
	// 	// 更新节点指向下一个节点
	// 	node = node.next

	// }

	return l.length
}

/**
 * 取值，取单链表中第i个元素
 */
func (l *LinkList) GetAtPosition(i int, e *ElemType) (int, error) {

	//判断位置是否合法
	if i < 1 || i > l.Length() {
		return 0, errors.New("超出线性表范围")
	}

	//累加器
	count := 0

	//获取首元节点
	node := l.head.next

	for node != nil {
		count += 1
		if count == i {
			break
		} else {
			node = node.next
		}

	}

	if node == nil {
		return 0, errors.New("未找到元素")
	} else {
		*e = node.data
		return 1, nil
	}

}

/**
 * 按值查找
 */

func (l *LinkList) GetAtElem(e ElemType, i *int) (int, error) {

	//累加器
	count := 0

	//获取首元节点
	node := l.head.next
	for node != nil {
		count += 1
		if node.data == e {
			break
		} else {
			node = node.next
		}

	}

	if node == nil {
		return 0, errors.New("未找到元素")
	} else {
		*i = count
		return 1, nil
	}

}

func (l *LinkList) Insert(e ElemType, i int) (int, error) {

	//判断位置是否合法
	if i < 1 || i > l.Length()+1 {
		return 0, errors.New("超出线性表范围")
	}

	//累加器
	count := 0

	//获取头节点
	node := l.head

	for node != nil {
		count += 1
		if count == i {
			break
		} else {
			node = node.next
		}

	}

	if node == nil {
		return 0, errors.New("未找到元素")
	} else {
		node.next = &Lnode{e, node.next}
		l.length += 1
		return 1, nil
	}

}

/**
 * 删除第i个节点的数据
 */
func (l *LinkList) Delete(i int) (int, error) {

	//判断位置是否合法
	if i < 1 || i > l.Length() {
		return 0, errors.New("超出线性表范围")
	}

	//累加器
	count := 0

	//获取首元节点
	node := l.head

	for node != nil {
		count += 1
		if count == i {
			break
		} else {
			node = node.next
		}

	}

	if node == nil {
		return 0, errors.New("未找到元素")
	} else {

		// n := node.next
		// node.next = n.next
		// n = nil

		node.next = node.next.next

		return 1, nil
	}

}

/**
 * 打印链表中的数据
 */
func (l *LinkList) Print() {

	node := l.head
	fmt.Println("head:", node.data, node.next)
	node = node.next
	for node != nil {
		fmt.Println("node:", node.data, node.next)
		node = node.next
	}

}
