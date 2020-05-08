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

func (b *BinaryHeap) isRightChild(index int) bool {
	if index%2 == 0 {
		return true
	}

	return false
}

func (b *BinaryHeap) bubblingUp(i int) {
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

		b.bubblingUp(parentIndex)
	}

	return
}

func (b *BinaryHeap) bubblingDown(i int) {
	lastIndex := len(b.Items) - 1
	if i == lastIndex {
		return
	}

	currentChild := b.Items[i]
	leftChildIndex := b.getLeftChildIndex(i)
	rightChildIndex := b.getRightChildIndex(i)

	if leftChildIndex >= lastIndex || rightChildIndex >= lastIndex {
		return
	}

	leftChild := b.Items[leftChildIndex]
	rightChild := b.Items[rightChildIndex]

	if leftChild > rightChild {
		if currentChild > rightChild {
			b.Items[i] = rightChild
			b.Items[rightChildIndex] = currentChild
			b.bubblingDown(rightChildIndex)
		}
	} else {
		if currentChild > leftChild {
			b.Items[i] = leftChild
			b.Items[leftChildIndex] = currentChild
			b.bubblingDown(leftChildIndex)
		}
	}
}

func (b BinaryHeap) search(item int) *int {
	for index, node := range b.Items {
		if node == item {
			return &index
		}
	}

	return nil
}

func (b *BinaryHeap) Add(item int) {
	b.Items = append(b.Items, item)
	b.bubblingUp(len(b.Items) - 1)
}

func (b *BinaryHeap) Remove(item int) {
	index := b.search(item)
	if index == nil {
		return
	}
	lastIndex := len(b.Items) - 1
	lastChild := b.Items[lastIndex]
	b.Items[*index] = lastChild
	b.Items = b.Items[:lastIndex]
	b.bubblingUp(*index)
}

func (b *BinaryHeap) Poll() {
	if len(b.Items) == 0 {
		return
	}
	lastIndex := len(b.Items) - 1
	lastChild := b.Items[lastIndex]
	b.Items[0] = lastChild
	b.Items = b.Items[:lastIndex]
	b.bubblingDown(0)
}

func (b *BinaryHeap) ReturnHeap() []int {
	return b.Items
}
