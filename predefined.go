package rps

// Classic are the classic rock/paper/scissors rules.
var Classic = NewRules().
	Add("rock", "scissors").
	Add("scissors", "paper").
	Add("paper", "rock")

// Extended are the classic rules plus well and match.
// The well beats rock and scissors and is beaten by match and paper.
// The match beats paper and well and is beaten by rock and scissors.
var Extended = NewRules().
	Add("rock", "scissors").
	Add("rock", "match").
	Add("scissors", "paper").
	Add("scissors", "match").
	Add("paper", "rock").
	Add("paper", "well").
	Add("match", "paper").
	Add("match", "well").
	Add("well", "rock").
	Add("well", "scissors")
