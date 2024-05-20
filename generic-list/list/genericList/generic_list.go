package genericlist

type GenericList[T comparable] struct {
	data []T
}

func New[T comparable]() *GenericList[T] {
	return &GenericList[T]{
		data: []T{},
	}
}

func (g *GenericList[T]) Insert(value T) {
	g.data = append(g.data, value)
}

func (g *GenericList[T]) Get(index int) T {
	if index > len(g.data)-1 {
		panic("Index too high!")
	}

	for i := 0; i < len(g.data); i++ {
		if i == index {
			return g.data[i]
		}
	}
	panic("Value not found!")
}

func (g *GenericList[T]) Remove(index int) {
	if index > len(g.data)-1 {
		panic("Index too high!")
	}

	for i := 0; i < len(g.data); i++ {
		if i == index {
			g.data = append(g.data[:i], g.data[i+1:]...)
			return
		}
	}
	panic("Value to be deletd not found!")
}

func (g *GenericList[T]) RemoveByValue(val T) {
	for i := 0; i < len(g.data); i++ {
		if val == g.data[i] {
			g.data = append(g.data[:i], g.data[i+1:]...)
			return
		}
	}
	panic("Value to be deletd not found!")

}
