package game

import (
	"fmt"
	"testing"

	"github.com/GrigoryKrasnochub/go-tic-tac-toe/game/gmmap"
	"github.com/stretchr/testify/assert"
)

type userGames struct {
	userGame       Game
	gameMap        [][]gmmap.CellValue
	cellCoordinate [2]int
	result         bool
}

func funcToTestsIsWinningCombinationExistForCell(t *testing.T, winCombinations []userGames) {
	a := assert.New(t)
	for i, testGameWin := range winCombinations {
		testGameWin.userGame.LoadGameMapFromSlice(testGameWin.gameMap)
		if testGameWin.result {
			a.True(testGameWin.userGame.IsWinningCombinationExistForCell(testGameWin.cellCoordinate[0], testGameWin.cellCoordinate[1]), fmt.Sprintf("%s %d, want %t, got %t", "Test Data №", i, true, false))
		} else {
			a.False(testGameWin.userGame.IsWinningCombinationExistForCell(testGameWin.cellCoordinate[0], testGameWin.cellCoordinate[1]), fmt.Sprintf("%s %d, want %t, got %t", "Test Data №", i, false, true))
		}
	}
}

func TestGame_IsWinningCombinationExistForCellBackSlash(t *testing.T) {
	testGameWins := []userGames{
		{
			Game{GameMapInterface: &gmmap.InMemoryGameMap{},
				UserTurn:     false,
				GameFinished: false,
			},
			[][]gmmap.CellValue{
				{gmmap.CrossCell, gmmap.EmptyCell, gmmap.EmptyCell},
				{gmmap.EmptyCell, gmmap.CrossCell, gmmap.EmptyCell},
				{gmmap.EmptyCell, gmmap.EmptyCell, gmmap.CrossCell},
			},
			[2]int{0, 0},
			true,
		},
		{
			Game{GameMapInterface: &gmmap.InMemoryGameMap{},
				UserTurn:     false,
				GameFinished: false,},
			[][]gmmap.CellValue{
				{gmmap.EmptyCell, gmmap.CrossCell, gmmap.EmptyCell, gmmap.EmptyCell, gmmap.EmptyCell},
				{gmmap.EmptyCell, gmmap.EmptyCell, gmmap.CrossCell, gmmap.EmptyCell, gmmap.EmptyCell},
				{gmmap.EmptyCell, gmmap.EmptyCell, gmmap.EmptyCell, gmmap.CrossCell, gmmap.EmptyCell},
			},
			[2]int{0, 1},
			true,
		},
		{
			Game{GameMapInterface: &gmmap.InMemoryGameMap{},
				UserTurn:     false,
				GameFinished: false,
			},
			[][]gmmap.CellValue{
				{gmmap.EmptyCell, gmmap.EmptyCell, gmmap.EmptyCell},
				{gmmap.CrossCell, gmmap.EmptyCell, gmmap.EmptyCell},
				{gmmap.EmptyCell, gmmap.CrossCell, gmmap.EmptyCell},
				{gmmap.EmptyCell, gmmap.EmptyCell, gmmap.CrossCell},
				{gmmap.EmptyCell, gmmap.EmptyCell, gmmap.EmptyCell},
			},
			[2]int{1, 0},
			true,
		},
		{
			Game{GameMapInterface: &gmmap.InMemoryGameMap{},
				UserTurn:     false,
				GameFinished: false,
			},
			[][]gmmap.CellValue{
				{gmmap.EmptyCell, gmmap.EmptyCell, gmmap.EmptyCell},
				{gmmap.CrossCell, gmmap.EmptyCell, gmmap.EmptyCell},
				{gmmap.EmptyCell, gmmap.CrossCell, gmmap.EmptyCell},
				{gmmap.EmptyCell, gmmap.EmptyCell, gmmap.CrossCell},
				{gmmap.EmptyCell, gmmap.EmptyCell, gmmap.EmptyCell},
			},
			[2]int{1, 0},
			true,
		},
		{
			Game{GameMapInterface: &gmmap.InMemoryGameMap{},
				UserTurn:     false,
				GameFinished: false,},
			[][]gmmap.CellValue{
				{gmmap.EmptyCell, gmmap.EmptyCell, gmmap.EmptyCell},
				{gmmap.CrossCell, gmmap.EmptyCell, gmmap.EmptyCell},
				{gmmap.EmptyCell, gmmap.CrossCell, gmmap.EmptyCell},
				{gmmap.EmptyCell, gmmap.EmptyCell, gmmap.CrossCell},
			},
			[2]int{1, 0},
			true,
		},
		{
			Game{GameMapInterface: &gmmap.InMemoryGameMap{},
				UserTurn:     false,
				GameFinished: false,
			},
			[][]gmmap.CellValue{
				{gmmap.EmptyCell, gmmap.CrossCell, gmmap.EmptyCell, gmmap.EmptyCell, gmmap.EmptyCell},
				{gmmap.EmptyCell, gmmap.EmptyCell, gmmap.EmptyCell, gmmap.EmptyCell, gmmap.EmptyCell},
				{gmmap.EmptyCell, gmmap.EmptyCell, gmmap.EmptyCell, gmmap.CrossCell, gmmap.EmptyCell},
			},
			[2]int{0, 1},
			false,
		},
	}
	funcToTestsIsWinningCombinationExistForCell(t, testGameWins)
}

