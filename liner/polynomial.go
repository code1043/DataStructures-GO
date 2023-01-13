package liner

type Polynomial struct { //多项式非零项的定义
	p float64 //系数
	e int32   //指数
}

type SqList_polynomial struct { //多项式的顺序存储结构类型为SqList
	elem   *Polynomial //存储空间的基地址
	length int         //多项式中当前项的个数
}
