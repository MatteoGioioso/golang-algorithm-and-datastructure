package main

import (
	"reflect"
	"testing"
)

func TestBinaryHeap_Add(t *testing.T) {
	type fields struct {
		Items []int
	}
	type args struct {
		item int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []int
	}{
		{name: "should add a small number as right child to the heap", fields: fields{Items: []int{1, 5, 1, 8, 6, 2}},
			args: args{
				item: 0,
			}, want: []int{0, 5, 1, 8, 6, 2, 1}},

		{name: "should add a small number as left child to the heap", fields: fields{Items: []int{1, 5, 1, 8, 6}},
			args: args{
				item: 0,
			}, want: []int{0, 5, 1, 8, 6, 1}},

		{name: "should add a big number as right child to the heap", fields: fields{Items: []int{1, 5, 1, 8}},
			args: args{
				item: 4,
			}, want: []int{1, 4, 1, 8, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &BinaryHeap{
				Items: tt.fields.Items,
			}
			b.Add(tt.args.item)
			got := b.ReturnHeap()

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Add() got = %v, want %v", got, tt.want)
			}
		})
	}
}

//func TestBinaryHeap_bubbling(t *testing.T) {
//	type fields struct {
//		Items []int
//	}
//	type args struct {
//		i int
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		args   args
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			b := BinaryHeap{
//				Items: tt.fields.Items,
//			}
//		})
//	}
//}

//func TestBinaryHeap_getLeftChildIndex(t *testing.T) {
//	type fields struct {
//		Items []int
//	}
//	type args struct {
//		parentIndex int
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		args   args
//		want   int
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			b := &BinaryHeap{
//				Items: tt.fields.Items,
//			}
//			if got := b.getLeftChildIndex(tt.args.parentIndex); got != tt.want {
//				t.Errorf("getLeftChildIndex() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

//func TestBinaryHeap_getParentIndexFromLeftChild(t *testing.T) {
//	type fields struct {
//		Items []int
//	}
//	type args struct {
//		leftChildIndex int
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		args   args
//		want   int
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			b := &BinaryHeap{
//				Items: tt.fields.Items,
//			}
//			if got := b.getParentIndexFromLeftChild(tt.args.leftChildIndex); got != tt.want {
//				t.Errorf("getParentIndexFromLeftChild() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

//func TestBinaryHeap_getParentIndexFromRightChild(t *testing.T) {
//	type fields struct {
//		Items []int
//	}
//	type args struct {
//		rightChildIndex int
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		args   args
//		want   int
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			b := &BinaryHeap{
//				Items: tt.fields.Items,
//			}
//			if got := b.getParentIndexFromRightChild(tt.args.rightChildIndex); got != tt.want {
//				t.Errorf("getParentIndexFromRightChild() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

//func TestBinaryHeap_getRightChildIndex(t *testing.T) {
//	type fields struct {
//		Items []int
//	}
//	type args struct {
//		parentIndex int
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		args   args
//		want   int
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			b := &BinaryHeap{
//				Items: tt.fields.Items,
//			}
//			if got := b.getRightChildIndex(tt.args.parentIndex); got != tt.want {
//				t.Errorf("getRightChildIndex() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

//func TestBinaryHeap_isRightChild(t *testing.T) {
//	type fields struct {
//		Items []int
//	}
//	type args struct {
//		index int
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		args   args
//		want   bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			b := BinaryHeap{
//				Items: tt.fields.Items,
//			}
//			if got := b.isRightChild(tt.args.index); got != tt.want {
//				t.Errorf("isRightChild() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
