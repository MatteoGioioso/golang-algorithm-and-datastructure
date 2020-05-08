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

func TestBinaryHeap_Poll(t *testing.T) {
	type fields struct {
		Items []int
	}
	tests := []struct {
		name   string
		fields fields
		want   []int
	}{
		{name: "Pool first element from queue", fields: fields{
			Items: []int{1, 5, 1, 8, 6, 2},
		}, want: []int{1, 5, 2, 8, 6}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := BinaryHeap{
				Items: tt.fields.Items,
			}
			b.Poll()
			got := b.ReturnHeap()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Add() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinaryHeap_Remove(t *testing.T) {
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
		{
			name:   "should remove an item",
			fields: fields{Items: []int{1, 5, 1, 8, 6, 2}},
			args:   args{item: 6},
			want:   []int{1, 2, 1, 8, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &BinaryHeap{
				Items: tt.fields.Items,
			}
			b.Remove(tt.args.item)
			got := b.ReturnHeap()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Add() got = %v, want %v", got, tt.want)
			}
		})
	}
}
