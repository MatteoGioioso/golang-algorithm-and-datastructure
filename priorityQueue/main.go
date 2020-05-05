package main

type BinaryHeap struct {
	Items []int
}

func (b *BinaryHeap) getLeftChildIndex(parentIndex int) int {
	return 2*parentIndex + 1
}

func (b *BinaryHeap) getRightChildIndex(parentIndex int) int {
	return 2*parentIndex + 2
}

func (b *BinaryHeap) getParentIndexFromLeftChild(leftChildIndex int) int {
	return (leftChildIndex - 1) / 2
}

func (b *BinaryHeap) getParentIndexFromRightChild(rightChildIndex int) int {
	return (rightChildIndex - 2) / 2
}

func (b BinaryHeap) isRightChild(index int) bool {
	if index%2 == 0 {
		return false
	}

	return true
}

func (b *BinaryHeap) Add(item int) {
	b.Items = append(b.Items, item)

	for i := len(b.Items); i >= 0; i-- {

	}
}

func main() {

}
