package acmachine

import "testing"

func TestQueuePush(t *testing.T) {
	q := NewQueue(1).(*queue)
	q.Push(1)
	q.Push(1)
	if len(q.arr) != 2 {
		t.Errorf("push arr is not 2 %d", len(q.arr))
	}
	q.Push(2)
	if len(q.arr) != 3 {
		t.Errorf("push arr is not 3 %d", len(q.arr))
	}

	if q.cap != 3 {
		t.Errorf("push cap is not 3 %d", q.cap)
	}

	q.Pop()
	q.Pop()
	q.Push(1)
	if len(q.arr) != 2 {
		t.Errorf("push arr is not 2 %d", len(q.arr))
	}
	if q.cap != 2 {
		t.Errorf("push cap is not 3 %d", q.cap)
	}
}

func TestQueuePop(t *testing.T) {
	q := NewQueue(1).(*queue)
	q.Push(1)
	q.Push(2)
	q.Push(3)
	q.Push(4)
	q.Push(5)

	if v, err := q.Pop(); err != nil || v.(int) != 1 {
		t.Errorf("pop value is not 1 %v", v)
	}
	if v, err := q.Pop(); err != nil || v.(int) != 2 {
		t.Errorf("pop value is not 2 %v", v)
	}
	if v, err := q.Pop(); err != nil || v.(int) != 3 {
		t.Errorf("pop value is not 3 %v", v)
	}
	if v, err := q.Pop(); err != nil || v.(int) != 4 {
		t.Errorf("pop value is not 4 %v", v)
	}
	if v, err := q.Pop(); err != nil || v.(int) != 5 {
		t.Errorf("pop value is not 5 %v", v)
	}
	if !q.Empty() {
		t.Errorf("empty is not null %d", len(q.arr))
	}
}

func TestQueueEmpty(t *testing.T) {
	q := NewQueue(1).(*queue)
	if !q.Empty() {
		t.Errorf("init queue is not empty")
	}
	q.Push(1)
	if q.Empty() {
		t.Errorf("queue empty after push")
	}

	q.Pop()
	if !q.Empty() {
		t.Errorf("queue not empty after pop")
	}
}