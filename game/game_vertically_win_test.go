package game

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsWinningCombinationExistForCellVertical1(t *testing.T) {
	t.Parallel()
	a := assert.New(t)
	userGame := Game{
		GameMap:      [][]int{
			{0,1,0},
			{0,1,0},
			{0,1,0},
		},
		UserTurn:     false,
		GameFinished: false,
	}
	a.True(IsWinningCombinationExistForCell(&userGame,2,1))
}

func TestIsWinningCombinationExistForCellVertical2(t *testing.T) {
	t.Parallel()
	a := assert.New(t)
	userGame := Game{
		GameMap:      [][]int{
			{0,1,0,0},
			{0,1,0,0},
			{0,1,0,0},
			{0,1,0,0},
		},
		UserTurn:     false,
		GameFinished: false,
	}
	a.True(IsWinningCombinationExistForCell(&userGame,2,1))
}

func TestIsWinningCombinationExistForCellVertical3(t *testing.T) {
	t.Parallel()
	a := assert.New(t)
	userGame := Game{
		GameMap:      [][]int{
			{0,0,0,0,1},
			{0,0,0,0,1},
			{0,0,0,0,1},
			{0,0,0,0,1},
			{0,0,0,0,1},
		},
		UserTurn:     false,
		GameFinished: false,
	}
	a.True(IsWinningCombinationExistForCell(&userGame,2,4))
}

func TestIsWinningCombinationExistForCellVertical4(t *testing.T) {
	t.Parallel()
	a := assert.New(t)
	userGame := Game{
		GameMap:      [][]int{
			{1,0,0,0,0},
			{1,0,2,0,0},
			{1,0,2,0,0},
			{1,0,2,0,0},
			{1,0,2,0,0},
		},
		UserTurn:     false,
		GameFinished: false,
	}
	a.True(IsWinningCombinationExistForCell(&userGame,3,0))
}

func TestIsWinningCombinationExistForCellVertical5(t *testing.T) {
	t.Parallel()
	a := assert.New(t)
	userGame := Game{
		GameMap:      [][]int{
			{1,0,0,0,0},
			{1,0,2,0,0},
			{2,0,2,0,0},
			{1,0,2,0,0},
			{1,0,2,0,0},
		},
		UserTurn:     false,
		GameFinished: false,
	}
	a.False(IsWinningCombinationExistForCell(&userGame,3,0))
}

func TestIsWinningCombinationExistForCellVertical6(t *testing.T) {
	t.Parallel()
	a := assert.New(t)
	userGame := Game{
		GameMap:      [][]int{
			{1,0,0,0,0},
			{1,0,2,0,0},
			{1,0,2,0,0},
			{2,0,2,0,0},
			{1,0,2,0,0},
		},
		UserTurn:     false,
		GameFinished: false,
	}
	a.False(IsWinningCombinationExistForCell(&userGame,3,0))
}






