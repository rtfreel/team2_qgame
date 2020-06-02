package models

import (
	"fmt"
	"testing"
)

func TestLocation_Interact(t *testing.T) {
	var tests = []struct {
		player       Player
		location     Location
		wantX, wantY int
	}{
		{
			Player{2, 4},
			&EmptyField{2, 5},
			2,
			5,
		},
		{
			Player{10, 5},
			&EmptyField{11, 5},
			11,
			5,
		},
	}

	for _, tt := range tests {
		x, y := tt.location.GetLocation()
		testname := fmt.Sprintf("Player(%d,%d) moves to EmptyField(%d,%d)", tt.player.X, tt.player.Y, x, y)
		t.Run(testname, func(t *testing.T) {
			tt.player.InteractWith(&tt.location)
			if x != tt.wantX || y != tt.wantY {
				t.Errorf("got Player(%d,%d), want Player(%d,%d)", x, y, tt.wantX, tt.wantY)
			}
		})
	}
}
