package util

// Val 用于替代nil来确认是否存在某个值
type Val[T any] interface {

	// IsExist 返回T类型的值是否存在
	// 当返回为true时，Val必须返回可用的数据
	// 当返回为false时，Val返回类型T的默认值
	IsExist() bool

	// Val 获取数据的值
	Val() T
}

type Value[T any] struct {
	inDB  bool
	value T
}

func NewValue[T any](inDB bool, value T) Val[T] {
	return Value[T]{
		inDB:  inDB,
		value: value,
	}
}

// IsExist 是否存在
func (v Value[T]) IsExist() bool {
	return v.inDB
}

// Val 获取数据的值
func (v Value[T]) Val() T {
	return v.value
}
