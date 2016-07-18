package types

import "testing"

func TestNotifyer(t *testing.T) {
	var d Notifier

	d = Outer{}
	t.Logf("Notifyer.Notify() returns [%s]", d.Notify())
	d = Outer2{}
	t.Logf("Notifyer.Notify() returns [%s]", d.Notify())

}
