package binarySearchTree_test

import (
	"algo/datastructures/trees/binarySearchTree"
	"fmt"
	"github.com/onsi/gomega"
	"testing"
)

func Test_stack_Insert(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	type args struct {
		values []int
	}
	type want struct {
		size int
	}

	tests := []struct {
		name    string
		args    args
		want    want
		wantErr bool
	}{
		{
			name: "should push items into the stack",
			args: args{values: []int{7, 20, 5, 15, 10, 4, 33, 2, 25, 6}},
			want: want{
				size: 10,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := binarySearchTree.New()

			for _, value := range tt.args.values {
				s.Insert(value)
			}

			g.Expect(s.Size()).Should(gomega.Equal(tt.want.size))
		})
	}
}

func Test_stack_LevelOrderTraversal(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	type args struct {
		values []int
	}
	type want struct {
		size                     int
		levelOrderTraversalArray []int
	}

	tests := []struct {
		name    string
		args    args
		want    want
		wantErr bool
	}{
		{
			name: "should push items into the stack",
			args: args{values: []int{7, 20, 5, 15, 10, 4, 33, 2, 25, 6}},
			want: want{
				size:                     10,
				levelOrderTraversalArray: []int{7, 5, 20, 4, 6, 15, 33, 2, 10, 25},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := binarySearchTree.New()

			for _, value := range tt.args.values {
				s.Insert(value)
			}

			g.Expect(s.Size()).Should(gomega.Equal(tt.want.size))
			s.LevelOrderTraversal()
			g.Expect(s.LevelOrderTraversal()).Should(gomega.Equal(tt.want.levelOrderTraversalArray))
		})
	}
}

func Test_stack_TreeTraversal(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	type args struct {
		values []int
	}
	type want struct {
		size               int
		treeTraversalArray []int
	}

	tests := []struct {
		name    string
		args    args
		want    want
		wantErr bool
	}{
		{
			name: "should push items into the stack",
			args: args{values: []int{7, 20, 5, 15, 10, 4, 33, 2, 25, 6}},
			want: want{
				size:               10,
				treeTraversalArray: []int{7, 5, 4, 2, 6, 20, 15, 10, 33, 25},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := binarySearchTree.New()

			for _, value := range tt.args.values {
				s.Insert(value)
			}

			g.Expect(s.Size()).Should(gomega.Equal(tt.want.size))
			g.Expect(s.TreeTraversal()).Should(gomega.Equal(tt.want.treeTraversalArray))
		})
	}
}

func Test_stack_TreeTraversalRecv(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	type args struct {
		values []int
	}
	type want struct {
		size               int
		treeTraversalArray []int
	}

	tests := []struct {
		name    string
		args    args
		want    want
		wantErr bool
	}{
		{
			name: "should push items into the stack",
			args: args{values: []int{7, 20, 5, 15, 10, 4, 33, 2, 25, 6}},
			want: want{
				size:               10,
				treeTraversalArray: []int{7, 5, 4, 2, 6, 20, 15, 10, 33, 25},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := binarySearchTree.New()

			for _, value := range tt.args.values {
				s.Insert(value)
			}

			g.Expect(s.Size()).Should(gomega.Equal(tt.want.size))
			g.Expect(s.TreeTraversalRecv()).Should(gomega.Equal(tt.want.treeTraversalArray))
		})
	}
}


func Test_stack_Search(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	type args struct {
		values []int
		search int
	}
	type want struct {
		found int
	}

	tests := []struct {
		name    string
		args    args
		want    want
		wantErr bool
	}{
		{
			name: "should push items into the stack",
			args: args{values: []int{7, 20, 5, 15, 10, 4, 33, 2, 25, 6}, search: 4},
			want: want{
				found: 4,
			},
			wantErr: false,
		},

		{
			name: "should push items into the stack",
			args: args{values: []int{7, 20, 5, 15, 10, 4, 33, 2, 25, 6}, search: 50},
			want: want{
				found: 0,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := binarySearchTree.New()

			for _, value := range tt.args.values {
				s.Insert(value)
			}

			search, err := s.Search(tt.args.search)
			if err != nil {
				if tt.wantErr {
					g.Expect(err).Should(gomega.MatchError(fmt.Errorf("not found")))
				} else {
					g.Expect(err).ShouldNot(gomega.HaveOccurred())
				}
			}

			g.Expect(search).Should(gomega.Equal(tt.want.found))
		})
	}
}

func Test_stack_Remove(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	type args struct {
		values   []int
		toRemove int
	}
	type want struct {
		levelOrderTraversalArray []int
		size                     int
	}

	tests := []struct {
		name    string
		args    args
		want    want
		wantErr bool
	}{
		{
			name: "should remove leaf node",
			args: args{values: []int{7, 20, 5, 15, 10, 4, 33, 2, 25, 6}, toRemove: 10},
			want: want{
				levelOrderTraversalArray: []int{7, 5, 20, 4, 6, 15, 33, 2, 25},
				size:                     9,
			},
			wantErr: false,
		},
		{
			name: "should remove a node that has a left subtree only",
			args: args{values: []int{7, 20, 5, 15, 10, 4, 33, 2, 25, 6}, toRemove: 15},
			want: want{
				levelOrderTraversalArray: []int{7, 5, 20, 4, 6, 10, 33, 2, 25},
				size:                     9,
			},
			wantErr: false,
		},
		{
			name: "should remove a node that has a right subtree only",
			args: args{values: []int{7, 20, 5, 15, 10, 4, 33, 2, 25, 6, 34, 35}, toRemove: 34},
			want: want{
				levelOrderTraversalArray: []int{7, 5, 20, 4, 6, 15, 33, 2, 10, 25, 35},
				size:                     11,
			},
			wantErr: false,
		},
		{
			name: "should remove a node that has both right and left subtrees",
			args: args{values: []int{7, 20, 5, 15, 10, 4, 33, 2, 25, 6}, toRemove: 20},
			want: want{
				levelOrderTraversalArray: []int{7, 5, 25, 4, 6, 15, 33, 2, 10},
				size:                     9,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := binarySearchTree.New()

			for _, value := range tt.args.values {
				s.Insert(value)
			}

			if err := s.Remove(tt.args.toRemove); err != nil {
				if tt.wantErr {
					g.Expect(err).Should(gomega.MatchError(fmt.Errorf("not found")))
				} else {
					g.Expect(err).ShouldNot(gomega.HaveOccurred())
				}
			}

			g.Expect(s.Size()).Should(gomega.Equal(tt.want.size))
			g.Expect(s.LevelOrderTraversal()).Should(gomega.Equal(tt.want.levelOrderTraversalArray))
		})
	}
}
