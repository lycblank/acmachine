package acmachine

import "errors"

var (
	QueueErrEmpty = errors.New("queue empty")
)

type Queue interface {
	Pop() (interface{}, error)
	Push(v interface{})
	Empty() bool
}

type queue struct {
	cap            int
	arr            []interface{}
	noUseThreshold int
}

func NewQueue(noUseThreshold int) Queue {
	return &queue{
		noUseThreshold: noUseThreshold,
	}
}

func (q *queue) Pop() (interface{}, error) {
	if len(q.arr) <= 0 {
		return nil, QueueErrEmpty
	}
	ret := q.arr[0]
	q.arr = q.arr[1:]
	return ret, nil
}

func (q *queue) Push(v interface{}) {
	if q.cap-len(q.arr) >= q.noUseThreshold {
		tmp := make([]interface{}, len(q.arr), len(q.arr)+1)
		copy(tmp, q.arr)
		q.arr = tmp
		q.cap = len(q.arr)
	}
	q.cap++
	q.arr = append(q.arr, v)
}

func (q *queue) Empty() bool {
	return len(q.arr) <= 0
}
