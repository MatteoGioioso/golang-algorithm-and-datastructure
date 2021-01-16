package linkedList_test

import (
	"algo/datastructures/linkedList"
	"github.com/onsi/gomega"
	"testing"
)

func Test_linkedList_Prepend(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	tests := []struct {
		name string
		args []string
	}{
		{name: "should insert items", args: []string{"one", "two", "three"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := linkedList.New()

			for _, arg := range tt.args {
				l.Prepend(arg)
			}

			g.Expect(l.Size()).Should(gomega.Equal(len(tt.args)))
			l.Print()
		})
	}
}

func Test_linkedList_Append(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	tests := []struct {
		name   string
		args   []string
	}{
		{name: "should append three values", args: []string{"one", "two", "three"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := linkedList.New()

			for _, arg := range tt.args {
				l.Append(arg)
			}

			g.Expect(l.Size()).Should(gomega.Equal(len(tt.args)))

			l.Print()
		})
	}
}

func Test_linkedList_GetAt(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	type args struct {
		position int
	}

	tests := []struct {
		name   string
		args  args
		want string
	}{
		{name: "should find the right value", args: args{position: 1}, want: "two"},
		{name: "should find the right value", args: args{position: 0}, want: "one"},
		{name: "should find nothing", args: args{position: 4}, want: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := linkedList.New()

			for _, arg := range []string{"one", "two", "three"} {
				l.Append(arg)
			}

			at := l.GetAt(tt.args.position)

			if at != nil {
				g.Expect(at).To(gomega.Equal(tt.want))
			} else {
				at = ""
				g.Expect(at).To(gomega.Equal(tt.want))
			}
		})
	}
}
