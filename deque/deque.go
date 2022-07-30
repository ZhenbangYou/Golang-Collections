package deque

import "strconv"

type Deque[T any] struct {
	data                 []T
	size                 int
	beginIndex, endIndex int
	cap                  int
}

const DequeMinCap = 8

const DequeInitCap = 8

func (receiver Deque[T]) Size() int {
	return receiver.size
}

func (receiver *Deque[T]) incSize() {
	receiver.size++
}

func (receiver *Deque[T]) decSize() {
	receiver.size--
}

func (receiver Deque[T]) IsEmpty() bool {
	return receiver.Size() == 0
}

func (receiver Deque[T]) IsNotEmpty() bool {
	return !receiver.IsEmpty()
}

func (receiver Deque[T]) Cap() int {
	return receiver.cap
}

func NewDeque[T any]() Deque[T] {
	return Deque[T]{data: make([]T, DequeInitCap), cap: DequeInitCap}
}

func (receiver Deque[T]) modulo(i int) int {
	for i < 0 {
		i += receiver.Cap()
	}
	return i % receiver.Cap()
}

func (receiver Deque[T]) inc(i int) int {
	return receiver.modulo(i + 1)
}

func (receiver Deque[T]) dec(i int) int {
	return receiver.modulo(i - 1)
}

func (receiver *Deque[T]) resize(newSize int) {
	newCap := receiver.max(DequeMinCap, newSize)
	newData := make([]T, newCap)
	for i := 0; i < receiver.Size(); i++ {
		newData[i] = receiver.Get(i)
	}
	receiver.data = newData
	receiver.beginIndex = 0
	receiver.endIndex = receiver.modulo(receiver.Size())
	receiver.cap = newSize
}

func (receiver *Deque[T]) expand() {
	receiver.resize(receiver.Cap() * 2)
}

func (receiver *Deque[T]) shrink() {
	receiver.resize((receiver.Cap() + 1) / 2)
}

func (receiver Deque[T]) isFull() bool {
	return receiver.Size() == receiver.Cap()
}

func (receiver Deque[T]) isToShrink() bool {
	return receiver.Cap() > DequeMinCap && receiver.Size() <= receiver.Cap()/4
}

func (receiver Deque[T]) Get(index int) T {
	if index < 0 || index >= receiver.Size() {
		panic("index [" + strconv.Itoa(index) +
			"] out of range [" + strconv.Itoa(receiver.Size()) + "]")
	}
	return receiver.data[receiver.modulo(receiver.beginIndex+index)]
}

func (receiver *Deque[T]) Set(index int, value T) {
	if index < 0 || index >= receiver.Size() {
		panic("index [" + strconv.Itoa(index) +
			"] out of range [" + strconv.Itoa(receiver.Size()) + "]")
	}
	receiver.data[receiver.modulo(receiver.beginIndex+index)] = value
}

func (receiver Deque[T]) Back() T {
	if receiver.IsEmpty() {
		panic("the deque is currently empty!")
	}
	return receiver.Get(receiver.Size() - 1)
}

func (receiver Deque[T]) Front() T {
	if receiver.IsEmpty() {
		panic("the deque is currently empty!")
	}
	return receiver.Get(0)
}

func (receiver *Deque[T]) PushBack(item T) {
	if receiver.isFull() {
		receiver.expand()
	}
	receiver.incSize()
	receiver.Set(receiver.Size()-1, item)
	receiver.endIndex = receiver.inc(receiver.endIndex)
}

func (receiver *Deque[T]) PushFront(item T) {
	if receiver.isFull() {
		receiver.expand()
	}
	receiver.incSize()
	receiver.beginIndex = receiver.dec(receiver.beginIndex)
	receiver.Set(0, item)
}

func (receiver *Deque[T]) PopBack() T {
	if receiver.IsEmpty() {
		panic("the deque is currently empty!")
	}
	res := receiver.Back()
	receiver.endIndex = receiver.dec(receiver.endIndex)
	receiver.decSize()
	if receiver.isToShrink() {
		receiver.shrink()
	}
	return res
}

func (receiver *Deque[T]) PopFront() T {
	if receiver.IsEmpty() {
		panic("the deque is currently empty!")
	}
	res := receiver.Front()
	receiver.beginIndex = receiver.inc(receiver.beginIndex)
	receiver.decSize()
	if receiver.isToShrink() {
		receiver.shrink()
	}
	return res
}

func (receiver Deque[T]) min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func (receiver Deque[T]) max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
