package datastructures

import "testing"

//     expect(heap.delete()).toEqual(7);
//     expect(heap.delete()).toEqual(8);
//     expect(heap.delete()).toEqual(69);
//     expect(heap.delete()).toEqual(420);
//     expect(heap.length).toEqual(0);

func toEqual(t *testing.T, value int, expected int) {
	if value != expected {
		t.Errorf("%v is not equal to %v", value, expected)
	}
}

func TestMinHeap(t *testing.T) {
	minHeap := MinHeap{}
	minHeap.Add(5)
	minHeap.Add(3)
	minHeap.Add(69)
	minHeap.Add(420)
	minHeap.Add(4)
	minHeap.Add(1)
	minHeap.Add(8)
	minHeap.Add(7)
	toEqual(t, minHeap.Length, 8)

	toEqual(t, minHeap.Delete(), 1)
	toEqual(t, minHeap.Delete(), 3)
	toEqual(t, minHeap.Delete(), 4)
	toEqual(t, minHeap.Delete(), 5)
	toEqual(t, minHeap.Length, 4)

	toEqual(t, minHeap.Delete(), 7)
	toEqual(t, minHeap.Delete(), 8)
	toEqual(t, minHeap.Delete(), 69)
	toEqual(t, minHeap.Delete(), 420)
	toEqual(t, minHeap.Length, 0)
}
