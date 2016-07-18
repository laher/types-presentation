// examples
// https://engo.io/tutorials/02-first-system
//
package types

type Embeddable struct {
}

func (e Embeddable) Notify() string {
	return "Embeddable"
}

type Outer struct {
	Embeddable
	Other int
}

type Outer2 struct {
	Embeddable
	Other int
}

func (o Outer2) Notify() string {
	return "Outer2"
}
