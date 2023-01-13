package liner

const (
	LIST_INIT_SIZE = 1000
)

type SqList struct {
	elem   [LIST_INIT_SIZE]int
	length int //当前长度
}
