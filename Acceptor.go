package main

type Pair struct {
	State, BallotNumber int
}

type Acceptor struct {
	Promise int
	Accepted Pair
}

func NewAcceptor() Acceptor {
	acceptor := new(Acceptor)
	acceptor.Promise  = 0 // TODO: надо делать флаш на диск
	acceptor.Accepted = Pair{0, 0}
	return *acceptor
}

func (acceptor Acceptor) prepare(ballotNumber int) Pair {
	if acceptor.Promise > ballotNumber {
		return Pair{-1,-1}
	}

	acceptor.Promise = ballotNumber // TODO: надо делать флаш на диск
	return acceptor.Accepted
}


func (acceptor *Acceptor) accept(ballotNumber int, newState int) Pair {
	if acceptor.Promise > ballotNumber || acceptor.Accepted.BallotNumber > ballotNumber {
		return Pair{-1,-1}
	}

	// TODO: надо делать флаш на диск
	acceptor.Promise = 0
	acceptor.Accepted = Pair{newState, ballotNumber}
	return Pair{-2, -2}
}
