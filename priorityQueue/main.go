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
		return true
	}

	return false
}

func (b BinaryHeap) bubbling(i int) {
	var parentIndex int

	if b.isRightChild(i) {
		parentIndex = b.getParentIndexFromRightChild(i)
	} else {
		parentIndex = b.getParentIndexFromLeftChild(i)
	}

	parent := b.Items[parentIndex]
	rightChild := b.Items[i]

	if rightChild < parent {
		b.Items[parentIndex] = rightChild
		b.Items[i] = parent

		if parentIndex == 0 {
			return
		}

		b.bubbling(parentIndex)
	}

	return
}

func (b *BinaryHeap) Add(item int) {
	b.Items = append(b.Items, item)
	b.bubbling(len(b.Items) - 1)
}

func (b BinaryHeap) ReturnHeap() []int {
	return b.Items
}

func main() {

}
