package queue

import "testing"

func TestAddQueue(t *testing.T) {
	q := New(3)
	for i := 0; i < 3; i++ {
		if len(q.items) != i {
			t.Errorf("incorrect queue element count: %v, want %v", len(q.items), i)
		}
		if !q.Append(i) {
			t.Errorf("failed to append element %v to queue", i)
		}

	}
	if q.Append(4) {
		t.Errorf("should be unable to add %v to queue, maximum capacity of %v reached", 4, q.capacity)
	}
}

func TestNext(t *testing.T) {
	q := New(3)
	for i := 0; i < 3; i++ {
		q.Append(i)
	}
	for i := 0; i < 3; i++ {
		item, ok := q.Next()
		if !ok {
			t.Errorf("should be able to get item from queue")
		}
		if item != i {
			t.Errorf("got item in wrong order %v, want %v", item, i)
		}
	}
	item, ok := q.Next()
	if ok {
		t.Errorf("queue should be empty, received %v", item)
	}
}
