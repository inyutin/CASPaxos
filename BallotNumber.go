package main

type Ballot struct {
	Number, Id int
}

func (first *Ballot) MoreThan(second Ballot) bool {
	if first.Number > second.Number {
		return true
	} else if first.Number == second.Number {
		if first.Id > second.Id {
			return true
		}
	}
	return false
}