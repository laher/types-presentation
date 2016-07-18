package types

import "testing"

func TestNotifier(t *testing.T) {
	var d Notifier

	d = Outer{}
	t.Logf("Notifier.Notify() returns [%s]", d.Notify())
	d = Outer2{}
	t.Logf("Notifier.Notify() returns [%s]", d.Notify())

}
