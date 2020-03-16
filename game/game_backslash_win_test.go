package game

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsWinningCombinationExistForCellBackSlash1(t *testing.T) {
	a := assert.New(t)
	userGame := Game{
		GameMap: [][]int{
			{1,0,0},
			{0,1,0},
			{0,0,1},
		},
		UserTurn:     false,
		GameFinished: false,
	}
	a.True(IsWinningCombinationExistForCell(&userGame,1,1))
}


func TestIsWinningCombinationExistForCellBackSlash2(t *testing.T) {
	a := assert.New(t)
	userGame := Game{
		GameMap: [][]int{
			{1,0,0},
			{0,1,0},
			{0,0,1},
		},
		UserTurn:     false,
		GameFinished: false,
	}
	a.True(IsWinningCombinationExistForCell(&userGame,1,1))
}


func TestIsWinningCombinationExistForCellBackSlash3(t *testing.T) {
	a := assert.New(t)
	userGame := Game{
		GameMap: [][]int{
			{0,1,0,0,0},
			{0,0,1,0,0},
			{0,0,0,1,0},
		},
		UserTurn: false,
		GameFinished: false,
	}
	a.True(IsWinningCombinationExistForCell(&userGame,0,1))
}


func TestIsWinningCombinationExistForCellBackSlash4(t *testing.T) {
	a := assert.New(t)
	userGame := Game{
		GameMap: [][]int{
			{0,1,0,0,0},
			{0,0,0,0,0},
			{0,0,0,1,0},
		},
		UserTurn: false,
		GameFinished: false,
	}
	a.False(IsWinningCombinationExistForCell(&userGame,0,1))
}


func TestIsWinningCombinationExistForCellBackSlash5(t *testing.T) {
	a := assert.New(t)
	userGame := Game{
		GameMap: [][]int{
			{0,0,0},
			{1,0,0},
			{0,1,0},
			{0,0,1},
			{0,0,0},
		},
		UserTurn: false,
		GameFinished: false,
	}
	a.True(IsWinningCombinationExistForCell(&userGame,1,0))
}


func TestIsWinningCombinationExistForCellBackSlash6(t *testing.T) {
	a := assert.New(t)
	userGame := Game{
		GameMap: [][]int{
			{0,0,0},
			{1,0,0},
			{0,1,0},
			{0,0,1},
		},
		UserTurn: false,
		GameFinished: false,
	}
	a.True(IsWinningCombinationExistForCell(&userGame,1,0))
}