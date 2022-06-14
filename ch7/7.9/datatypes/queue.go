package utils

import (
	"errors"
	"fmt"

	ti "github.com/DagmarC/gopl-solutions/ch7/7.8/trackpkg"
)

type Queue struct {
	elements []ti.Less
}

func (q *Queue) Init() {
	q.elements = make([]ti.Less, 0, 5)
}

func (q *Queue) Elements() []ti.Less {
	return q.elements
}

func (q *Queue) Reverse() []ti.Less {
	for i := 0; i < q.Length()/2; i++ {
		j := q.Length() - i - 1
		q.elements[i], q.elements[j] = q.elements[j], q.elements[i]
	}
	return q.elements
}

func (q *Queue) Enqueue(el ti.Less) {
	q.elements = append(q.elements, el)
}

func (q *Queue) Dequeue() (ti.Less, error) {
	el, err := q.Peek()
	if err != nil {
		return nil, err
	}
	q.elements = q.elements[1:]
	return el, nil
}

func (q *Queue) Peek() (ti.Less, error) {
	if q.IsEmpty() {
		return nil, errors.New("queue is empty")
	}
	return q.elements[0], nil
}

func (q *Queue) Length() int {
	return len(q.elements)
}

func (q *Queue) IsEmpty() bool {
	return q.Length() == 0
}

func (q *Queue) Print() {
	fmt.Println(q.elements)
}

func (q *Queue) Get(x int) ti.Less {
	if x < 0 || x >= q.Length() {
		return nil
	}
	return q.elements[x]
}

func (q *Queue) RemoveLast() {
	len := q.Length()
	if len == 1 {
		q.elements = []ti.Less{} // make it empty
	}
	q.elements = q.elements[:len-2]
}

// func (q *Queue) Max() int {
// 	if !q.IsEmpty() {
// 		sort.Ints(q.elements)
// 		return q.elements[len(q.elements)-1]
// 	}
// 	return 0
// }

// func (q *Queue) Min() int {
// 	if !q.IsEmpty() {
// 		sort.Ints(q.elements)
// 		return q.elements[0]
// 	}
// 	return 0
// }
