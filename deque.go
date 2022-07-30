package deque

import "strconv"

type Deque struct {
	data                 []int
	size                 int
	beginIndex, endIndex int
}

const DequeMinCap = 8

const DequeInitCap = 8

func (receiver Deque) Size() int {
	return receiver.size
}

func (receiver *Deque) incSize() {
	receiver.size++
}

func (receiver *Deque) decSize() {
	receiver.size--
}

func (receiver Deque) IsEmpty() bool {
	return receiver.Size() == 0
}

func (receiver Deque) IsNotEmpty() bool {
	return !receiver.IsEmpty()
}

func (receiver Deque) Cap() int {
	return len(receiver.data)
}

func NewDeque() Deque {
	return Deque{data: make([]int, DequeInitCap)}
}

func (receiver Deque) modulo(i int) int {
	for i < 0 {
		i += receiver.Cap()
	}
	return i % receiver.Cap()
}

func (receiver Deque) inc(i int) int {
	return receiver.modulo(i + 1)
}

func (receiver Deque) dec(i int) int {
	return receiver.modulo(i - 1)
}

func (receiver *Deque) resize(newSize int) {
	newCap := receiver.max(DequeMinCap, newSize)
	newData := make([]int, newCap)
	for i := 0; i < receiver.Size(); i++ {
		newData[i] = receiver.Get(i)
	}
	receiver.data = newData
	receiver.beginIndex = 0
	receiver.endIndex = receiver.modulo(receiver.Size())
}

func (receiver *Deque) expand() {
	receiver.resize(receiver.Cap() * 2)
}

func (receiver *Deque) shrink() {
	receiver.resize((receiver.Cap() + 1) / 2)
}

func (receiver Deque) isFull() bool {
	return receiver.Size() == receiver.Cap()
}

func (receiver Deque) isToShrink() bool {
	return receiver.Cap() > DequeMinCap && receiver.Size() <= receiver.Cap()/4
}

func (receiver Deque) Get(index int) int {
	if index < 0 || index >= receiver.Size() {
		panic("index [" + strconv.Itoa(index) +
			"] out of range [" + strconv.Itoa(receiver.Size()) + "]")
	}
	return receiver.data[receiver.modulo(receiver.beginIndex+index)]
}

func (receiver *Deque) Set(index int, value int) {
	if index < 0 || index >= receiver.Size() {
		panic("index [" + strconv.Itoa(index) +
			"] out of range [" + strconv.Itoa(receiver.Size()) + "]")
	}
	receiver.data[receiver.modulo(receiver.beginIndex+index)] = value
}

func (receiver Deque) Back() int {
	if receiver.IsEmpty() {
		panic("the deque is currently empty!")
	}
	return receiver.Get(receiver.Size() - 1)
}

func (receiver Deque) Front() int {
	if receiver.IsEmpty() {
		panic("the deque is currently empty!")
	}
	return receiver.Get(0)
}

func (receiver *Deque) PushBack(item int) {
	if receiver.isFull() {
		receiver.expand()
	}
	receiver.incSize()
	receiver.Set(receiver.Size()-1, item)
	receiver.endIndex = receiver.inc(receiver.endIndex)
}

func (receiver *Deque) PushFront(item int) {
	if receiver.isFull() {
		receiver.expand()
	}
	receiver.incSize()
	receiver.beginIndex = receiver.dec(receiver.beginIndex)
	receiver.Set(0, item)
}

func (receiver *Deque) PopBack() int {
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

func (receiver *Deque) PopFront() int {
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

func (receiver Deque) min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func (receiver Deque) max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func (receiver Deque) ToString() string {
	res := "["
	if receiver.IsNotEmpty() {
		res += strconv.Itoa(receiver.Front())
	}
	for i := 1; i < receiver.Size(); i++ {
		res += ", " + strconv.Itoa(receiver.Get(i))
	}
	res += "]"
	return res
}
