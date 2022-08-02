package queue

import "strconv"

type PriorityQueue[T any] struct {
	data               []T
	lessThanComparator func(T, T) bool
}

func NewPriorityQueue[T any](lessThanComparator func(T, T) bool) PriorityQueue[T] {
	return PriorityQueue[T]{
		data:               make([]T, 0),
		lessThanComparator: lessThanComparator,
	}
}

func (receiver PriorityQueue[_]) Size() int {
	return len(receiver.data)
}

func (receiver PriorityQueue[_]) IsEmpty() bool {
	return receiver.Size() == 0
}

func (receiver PriorityQueue[_]) IsNotEmpty() bool {
	return !receiver.IsEmpty()
}

func (receiver PriorityQueue[_]) validIndex(i int) bool {
	return i >= 0 && i < receiver.Size()
}

func (receiver *PriorityQueue[_]) swap(i, j int) {
	if !(receiver.validIndex(i) && receiver.validIndex(j)) {
		panic("swap [" + strconv.Itoa(i) + ", " +
			strconv.Itoa(j) + "] out of range [" +
			strconv.Itoa(receiver.Size()) + "]")
	}
	tmp := receiver.data[i]
	receiver.data[i] = receiver.data[j]
	receiver.data[j] = tmp
}

func (receiver PriorityQueue[_]) fatherIndex(i int) int {
	return (i - 1) / 2
}

func (receiver PriorityQueue[_]) leftChildIndex(i int) int {
	return 2*i + 1
}

func (receiver PriorityQueue[_]) rightChildIndex(i int) int {
	return 2*i + 2
}

func (receiver PriorityQueue[_]) lessThanByIndex(i, j int) bool {
	return receiver.lessThanComparator(receiver.data[i], receiver.data[j])
}

func (receiver *PriorityQueue[T]) Push(value T) {
	receiver.data = append(receiver.data, value)
	cur := receiver.Size() - 1
	for cur > 0 {
		father := receiver.fatherIndex(cur)
		if receiver.lessThanByIndex(cur, father) {
			receiver.swap(cur, father)
		}
		cur = father
	}
}

func (receiver PriorityQueue[T]) Top() T {
	if receiver.IsEmpty() {
		panic("Empty Priority Queue!")
	}
	return receiver.data[0]
}

func (receiver *PriorityQueue[T]) Pop() T {
	res := receiver.Top()
	receiver.swap(0, receiver.Size()-1)
	receiver.data = receiver.data[0 : receiver.Size()-1]
	cur := 0
	for receiver.leftChildIndex(cur) < receiver.Size() {
		left := receiver.leftChildIndex(cur)
		right := receiver.rightChildIndex(cur)
		if right == receiver.Size() {
			if receiver.lessThanByIndex(left, cur) {
				receiver.swap(cur, left)
			}
			break
		} else {
			less := right
			if receiver.lessThanByIndex(left, right) {
				less = left
			}
			if receiver.lessThanByIndex(less, cur) {
				receiver.swap(cur, less)
				cur = less
			} else {
				break
			}
		}
	}
	return res
}
