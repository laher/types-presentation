package types

import "testing"

func TestSubtyping(t *testing.T) {
	subtype := Subtype{}
	t.Logf("Subtype.Notify() returns [%s]", subtype.Notify())
	//does not compile:
	//t.Logf("Subtype.Outer2.Notify() returns [%s]", subtype.Outer2.Notify())
	intSubType := IntSubtype(3)
	t.Logf("IntSubtype.Notify() returns [%s]", intSubType.Notify())
	sliceSubtype := SliceSubtype{"a", "b"}
	t.Logf("SliceSubtype.Notify() returns [%s]", sliceSubtype.Notify())
}
