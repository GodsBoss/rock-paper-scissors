package rps_test

import (
	"testing"

	rps "github.com/GodsBoss/rock-paper-scissors"
)

func TestClassicRockPaperScissors(t *testing.T) {
	type testCase struct {
		p1             string
		p2             string
		expectedWinner rps.Player
	}
	testCases := map[string]testCase{
		"invalid player one choice leads to draw": {
			p1:             "lizard",
			p2:             "rock",
			expectedWinner: rps.NoPlayer,
		},
		"invalid player two choice leads to draw": {
			p1:             "paper",
			p2:             "wizard",
			expectedWinner: rps.NoPlayer,
		},
		"equality leads to draw": {
			p1:             "scissors",
			p2:             "scissors",
			expectedWinner: rps.NoPlayer,
		},
		"player one wins": {
			p1:             "paper",
			p2:             "rock",
			expectedWinner: rps.PlayerOne,
		},
		"player two wins": {
			p1:             "scissors",
			p2:             "rock",
			expectedWinner: rps.PlayerTwo,
		},
	}

	for name := range testCases {
		t.Run(
			name,
			func(tc testCase) func(*testing.T) {
				return func(t *testing.T) {
					winner := rps.Classic.Winner(rps.Choice(tc.p1), rps.Choice(tc.p2))
					if winner != tc.expectedWinner {
						t.Errorf("expected %s to win, but %s won", tc.expectedWinner, winner)
					}
				}
			}(testCases[name]),
		)
	}
}
