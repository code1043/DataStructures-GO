package liner

// 图书信息
type Book struct {
	ISBN  string  // 图书ISBN
	Name  string  // 图书名字
	Price float64 // 图书价格
}

type SqList_book struct { //
	elem   *Book
	length int
}
