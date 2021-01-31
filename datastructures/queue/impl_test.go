package queue_test

import (
	"algo/datastructures/queue"
	"github.com/onsi/gomega"
	"testing"
)

func Test_stack_Queue(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	type args struct {
		value interface{}
	}
	type want struct {
		size     int
		peekElem string
	}

	tests := []struct {
		name    string
		args    args
		want    want
		wantErr bool
	}{
		{name: "should push items into the stack", want: want{
			size:     3,
			peekElem: "one",
		}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := queue.New()

			s.Enqueuing("one")
			s.Enqueuing("two")
			s.Enqueuing("three")

			g.Expect(s.Size()).Should(gomega.Equal(tt.want.size))
			g.Expect(s.Peeking()).Should(gomega.Equal(tt.want.peekElem))
		})
	}
}

func Test_stack_Pop(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	type args struct {
		value interface{}
	}
	type want struct {
		size      int
		firstItem string
	}

	tests := []struct {
		name    string
		args    args
		want    want
		wantErr bool
	}{
		{name: "should push items into the stack", want: want{
			size:      3,
			firstItem: "one",
		}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := queue.New()

			s.Enqueuing("one")
			s.Enqueuing("two")
			s.Enqueuing("three")
			s.Enqueuing("four")

			g.Expect(s.Peeking()).Should(gomega.Equal("one"))

			item, err := s.Dequeuing()
			g.Expect(err).ShouldNot(gomega.HaveOccurred())

			g.Expect(s.Size()).Should(gomega.Equal(tt.want.size))
			g.Expect(item).Should(gomega.Equal(tt.want.firstItem))
			g.Expect(s.Peeking()).Should(gomega.Equal("two"))
		})
	}
}
