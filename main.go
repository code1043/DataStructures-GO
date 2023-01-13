package main

import (
	"dataStructure/liner"
	"fmt"
)

func main() {

	exampleSqList()

}

/**
 * 顺序存储结构案例
 */
func exampleSqList() {

	s := liner.NewSqList()

	for i := 1; i <= 10; i++ {
		s.Insert(liner.ElemType(i), i)

	}

	fmt.Println(s)

	_, err := s.Insert(liner.ElemType(100), 5)
	fmt.Println(s, err)

	_, err = s.Insert(liner.ElemType(100), 12)
	fmt.Println(s, err)

	_, err = s.Insert(liner.ElemType(100), 15)
	fmt.Println(s, err)

	//以下有一个bug
	p, err := s.Get(12)
	fmt.Println(p, err)

	_, err = s.DeleteAtElem(p)
	fmt.Println(s, err)

	_, err = s.Delete(11)
	fmt.Println(s, err)

	_, err = s.Delete(12)
	fmt.Println(s, err)
}
