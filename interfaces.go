package types

import "fmt"

type Notifier interface {
	Notify() string
}

func Log(n Notifier) {
	fmt.Println(n.Notify())
}
