package rete

import (
	"testing"
	"container/list"
)

func Test_constant_test_node_activation(t *testing.T)  {
	mem := &AlphaMemory{
		items: list.New(),
		successors: list.New(),
	}
	node := ConstantTestNode{
		field_to_test: IDENTIFIER,
		field_must_equal: "B1",
		output_memory: mem,
		children: list.New(),
	}
	w0 := &WME{
		fields: [3]string{"B1", "color", "red"},
		alpha_mems: list.New(),
	}
	w1 := &WME{
		fields: [3]string{"B2", "color", "red"},
		alpha_mems: list.New(),
	}

	node.activation(w0)
	node.activation(w1)

	if contain(mem.items, w0) == nil {
		t.Error("w0 not in alpha memory")
	}
	if contain(mem.items, w1) != nil {
		t.Error("w1 in alpha memory")
	}
	RemoveWme(w0)
	if mem.items.Len() != 0 {
		t.Error("alpha memory not empty")
	}
}