func TestGame_IsWinningCombinationExistForCellSlash(t *testing.T) {
	testGameWins := []userGames{
		{
			Game{GameMapInterface: &gmmap.InMemoryGameMap{},
				UserTurn:     false,
				GameFinished: false,
			}, [][]gmmap.CellValue{
			{gmmap.EmptyCell, gmmap.EmptyCell, gmmap.CrossCell},
			{gmmap.EmptyCell, gmmap.CrossCell, gmmap.EmptyCell},
			{gmmap.CrossCell, gmmap.EmptyCell, gmmap.EmptyCell},
		}, [2]int{1, 1},
			true,
		},
		{
			Game{GameMapInterface: &gmmap.InMemoryGameMap{},
				UserTurn:     false,
				GameFinished: false,
			},
			[][]gmmap.CellValue{
				{gmmap.EmptyCell, gmmap.EmptyCell, gmmap.EmptyCell, gmmap.CrossCell, gmmap.EmptyCell},
				{gmmap.EmptyCell, gmmap.EmptyCell, gmmap.CrossCell, gmmap.EmptyCell, gmmap.EmptyCell},
				{gmmap.EmptyCell, gmmap.CrossCell, gmmap.EmptyCell, gmmap.EmptyCell, gmmap.EmptyCell},
			},
			[2]int{1, 2},
			true,
		},
		{
			Game{GameMapInterface: &gmmap.InMemoryGameMap{},
				UserTurn:     false,
				GameFinished: false,
			},
			[][]gmmap.CellValue{
				{gmmap.EmptyCell, gmmap.EmptyCell, gmmap.EmptyCell, gmmap.CrossCell},
				{gmmap.EmptyCell, gmmap.EmptyCell, gmmap.CrossCell, gmmap.EmptyCell},
				{gmmap.EmptyCell, gmmap.CrossCell, gmmap.EmptyCell, gmmap.EmptyCell},
			},
			[2]int{1, 2},
			true,
		},
		{
			Game{GameMapInterface: &gmmap.InMemoryGameMap{},
				UserTurn:     false,
				GameFinished: false,
			},
			[][]gmmap.CellValue{
				{gmmap.EmptyCell, gmmap.EmptyCell, gmmap.EmptyCell, gmmap.EmptyCell},
				{gmmap.EmptyCell, gmmap.EmptyCell, gmmap.EmptyCell, gmmap.CrossCell},
				{gmmap.EmptyCell, gmmap.EmptyCell, gmmap.CrossCell, gmmap.EmptyCell},
				{gmmap.EmptyCell, gmmap.CrossCell, gmmap.EmptyCell, gmmap.EmptyCell},
				{gmmap.CrossCell, gmmap.EmptyCell, gmmap.EmptyCell, gmmap.EmptyCell},
				{gmmap.EmptyCell, gmmap.EmptyCell, gmmap.EmptyCell, gmmap.EmptyCell},
			},
			[2]int{1, 3},
			true,
		},
		{
			Game{GameMapInterface: &gmmap.InMemoryGameMap{},
				UserTurn:     false,
				GameFinished: false,
			},
			[][]gmmap.CellValue{
				{gmmap.EmptyCell, gmmap.EmptyCell, gmmap.EmptyCell, gmmap.EmptyCell},
				{gmmap.EmptyCell, gmmap.EmptyCell, gmmap.EmptyCell, gmmap.EmptyCell},
				{gmmap.EmptyCell, gmmap.EmptyCell, gmmap.EmptyCell, gmmap.CrossCell},
				{gmmap.EmptyCell, gmmap.EmptyCell, gmmap.CrossCell, gmmap.EmptyCell},
				{gmmap.EmptyCell, gmmap.CrossCell, gmmap.EmptyCell, gmmap.EmptyCell},
				{gmmap.CrossCell, gmmap.EmptyCell, gmmap.EmptyCell, gmmap.EmptyCell},
				{gmmap.EmptyCell, gmmap.EmptyCell, gmmap.EmptyCell, gmmap.EmptyCell},
			},
			[2]int{2, 3},
			true,
		},
	}
	funcToTestsIsWinningCombinationExistForCell(t, testGameWins)
}

