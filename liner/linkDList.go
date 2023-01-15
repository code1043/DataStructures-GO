package liner

import (
	"errors"
	"fmt"
)

/**
 * 线性表的链式存储结构实现 -- 双向链表
 */

type LDnode struct {
	data        ElemType //数据域
	prior, next *LDnode  //指针域
}

// 声明链表的类型
type LinkDList struct {
	length int     // 链表的长度
	link   *LDnode // 数据元素的指针
}

/**
 * 创建一个双向单链表
 * 长度初始化未1
 * head初始化未节点
 */
func NewLinkDList() *LinkDList {
	return &LinkDList{length: 0, link: &LDnode{data: 1043}}
}

/**
 * 销毁指针
 */
func DestroyLinkDList(l **LinkDList) {

	/**
	 * 链表：头指针、头节点、首元节点
	 */
	//调用链表清空方法 清空所有节点
	(*l).Clear()

	//头节点销毁
	(*l).link = nil

	//链表销毁
	(*l) = nil
}

/**
 * 清空列表
 * 即保留头指针和头节点
 */
func (l *LinkDList) Clear() {
	/**
	 * 链表：头指针、头节点
	 */

	// 指向头节点
	p := l.link
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

func (l *LinkDList) Length() int {

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
 * 判断链表是否为空
 * 判断逻辑，头节点的指针域是否为nil
 * 若是则链表为空
 * 否则链表不为空
 */
func (l *LinkDList) IsEmpty() bool {
	return l.link.next == nil
}

func (l *LinkDList) Insert(e ElemType, i int) (int, error) {

	//判断位置是否合法
	if i < 1 || i > l.Length()+1 {
		return 0, errors.New("超出线性表范围")
	}

	//累加器
	count := 0

	//获取头节点
	node := l.link

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
		fmt.Println("开始插入")
		// fmt.Println(node)

		// 创建一个新的节点 newNode
		// 新节点的prior指向 node
		// 新节点的next指向 node.next
		// node.next.prior 指向newNode
		// node.next 指向 newNode

		newNode := &LDnode{data: e, prior: node, next: node.next}
		if node.next != nil {
			node.next.prior = newNode
			fmt.Println("prior", node.next.prior)
		}
		node.next = newNode

		l.length += 1
		fmt.Println("插入结束")

		// node.next.prior = &LDnode{data: e, prior: node, next: node.next}
		// node.next = node.next.prior
		return 1, nil
	}

}

/**
 * 删除第i个节点的数据
 */
func (l *LinkDList) Delete(i int) (int, error) {

	//判断位置是否合法
	if i < 1 || i > l.Length() {
		return 0, errors.New("超出线性表范围")
	}

	//累加器
	count := 0

	//获取数据元素域节点
	node := l.link

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

		node.next = node.next.next
		node.next.prior = node
		return 1, nil
	}

}

/**
 * 打印链表中的数据
 */
func (l *LinkDList) Print() {

	node := l.link
	fmt.Println("head:", node.data, node.next)
	node = node.next
	for node != nil {
		fmt.Printf("node: %v nodeptr: %p prior: %p  next: %p\n", node.data, node, node.prior, node.next)
		node = node.next
	}

}
