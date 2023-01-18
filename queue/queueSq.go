package queue

import (
	"errors"
	"fmt"
)

/**
 * 队列的顺序存储实现
 */

type QueueSq struct {
	elem  []ElemType
	front int
	rear  int
	cap   int
}

func NewQueueSq(cap int) *QueueSq {
	return &QueueSq{elem: make([]ElemType, cap), front: 0, rear: 0, cap: cap}
}

func (q *QueueSq) EnQueueSq(e ElemType) (int, error) {

	if q.IsFull() {
		return 0, errors.New("队列已满")
	}

	q.elem[q.rear] = e
	q.rear = (q.rear + 1) % q.cap

	return 1, nil

}

func (q *QueueSq) DeQueueSq(e *ElemType) (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("队列已空")
	}
	*e = q.elem[q.front]
	q.elem[q.front] = 0 //这里不应该这样写
	q.front = (q.front + 1) % q.cap

	return 1, nil
}

func (q *QueueSq) IsEmpty() bool {
	return q.front == q.rear
}

func (q *QueueSq) IsFull() bool {
	return (q.rear+1)%q.cap == q.front //少用一个空间法
}

func (q *QueueSq) Length() int {

	/**
	 * 因为是循环队列，所以可以把该队列想象成钟表。从A到B，有两个方向可供选择。
	 * 因此，负数的正补数就是从A到B的距离。负数的补数 = 负数 + "队列容量"，求得。
	 * 因为正数的补数是其本身，若 + 队列容量 则超出队列容量 % 队列容量 丢弃即可。
	 * 同余式
	 */

	return ((q.rear - q.front) + q.cap) % q.cap //
}

func (q *QueueSq) Print() {

	len := q.Length()

	fmt.Println("-----------开始打印信息-------------")
	fmt.Println(q)
	for i := 0; i < len; i++ {
		index := (q.front + i) % q.cap
		fmt.Printf("%v ", q.elem[index])
	}
	fmt.Println()
	fmt.Println("-----------结束打印信息-------------")

}
