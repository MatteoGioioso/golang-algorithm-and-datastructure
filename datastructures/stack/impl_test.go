package stack_test

import (
	"algo/datastructures/stack"
	"github.com/onsi/gomega"
	"testing"
)

func Test_stack_Push(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	type args struct {
		value interface{}
	}
	type want struct {
		size          int
		topOfTheStack string
	}

	tests := []struct {
		name    string
		args    args
		want    want
		wantErr bool
	}{
		{name: "should push items into the stack", want: want{
			size:          2,
			topOfTheStack: "",
		}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := stack.New()

			s.Push("one")
			s.Push("two")

			g.Expect(s.Size()).Should(gomega.Equal(tt.want.size))
		})
	}
}

func Test_stack_Pop(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	type args struct {
		value interface{}
	}
	type want struct {
		size          int
		popItem string
	}

	tests := []struct {
		name    string
		args    args
		want    want
		wantErr bool
	}{
		{name: "should push items into the stack", want: want{
			size:          3,
			popItem: "four",
		}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := stack.New()

			s.Push("one")
			s.Push("two")
			s.Push("three")
			s.Push("four")

			pop, err := s.Pop()
			g.Expect(err).ShouldNot(gomega.HaveOccurred())

			g.Expect(s.Size()).Should(gomega.Equal(tt.want.size))
			g.Expect(pop.(string)).Should(gomega.Equal(tt.want.popItem))
		})
	}
}

