package main

import (
	"github.com/Workiva/go-datastructures/set"
	"math/rand"
	"sort"
	"time"
)

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max - min) + min
}

type Proposer struct {
	Acceptors *set.Set
	Quorum  int
	State   int
	BallotNumber 	Ballot
	Id int
}

func NewProposer(acceptors *set.Set, id int) *Proposer {
	proposer := new(Proposer)
	proposer.Acceptors  = acceptors
	proposer.Quorum = (int(acceptors.Len()) - 1) / 2
	proposer.State = 0
	proposer.BallotNumber = Ballot {1, id}
	proposer.Id = id
	return proposer
}

func (proposer *Proposer) receive(f func(x int) int) (int, bool) {
	proposer.BallotNumber = proposer.generateBallotNumber()
	if proposer.sendPrepare(proposer.BallotNumber) {
		return proposer.sendAccept(f, proposer.BallotNumber)
	} else {
		return 0, false
	}
}

func (proposer *Proposer) generateBallotNumber() Ballot {
	return Ballot {proposer.BallotNumber.Number + random(1, 100), proposer.BallotNumber.Id}
}


func (proposer *Proposer) sendPrepare(ballotNumber Ballot) bool{
	conformations :=  make([]Pair, 0, 0)

	// TODO: надо делать в разных поток отправку prepare и проверку
	for _, acceptor := range proposer.Acceptors.Flatten() {
		(Acceptor*) acceptor
		conformation := acceptor.prepare(ballotNumber)
		if conformation.State != -1 {
			conformations = append(conformations, conformation)
		}
	}

	for true {
		if len(conformations) >= proposer.Quorum+ 1 {
			break
		} else if len(conformations) >= proposer.Quorum+ 1 {
			return false
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
		proposer.State = 0
	} else {
		proposer.State = getHighestConfirmation(conformations).State
	}
	return true
}


func getHighestConfirmation(conformations []Pair) Pair {
	ballots := make([]Ballot, 1, 1)
	for _, conformation := range conformations {
		ballots = append(ballots, conformation.BallotNumber)
	}

	sort.Slice(ballots, func(i, j int) bool {
		return ballots[j].MoreThan(ballots[i])
	})

	highestBallot := ballots[len(ballots) - 1]

	for _, conformation := range conformations {
		if conformation.BallotNumber == highestBallot {
			return conformation
		}
	}
	return conformations[0]
}

func (proposer *Proposer) sendAccept(f func(x int) int, ballotNumber Ballot) (int, bool) {
	proposer.State = f(proposer.State)
	acceptations := make([]Pair, 0, 0)
	for id, acceptor := range proposer.Acceptors {
		acceptation := acceptor.accept(ballotNumber, proposer.State)
		proposer.Acceptors[id] = acceptor
		if acceptation.State != -1 {
			acceptations = append(acceptations, acceptation)
		}
	}

	if len(acceptations) >= proposer.Quorum+ 1 {
		return proposer.State, true
	}
	time.Sleep(5)

	return proposer.State, false
}
