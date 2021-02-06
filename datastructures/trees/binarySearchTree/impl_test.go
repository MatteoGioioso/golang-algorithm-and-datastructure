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
			s.LevelOrderTraversal()
			g.Expect(s.LevelOrderTraversal()).Should(gomega.Equal([]int{7, 5, 20, 4, 6, 15, 33, 2, 10, 25}))
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
