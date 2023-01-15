package liner

/**
 * 线性表的顺序存储结构实现
 */

import (
	"errors"
)

type ElemType int

type SqList struct {
	elem   []ElemType
	length int //当前长度
}

// 创建一个线性表，存储结构为顺序
func NewSqList() *SqList {
	return &SqList{}
}

// 销毁线性表
func DestroySqList(sq **SqList) {
	(*sq).elem = nil
	(*sq).length = 0
	*sq = nil
}

/**
 * 获取值
 * 根据位置i获取相应位置数据元素的内容
 * 成功: 返回该位置i相应元素
 * 失败: 返回0并且返回相应错误信息
 *
 * 注意，这里考虑用户传的位置是基于逻辑位置，即从1开始。物理存储是从0开始，因此-1
 *
 * 算法复杂度常量阶O(1)
 */
func (sq *SqList) Get(i int) (ElemType, error) {

	//判断线性表是否为空
	if sq.IsEmpty() {
		return 0, errors.New("线性表为空")
	}

	//判断位置i是否越界
	if i < 1 || i > sq.length {
		return 0, errors.New("位置越界")
	}

	return sq.elem[i-1], nil
}

/**
 * 按值查找
 * 值是否在线性表中
 * 存在，返回该线性表中的位置序号
 * 不存在，返回0
 *
 * 注意，这里返回的是线性表的逻辑位置，即从1开始，因此返回的位置需要+1
 *
 * 平均查找长度ASL
 * 最好为1,最坏为n
 * (1 + 2 + 3 + ... + n) / n = n * ( 1 + n) / 2 / n    所以它的同阶为O(n)
 * 实际上最坏是n + 1
 */
func (sq *SqList) Locate(e ElemType) (int, error) {

	position := 0

	if sq.IsEmpty() {
		return 0, errors.New("线性表为空")
	}

	for i := 0; i < sq.length; i++ {
		if e == sq.elem[i] {
			position = i + 1
			break
		}
	}

	return position, nil

}

/**
 * 插入元素到相应的位置
 * 成功,返回逻辑位置索引
 * 失败,返回0
 *
 * 最坏移动n次
 * 最好移动0次
 * 平均同阶为O(n)
 */
func (sq *SqList) Insert(e ElemType, position int) (int, error) {

	//判断位置是否合法
	if position < 1 || position > (sq.length+1) {
		return 0, errors.New("插入位置不合法")
	}

	//如果位置等于length + 1，表示用户想把元素作为最后一个元素的后继
	//插入元素

	//开辟一个空间
	// 第一种实现方法
	//
	// sq.elem = append(sq.elem, 0)
	// p := position - 1
	// for i := sq.length - 1; i > p; i-- {
	// 	sq.elem[i+1] = sq.elem[i]
	// }
	// sq.elem[p] = e
	// sq.length += 1

	// 第二种实现方法
	if position == sq.length+1 {
		sq.elem = append(sq.elem, e)
	} else {
		sq.elem = append(sq.elem, 0)
		p := position - 1
		for i := sq.length; i > p; i-- {
			sq.elem[i] = sq.elem[i-1]
		}
		sq.elem[p] = e
	}
	//更新长度
	sq.length += 1

	return position, nil

}

/**
 * 删除一个元素
 * 根据位置删除元素
 * 成功，返回位置
 * 失败，返回0
 *
 */
func (sq *SqList) Delete(position int) (int, error) {
	//判断位置是否合法
	if position < 1 || position > (sq.length+1) {
		return 0, errors.New("删除位置不合法")
	}

	//减少一个空间
	sq.length -= 1
	for i := position - 1; i < sq.length; i++ {
		sq.elem[i] = sq.elem[i+1]
	}
	sq.elem[sq.length] = 0

	return position, nil
}

/**
 * 删除一个元素
 * 根据元素删除元素
 * 成功，返回删除元素
 * 失败，返回0
 */
func (sq *SqList) DeleteAtElem(e ElemType) (ElemType, error) {

	//获取位置
	p, err := sq.Locate(e)
	if err != nil {
		return 0, errors.New("未找到元素")
	}

	//删除位置
	_, err = sq.Delete(p)
	if err != nil {
		return 0, errors.New("未找到元素")
	}

	return e, nil
}

// 清空线性表
func (sq *SqList) Clear() {
	sq.length = 0
}

// 获取线性表的长度
func (sq *SqList) GetLength() int {
	return sq.length
}

// 判断先线性表是否为空
func (sq *SqList) IsEmpty() bool {
	return sq.length == 0
}
