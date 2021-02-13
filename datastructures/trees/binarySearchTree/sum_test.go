package binarySearchTree_test

import (
	"algo/datastructures/trees/binarySearchTree"
	"github.com/onsi/gomega"
	"testing"
)

func Test_stack_Sum(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	type args struct {
		values []int
	}
	type want struct {
		total int
	}

	tests := []struct {
		name    string
		args    args
		want    want
		wantErr bool
	}{
		{
			name: "should push items into the stack",
			args: args{values: []int{2, 15, 1, 8}},
			want: want{
				total: 2 + 15 + 1 + 8,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := binarySearchTree.Sum(tt.args.values)

			g.Expect(s).Should(gomega.Equal(tt.want.total))
		})
	}
}
