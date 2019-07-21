package rps

// Choice is the choice a player makes.
type Choice string

// Rules represent rock/paper/scissors-style rules.
type Rules map[Choice]map[Choice]bool

// NewRules creates new, empty rules.
func NewRules() Rules {
	return make(Rules)
}

// Add adds a rule for winning/losing. If a player chooses the winning choice and
// the other player chooses the losing one, the former player wins.
func (rules Rules) Add(winning, losing Choice) Rules {
	if rules[winning] == nil {
		rules[winning] = make(map[Choice]bool)
	}
	if rules[losing] == nil {
		rules[losing] = make(map[Choice]bool)
	}
	rules[winning][losing] = true
	rules[losing][winning] = false
	return rules
}

// Winner determines the winner of the player's choices according to the rules.
// Unknown choices lead to a draw.
func (rules Rules) Winner(playerOneChoice, playerTwoChoice Choice) Player {
	if rules[playerOneChoice] == nil {
		return NoPlayer
	}
	result, found := rules[playerOneChoice][playerTwoChoice]
	if !found {
		return NoPlayer
	}
	if result {
		return PlayerOne
	}
	return PlayerTwo
}

// Player is a player, returned by Rules.Winner().
type Player int

const (
	// NoPlayer is returned by Rules.Winner if the game ends in a draw.
	NoPlayer Player = iota

	// PlayerOne is returned by Rules.Winner if the first player won.
	PlayerOne

	// PlayerTwo is returned by Rules.Winner if the second player won.
	PlayerTwo
)

// String formats player to a string.
func (player Player) String() string {
	switch player {
	case NoPlayer:
		return "none"
	case PlayerOne:
		return "player one"
	case PlayerTwo:
		return "player two"
	default:
		return "unknown"
	}
}
