package main

import (
	"fmt"

	"github.com/corymurphy/adventofcode/shared"
)

type Monkey struct {
	Id             int
	Inspected      int
	Items          *shared.IntQueue
	ThrowTest      int
	IfFalse        int
	IfTrue         int
	WorryModifier  int
	WorryOperation Operation
}

func (m *Monkey) ItemsToString() string {

	item, err := m.Items.Dequeue()
	items := ""

	for err == nil {
		if items == "" {
			items = fmt.Sprint(item)
		} else {
			items = items + " " + fmt.Sprint(item)
		}
		item, err = m.Items.Dequeue()
	}
	return items
}
