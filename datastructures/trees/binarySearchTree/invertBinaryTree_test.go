package binarySearchTree_test

import (
	"algo/datastructures/trees/binarySearchTree"
	"fmt"
	"github.com/onsi/gomega"
	"testing"
)

func Test_stack_InvertBinaryTree(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	type args struct {
		values []int
	}
	type want struct {
		invertedTree []int
	}

	tests := []struct {
		name    string
		args    args
		want    want
		wantErr bool
	}{
		{
			name: "should push items into the stack",
			args: args{values: []int{7, 20, 5, 15, 3, 4, 33, 2, 6}},
			want: want{
				invertedTree: []int{7, 20, 5, 33, 15, 6, 3, 4, 2},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			invertedTree := binarySearchTree.InvertTree(tt.args.values)
			fmt.Println(invertedTree)

			g.Expect(invertedTree).Should(gomega.Equal(tt.want.invertedTree))
		})
	}
}


