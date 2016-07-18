package types

import "testing"

func TestInterfaces(t *testing.T) {
	var d Notifier

	d = Outer{}
	Log(d)
	t.Logf("Notifier.Notify() returns [%s]", d.Notify())
	d = Outer2{}
	Log(d)
	t.Logf("Notifier.Notify() returns [%s]", d.Notify())

}
