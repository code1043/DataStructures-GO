package queue

import (
	"errors"
	"fmt"
)

/**
 * 队列的顺序存储实现
 */

type node struct {
	elem ElemType
	next *node
}

type QueueLink struct {
	rear  *node
	front *node
	len   int
}

func NewQueueQueueLink() *QueueLink {

	q := QueueLink{front: &node{elem: 1043, next: nil}, len: 0}
	q.rear = q.front
	return &q
}

func (q *QueueLink) EnQueueSq(e ElemType) (int, error) {

	node := &node{elem: e, next: nil}

	q.rear.next = node
	q.rear = q.rear.next

	q.len += 1
	return 1, nil

}

func (q *QueueLink) DeQueueSq(e *ElemType) (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("队列已空")
	}

	//获取首元节点
	n := q.front.next
	*e = n.elem
	q.front.next = n.next
	q.len -= 1

	if q.front.next == nil {
		q.rear.next = nil
	}
	return 1, nil
}

func (q *QueueLink) IsEmpty() bool {
	return q.front.next == q.rear.next
}

func (q *QueueLink) Length() int {

	return q.len
}

func (q *QueueLink) Print() {

	fmt.Println("-----------开始打印信息-------------")
	fmt.Println(q)
	p := q.front
	for p != nil {
		fmt.Printf("%v ", p.elem)
		p = p.next
	}
	fmt.Println()
	fmt.Println("-----------结束打印信息-------------")

}
