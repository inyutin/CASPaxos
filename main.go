package main

import (
	"fmt"
	"github.com/Workiva/go-datastructures/set"
	"time"
)

func main() {
	a1 := NewAcceptor()
	a2 := NewAcceptor()
	a3 := NewAcceptor()
	a4 := NewAcceptor()
	a5 := NewAcceptor()

	write := func(x int) int {
		return 2
	}

	inc := func(x int) int {
		return x+1
	}

	acceptorsList := set.Set{}
	acceptorsList.Add(a1, a2, a3, a4, a5)
	p1 := NewProposer(acceptorsList,1)
	p2 := NewProposer(acceptorsList, 2)
	p3 := NewProposer(acceptorsList, 3)

	fmt.Println("Init state: 0")
	go func() {
		result, state := p1.receive(write)
		fmt.Printf("Id: %d, Result: %d, State: %t\n", p1.Id, result, state)
	}()

	go func() {
		result, state := p2.receive(inc)
		fmt.Printf("Id: %d, Result: %d, State: %t\n", p2.Id, result, state)
	}()

	go func() {
		result, state := p3.receive(inc)
		fmt.Printf("Id: %d, Result: %d, State: %t\n", p3.Id, result, state)
	}()

	time.Sleep(60 * time.Second)
}