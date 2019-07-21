package rps_test

import (
	"testing"

	rps "github.com/GodsBoss/rock-paper-scissors"
)

func TestPlayerNamesAreDifferentFromEachOther(t *testing.T) {
	players := []rps.Player{rps.NoPlayer, rps.PlayerOne, rps.PlayerTwo, rps.Player(9001)}
	for i := range players {
		for j := range players {
			if i != j && players[i].String() == players[j].String() {
				t.Errorf("expected players[%d] and players[%d] to have different names, but both are '%s'", i, j, players[i].String())
			}
		}
	}
}
