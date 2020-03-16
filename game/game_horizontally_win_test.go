package game

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsWinningCombinationExistForCellHorizontal1(t *testing.T) {
	a := assert.New(t)
	userGame := Game{
		GameMap:      [][]int{
			{1,1,1,1,1},
			{0,0,2,0,0},
			{0,0,2,0,0},
			{0,0,2,0,0},
			{0,0,2,0,0},
		},
		UserTurn:     false,
		GameFinished: false,
	}
	a.True(IsWinningCombinationExistForCell(&userGame,0,0))
}

func TestIsWinningCombinationExistForCellHorizontal2(t *testing.T) {
	a := assert.New(t)
	userGame := Game{
		GameMap:      [][]int{
			{0,0,0,0,0},
			{1,0,2,0,0},
			{1,1,1,1,1},
			{2,0,2,0,0},
			{2,0,2,0,0},
		},
		UserTurn:     false,
		GameFinished: false,
	}
	a.True(IsWinningCombinationExistForCell(&userGame,2,0))
}

func TestIsWinningCombinationExistForCellHorizontal3(t *testing.T) {
	a := assert.New(t)
	userGame := Game{
		GameMap:      [][]int{
			{0,0,0,0,0},
			{1,0,2,0,0},
			{0,0,0,0,0},
			{2,0,2,0,0},
			{1,1,1,1,1},
		},
		UserTurn:     false,
		GameFinished: false,
	}
	a.True(IsWinningCombinationExistForCell(&userGame,4,0))
}

func TestIsWinningCombinationExistForCellHorizontal4(t *testing.T) {
	a := assert.New(t)
	userGame := Game{
		GameMap:      [][]int{
			{0,0,0,0,0},
			{1,0,2,0,0},
			{0,0,0,0,0},
			{2,0,2,0,0},
			{1,1,1,0,1},
		},
		UserTurn:     false,
		GameFinished: false,
	}
	a.False(IsWinningCombinationExistForCell(&userGame,4,0))
}






