package binarySearchTree_test

import (
	"algo/datastructures/trees/binarySearchTree"
	"github.com/onsi/gomega"
	"testing"
)

func Test_stack_FindAllPaths(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	type args struct {
		values []int
	}
	type want struct {
		allPaths [][]int
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
				allPaths: [][]int{{7, 5, 4, 2}, {7, 5, 6}, {7, 20, 15, 10}, {7, 20, 33, 25}},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := binarySearchTree.FindAllPaths(tt.args.values)

			g.Expect(s).Should(gomega.Equal(tt.want.allPaths))
		})
	}
}