func TestGame_IsWinningCombinationExistForCellVertical(t *testing.T) {
	testGameWins := []userGames{
		{
			Game{GameMapInterface: &gmmap.InMemoryGameMap{},
				UserTurn:     false,
				GameFinished: false,
			},
			[][]gmmap.CellValue{
				{gmmap.EmptyCell, gmmap.CrossCell, gmmap.EmptyCell},
				{gmmap.EmptyCell, gmmap.CrossCell, gmmap.EmptyCell},
				{gmmap.EmptyCell, gmmap.CrossCell, gmmap.EmptyCell},
			},
			[2]int{2, 1},
			true,
		},
		{
			Game{GameMapInterface: &gmmap.InMemoryGameMap{},
				UserTurn:     false,
				GameFinished: false,
			},
			[][]gmmap.CellValue{
				{gmmap.EmptyCell, gmmap.CrossCell, gmmap.EmptyCell, gmmap.EmptyCell},
				{gmmap.EmptyCell, gmmap.CrossCell, gmmap.EmptyCell, gmmap.EmptyCell},
				{gmmap.EmptyCell, gmmap.CrossCell, gmmap.EmptyCell, gmmap.EmptyCell},
				{gmmap.EmptyCell, gmmap.CrossCell, gmmap.EmptyCell, gmmap.EmptyCell},
			},
			[2]int{2, 1},
			true,
		},
		{
			Game{GameMapInterface: &gmmap.InMemoryGameMap{},
				UserTurn:     false,
				GameFinished: false,
			},
			[][]gmmap.CellValue{
				{gmmap.EmptyCell, gmmap.EmptyCell, gmmap.EmptyCell, gmmap.EmptyCell, gmmap.CrossCell},
				{gmmap.EmptyCell, gmmap.EmptyCell, gmmap.EmptyCell, gmmap.EmptyCell, gmmap.CrossCell},
				{gmmap.EmptyCell, gmmap.EmptyCell, gmmap.EmptyCell, gmmap.EmptyCell, gmmap.CrossCell},
				{gmmap.EmptyCell, gmmap.EmptyCell, gmmap.EmptyCell, gmmap.EmptyCell, gmmap.CrossCell},
				{gmmap.EmptyCell, gmmap.EmptyCell, gmmap.EmptyCell, gmmap.EmptyCell, gmmap.CrossCell},
			},
			[2]int{2, 4},
			true,
		},
		{
			Game{GameMapInterface: &gmmap.InMemoryGameMap{},
				UserTurn:     false,
				GameFinished: false,
			},
			[][]gmmap.CellValue{
				{gmmap.CrossCell, gmmap.EmptyCell, gmmap.EmptyCell, gmmap.EmptyCell, gmmap.EmptyCell},
				{gmmap.CrossCell, gmmap.EmptyCell, gmmap.CircleCell, gmmap.EmptyCell, gmmap.EmptyCell},
				{gmmap.CrossCell, gmmap.EmptyCell, gmmap.CircleCell, gmmap.EmptyCell, gmmap.EmptyCell},
				{gmmap.CrossCell, gmmap.EmptyCell, gmmap.CircleCell, gmmap.EmptyCell, gmmap.EmptyCell},
				{gmmap.CrossCell, gmmap.EmptyCell, gmmap.CircleCell, gmmap.EmptyCell, gmmap.EmptyCell},
			},
			[2]int{3, 0},
			true,
		},
		{
			Game{GameMapInterface: &gmmap.InMemoryGameMap{},
				UserTurn:     false,
				GameFinished: false,
			},
			[][]gmmap.CellValue{
				{gmmap.CrossCell, gmmap.EmptyCell, gmmap.EmptyCell, gmmap.EmptyCell, gmmap.EmptyCell},
				{gmmap.CrossCell, gmmap.EmptyCell, gmmap.CircleCell, gmmap.EmptyCell, gmmap.EmptyCell},
				{gmmap.CrossCell, gmmap.EmptyCell, gmmap.CircleCell, gmmap.EmptyCell, gmmap.EmptyCell},
				{gmmap.CrossCell, gmmap.EmptyCell, gmmap.CircleCell, gmmap.EmptyCell, gmmap.EmptyCell},
				{gmmap.CrossCell, gmmap.EmptyCell, gmmap.CircleCell, gmmap.EmptyCell, gmmap.EmptyCell},
			},
			[2]int{3, 0},
			true,
		},
		{
			Game{GameMapInterface: &gmmap.InMemoryGameMap{},
				UserTurn:     false,
				GameFinished: false,
			},
			[][]gmmap.CellValue{
				{gmmap.CrossCell, gmmap.EmptyCell, gmmap.EmptyCell, gmmap.EmptyCell, gmmap.EmptyCell},
				{gmmap.CrossCell, gmmap.EmptyCell, gmmap.CircleCell, gmmap.EmptyCell, gmmap.EmptyCell},
				{gmmap.CircleCell, gmmap.EmptyCell, gmmap.CircleCell, gmmap.EmptyCell, gmmap.EmptyCell},
				{gmmap.CrossCell, gmmap.EmptyCell, gmmap.CircleCell, gmmap.EmptyCell, gmmap.EmptyCell},
				{gmmap.CrossCell, gmmap.EmptyCell, gmmap.CircleCell, gmmap.EmptyCell, gmmap.EmptyCell},
			},
			[2]int{3, 0},
			false,
		},
		{
			Game{GameMapInterface: &gmmap.InMemoryGameMap{},
				UserTurn:     false,
				GameFinished: false,
			},
			[][]gmmap.CellValue{
				{gmmap.CrossCell, gmmap.EmptyCell, gmmap.EmptyCell, gmmap.EmptyCell, gmmap.EmptyCell},
				{gmmap.CrossCell, gmmap.EmptyCell, gmmap.CircleCell, gmmap.EmptyCell, gmmap.EmptyCell},
				{gmmap.CrossCell, gmmap.EmptyCell, gmmap.CircleCell, gmmap.EmptyCell, gmmap.EmptyCell},
				{gmmap.CircleCell, gmmap.EmptyCell, gmmap.CircleCell, gmmap.EmptyCell, gmmap.EmptyCell},
				{gmmap.CrossCell, gmmap.EmptyCell, gmmap.CircleCell, gmmap.EmptyCell, gmmap.EmptyCell},
			},
			[2]int{3, 0},
			false,
		},
	}
	funcToTestsIsWinningCombinationExistForCell(t, testGameWins)
}

