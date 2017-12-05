package ds

import (
	"testing"
)

func TestQueue(t *testing.T) {
	q := NewQueue()
	if q == nil {
		t.Fatal("can't create a queue")
	}

	for i := 0; i < 1000000; i++ {
		q.Put(i)

		if q.Size() != 1 {
			t.Fatalf("Size() failed with error:%v, expected=%v", q.Size(), 1)
		}

		el, err := q.Get()
		if err != nil {
			t.Fatalf("Get() failed with error:empty queue i(%v)", i)
		}

		if el != i {
			t.Fatalf("Get() failed with error:el(%v) not equal to i(%v)", el, i)
		}
	}

	q.Put(1)
	q.Put(2)
	if q.Size() != 2 {
		t.Fatalf("Size() failed with error:%v, expected=%v", q.Size(), 2)
	}
	q.Clear()

	if q.Size() != 0 {
		t.Fatalf("Size() failed with error:%v, expected=%v", q.Size(), 0)
	}

	el, err := q.Get()
	if err == nil {
		t.Fatalf("Get() failed with error:get element(%v) from empty queue", el)
	}

	// test struct queue
	type testData struct {
		name string
		age  int
	}
	td := testData{name: "n", age: 1}

	q.Put(td)
	el2, err := q.Get()
	if err != nil {
		t.Fatalf("Get() struct failed with error:%v", err)
	}
	data := el2.(testData)
	if data.name != td.name || data.age != td.age {
		t.Fatalf("Get() struct failed with error:got=%v, expected=%v", data, el2)
	}
}
