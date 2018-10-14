package main

import (
	"math/rand"
	"sort"
	"time"
)

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max - min) + min
}

type Proposer struct {
	Acceptors []Acceptor
	F int
	state int
}

func NewProposer(acceptors []Acceptor) *Proposer {
	proposer := new(Proposer)
	proposer.Acceptors  = acceptors
	proposer.F = (len(acceptors) - 1) / 2
	proposer.state = 0
	return proposer
}

func (proposer *Proposer) receive(f func(x int) int) int {
	ballotNumber := proposer.generateBallotNumber()
	proposer.sendPrepare(ballotNumber)
	return proposer.sendAccept(f, ballotNumber)
}

func (proposer *Proposer) generateBallotNumber() int {
	return random(1, 100)
}


func (proposer *Proposer) sendPrepare(ballotNumber int) {
	conformations :=  make([]Pair, 0, 0)
	for _, acceptor := range proposer.Acceptors {
		conformation := acceptor.prepare(ballotNumber)
		if conformation.State != -1 {
			conformations = append(conformations, conformation)
		}
	}

	for true {
		if len(conformations) >= proposer.F + 1 {
			break
		} else {
			time.Sleep(5)
		}
	}
	sum := 0
	totalListOfConfirmationValues :=  make([]int, 0, 0)
	for _, conformation := range conformations {
		sum += conformation.State
		totalListOfConfirmationValues = append(totalListOfConfirmationValues, conformation.State)
	}

	if sum == 0 {
		proposer.state = 0
	} else {
		proposer.state = getHighestConfirmation(conformations).State
	}
}


func getHighestConfirmation(conformations []Pair) Pair {
	ballots := make([]int, 1, 1)
	for _, conformation := range conformations {
		ballots = append(ballots, conformation.ballotNumber)
	}
	sort.Ints(ballots)
	highestBallot := ballots[len(ballots) - 1]

	for _, conformation := range conformations {
		if conformation.ballotNumber == highestBallot {
			return conformation
		}
	}
	return conformations[0]
}

func (proposer *Proposer) sendAccept(f func(x int) int, ballotNumber int) int {
	proposer.state = f(proposer.state)
	acceptations := make([]Pair, 0, 0)
	for id, acceptor := range proposer.Acceptors {
		acceptation := acceptor.accept(ballotNumber, proposer.state)
		proposer.Acceptors[id] = acceptor
		if acceptation.State != -1 {
			acceptations = append(acceptations, acceptation)
		}
	}

	for true {
		if len(acceptations) >= proposer.F + 1 {
			break
		}
		time.Sleep(5)
	}
	return proposer.state
}
