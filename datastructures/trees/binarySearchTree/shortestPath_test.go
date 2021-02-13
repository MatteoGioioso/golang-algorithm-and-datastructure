package binarySearchTree_test

import (
	"algo/datastructures/trees/binarySearchTree"
	"github.com/onsi/gomega"
	"testing"
)

func Test_stack_ShortestPath(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	type args struct {
		values []int
	}
	type want struct {
		shortest int
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
				shortest: 3,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := binarySearchTree.ShortestPath(tt.args.values)

			g.Expect(s).Should(gomega.Equal(tt.want.shortest))
		})
	}
}

