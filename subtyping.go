package types

import "fmt"

type Subtype Outer2

func (_ Subtype) Notify() string {
	return "Subtype"
}

type IntSubtype int

func (i IntSubtype) Notify() string {
	return fmt.Sprintf("Subtype:[%d]", i)
}

type SliceSubtype []string

func (i SliceSubtype) Notify() string {
	return fmt.Sprintf("SliceSubtype:%v", i)
}