func TestGame_IsWinningCombinationExistForCellHorizontal(t *testing.T) {
	testGameWins := []userGames{
		{
			Game{GameMapInterface: &gmmap.InMemoryGameMap{},
				UserTurn:     false,
				GameFinished: false,
			},
			[][]gmmap.CellValue{
				{gmmap.CrossCell, gmmap.CrossCell, gmmap.CrossCell, gmmap.CrossCell, gmmap.CrossCell},
				{gmmap.EmptyCell, gmmap.EmptyCell, gmmap.CircleCell, gmmap.EmptyCell, gmmap.EmptyCell},
				{gmmap.EmptyCell, gmmap.EmptyCell, gmmap.CircleCell, gmmap.EmptyCell, gmmap.EmptyCell},
				{gmmap.EmptyCell, gmmap.EmptyCell, gmmap.CircleCell, gmmap.EmptyCell, gmmap.EmptyCell},
				{gmmap.EmptyCell, gmmap.EmptyCell, gmmap.CircleCell, gmmap.EmptyCell, gmmap.EmptyCell},
			},
			[2]int{0, 0},
			true,
		},
		{
			Game{GameMapInterface: &gmmap.InMemoryGameMap{},
				UserTurn:     false,
				GameFinished: false,
			},
			[][]gmmap.CellValue{
				{gmmap.EmptyCell, gmmap.EmptyCell, gmmap.EmptyCell, gmmap.EmptyCell, gmmap.EmptyCell},
				{gmmap.CrossCell, gmmap.EmptyCell, gmmap.CircleCell, gmmap.EmptyCell, gmmap.EmptyCell},
				{gmmap.CrossCell, gmmap.CrossCell, gmmap.CrossCell, gmmap.CrossCell, gmmap.CrossCell},
				{gmmap.CircleCell, gmmap.EmptyCell, gmmap.CircleCell, gmmap.EmptyCell, gmmap.EmptyCell},
				{gmmap.CircleCell, gmmap.EmptyCell, gmmap.CircleCell, gmmap.EmptyCell, gmmap.EmptyCell},
			},
			[2]int{2, 0},
			true,
		},
		{
			Game{GameMapInterface: &gmmap.InMemoryGameMap{},
				UserTurn:     false,
				GameFinished: false,
			},
			[][]gmmap.CellValue{
				{gmmap.EmptyCell, gmmap.EmptyCell, gmmap.EmptyCell, gmmap.EmptyCell, gmmap.EmptyCell},
				{gmmap.CrossCell, gmmap.EmptyCell, gmmap.CircleCell, gmmap.EmptyCell, gmmap.EmptyCell},
				{gmmap.EmptyCell, gmmap.EmptyCell, gmmap.EmptyCell, gmmap.EmptyCell, gmmap.EmptyCell},
				{gmmap.CircleCell, gmmap.EmptyCell, gmmap.CircleCell, gmmap.EmptyCell, gmmap.EmptyCell},
				{gmmap.CrossCell, gmmap.CrossCell, gmmap.CrossCell, gmmap.CrossCell, gmmap.CrossCell},
			},
			[2]int{4, 0},
			true,
		},
		{
			Game{GameMapInterface: &gmmap.InMemoryGameMap{},
				UserTurn:     false,
				GameFinished: false,
			},
			[][]gmmap.CellValue{
				{gmmap.EmptyCell, gmmap.EmptyCell, gmmap.EmptyCell, gmmap.EmptyCell, gmmap.EmptyCell},
				{gmmap.CrossCell, gmmap.EmptyCell, gmmap.CircleCell, gmmap.EmptyCell, gmmap.EmptyCell},
				{gmmap.EmptyCell, gmmap.EmptyCell, gmmap.EmptyCell, gmmap.EmptyCell, gmmap.EmptyCell},
				{gmmap.CircleCell, gmmap.EmptyCell, gmmap.CircleCell, gmmap.EmptyCell, gmmap.EmptyCell},
				{gmmap.CrossCell, gmmap.CrossCell, gmmap.CrossCell, gmmap.EmptyCell, gmmap.CrossCell},
			},
			[2]int{4, 0},
			false,
		},
	}
	funcToTestsIsWinningCombinationExistForCell(t, testGameWins)
}

func TestIsWinningCombinationExistForCellFalse(t *testing.T) {
	a := assert.New(t)
	testGameWins := []userGames{
	}
	for _, testGameWin := range testGameWins {
		a.False(testGameWin.userGame.IsWinningCombinationExistForCell(testGameWin.cellCoordinate[0], testGameWin.cellCoordinate[1]))
	}
}
