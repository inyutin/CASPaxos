package main

type Pair struct {
	State int
	BallotNumber Ballot
}

type Acceptor struct {
	Promise Ballot
	Accepted Pair
}

func NewAcceptor() Acceptor {
	acceptor := new(Acceptor)
	acceptor.Promise  = Ballot {0, 0}
	acceptor.Accepted = Pair{0, Ballot{0, 0}}
	return *acceptor
}

func (acceptor Acceptor) prepare(ballotNumber Ballot) Pair {
	if acceptor.Promise.MoreThan(ballotNumber) {
		return Pair{-1, Ballot{-1, -1}}
	}

	acceptor.Promise = ballotNumber
	return acceptor.Accepted
}


func (acceptor *Acceptor) accept(ballotNumber Ballot, newState int) Pair {
	if acceptor.Promise.MoreThan(ballotNumber) || acceptor.Accepted.BallotNumber.MoreThan(ballotNumber) {
		return Pair{-1,Ballot{-1, -1}}
	}

	acceptor.Promise = Ballot {0, 0}
	acceptor.Accepted = Pair{newState, ballotNumber}
	return Pair{-2, Ballot{-2, -2}}
}
