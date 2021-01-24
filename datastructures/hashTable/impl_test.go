package hashTable_test

import (
	"algo/datastructures/hashTable"
	"fmt"
	"github.com/onsi/gomega"
	"testing"
)

func Test_hashTable_Set(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	type args struct {
		key   string
		value interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "should insert an item", args: args{
			key:   "itemOne",
			value: "my item",
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := hashTable.New()
			set := h.Set(tt.args.key, tt.args.value)

			g.Expect(set).Should(gomega.Equal(true))
			g.Expect(h.Size()).Should(gomega.Equal(1))

			value := h.Get(tt.args.key)
			g.Expect(value).Should(gomega.Equal(tt.args.value))
		})
	}
}

func Test_hashTable_Delete(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "should delete an item that exists", args: args{key: "itemTwo"}},
		{name: "should delete an item that does not exists", args: args{key: "itemFour"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := hashTable.New()

			h.Set("itemOne", "my item one")
			h.Set("itemTwo", "my item two")
			h.Set("itemThree", "my item three")

			g.Expect(h.Size()).Should(gomega.Equal(3))

			b := h.Delete(tt.args.key)

			g.Expect(b).Should(gomega.Equal(true))
			g.Expect(h.Size()).Should(gomega.Equal(2))

			g.Expect(h.Get(tt.args.key)).Should(gomega.BeNil())
		})
	}
}

func Test_hashTable_Load(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	type args struct {
		key      string
		fillSize int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "should load the hash table with many elements and still operate normally",
			args: args{key: "item19", fillSize: 20},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := hashTable.New()

			for i := 0; i < tt.args.fillSize; i++ {
				h.Set(fmt.Sprintf("item%d", i), fmt.Sprintf("my item %v", i))
			}

			g.Expect(h.Size()).Should(gomega.Equal(tt.args.fillSize))

			for i := 0; i < tt.args.fillSize; i++ {
				g.Expect(h.Get(fmt.Sprintf("item%d", i))).
					Should(gomega.Equal(fmt.Sprintf("my item %v", i)))
			}

			b := h.Delete(tt.args.key)

			g.Expect(b).Should(gomega.Equal(true))
			g.Expect(h.Size()).Should(gomega.Equal(tt.args.fillSize - 1))

			g.Expect(h.Get(tt.args.key)).Should(gomega.BeNil())
		})
	}
}
