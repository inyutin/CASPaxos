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

	write := func(x int) int {
		return 2
	}

	read := func(x int) int {
		return x
	}

	cas := func(x int) int {
		if x == 2 {
			return 3
		}
		return x
	}

	acceptorsList := []Acceptor{a1, a2, a3, a4, a5}
	p1 := NewProposer(acceptorsList,1)
	p2 := NewProposer(acceptorsList, 2)
	p3 := NewProposer(acceptorsList, 3)

	go func() {
		result := p1.receive(write)
		fmt.Println(result)
	}()

	go func() {
		result := p2.receive(read)
		fmt.Println(result)
	}()

	go func() {
		result := p3.receive(cas)
		fmt.Println(result)
	}()

	time.Sleep(60 * time.Second)
}