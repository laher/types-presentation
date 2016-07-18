package types

import "testing"

func TestEmbedding(t *testing.T) {
	outer := Outer{}
	t.Logf("Outer.Notify() returns [%s]", outer.Notify())

	outer2 := Outer2{}
	t.Logf("Outer2.Notify() returns [%s]", outer2.Notify())
	t.Logf("Outer2.Embeddable.Notify() returns [%s]", outer2.Embeddable.Notify())
}
