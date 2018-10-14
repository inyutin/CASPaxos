package main

import (
	"fmt"
	"time"
)

func main() {
	a1 := NewAcceptor()
	a2 := NewAcceptor()
	a3 := NewAcceptor()
	a4 := NewAcceptor()
	a5 := NewAcceptor()

	change := func(x int) int{
		return x + 4
	}

	acceptorsList := []Acceptor{a1, a2, a3, a4, a5}
	p1 := NewProposer(acceptorsList)
	p2 := NewProposer(acceptorsList)
	p3 := NewProposer(acceptorsList)

	go func() {
		result := p1.receive(change)
		fmt.Println(result)
	}()

	go func() {
		result := p2.receive(change)
		fmt.Println(result)
	}()

	go func() {
		result := p3.receive(change)
		fmt.Println(result)
	}()

	time.Sleep(60 * time.Second)
}