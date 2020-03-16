package game

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsWinningCombinationExistForCellSlash1(t *testing.T) {
	a := assert.New(t)
	userGame := Game{
		GameMap: [][]int{
			{0,0,1},
			{0,1,0},
			{1,0,0},
		},
		UserTurn:     false,
		GameFinished: false,
	}
	a.True(IsWinningCombinationExistForCell(&userGame,1,1))
}


func TestIsWinningCombinationExistForCellSlash2(t *testing.T) {
	a := assert.New(t)
	userGame := Game{
		GameMap: [][]int{
			{0,0,0,1,0},
			{0,0,1,0,0},
			{0,1,0,0,0},
		},
		UserTurn:     false,
		GameFinished: false,
	}
	a.True(IsWinningCombinationExistForCell(&userGame,1,2))
}


func TestIsWinningCombinationExistForCellSlash3(t *testing.T) {
	a := assert.New(t)
	userGame := Game{
		GameMap: [][]int{
			{0,0,0,1},
			{0,0,1,0},
			{0,1,0,0},
		},
		UserTurn:     false,
		GameFinished: false,
	}
	a.True(IsWinningCombinationExistForCell(&userGame,1,2))
}


func TestIsWinningCombinationExistForCellSlash4(t *testing.T) {
	a := assert.New(t)
	userGame := Game{
		GameMap: [][]int{
			{0,0,0,0},
			{0,0,0,1},
			{0,0,1,0},
			{0,1,0,0},
			{1,0,0,0},
			{0,0,0,0},
		},
		UserTurn:     false,
		GameFinished: false,
	}
	a.True(IsWinningCombinationExistForCell(&userGame,1,3))
}


func TestIsWinningCombinationExistForCellSlash5(t *testing.T) {
	a := assert.New(t)
	userGame := Game{
		GameMap: [][]int{
			{0,0,0,0},
			{0,0,0,0},
			{0,0,0,1},
			{0,0,1,0},
			{0,1,0,0},
			{1,0,0,0},
			{0,0,0,0},
		},
		UserTurn:     false,
		GameFinished: false,
	}
	a.True(IsWinningCombinationExistForCell(&userGame,2,3))
}
