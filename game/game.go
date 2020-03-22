package game

import (
	"errors"
	"goTicTacToe/game/gmmap"
)

const emptyCellValue = 0

type Game struct {
	UserTurn     bool
	GameFinished bool
	gmmap.GameMapInterface
}

func New(mapSizeX int, mapSizeY int) (*Game, error) {
	userGame := &Game{UserTurn: true, GameFinished: false, GameMapInterface: &gmmap.InMemoryGameMap{}}
	newMapError := userGame.NewMap(mapSizeX, mapSizeY)
	if newMapError != nil {
		return nil, newMapError
	}

	return userGame, nil
}

func (game Game) IsGameEnded() bool {
	return !isNextTurnAvailable(game) || game.GameFinished
}

func (game *Game) MakeTurn(x int, y int) (bool, error) {
	gameMap := game.GetMap()
	if x > len(gameMap) || x < 0 || y > len(gameMap[0]) || y < 0 {
		return false, errors.New("cell coordinates do not exist")
	}

	if gameMap[x][y] != emptyCellValue {
		return false, errors.New("cell is not empty")
	}

	var writeValueToCellError error
	if game.UserTurn {
		writeValueToCellError = game.WriteValueToCell(1, x, y)
	} else {
		writeValueToCellError = game.WriteValueToCell(2, x, y)
	}

	if writeValueToCellError != nil {
		return false, writeValueToCellError
	}

	game.UserTurn = !game.UserTurn
	return true, nil
}

func (game *Game) IsWinningCombinationExistForCell(x int, y int) bool {
	gameMap := game.GetMap()

	cellValue := gameMap[x][y]
	gameMapXLen := len(gameMap)
	gameMapYLen := len(gameMap)

	if x >= 0 && x < gameMapXLen && y >= 0 && y < gameMapYLen && cellValue != emptyCellValue {
		//TODO написать более изысканную проверку
		verticalCellValueStrike := getVerticallyValueStrikeForCell(gameMap, x, y)

		horizontalCellValueStrike := getHorizontalValueStrikeForCell(gameMap, x, y)

		backSlashValueStrike := getBackSlashValueStrikeForCell(gameMap, x, y)

		slashValueStrike := getSlashValueStrikeForCell(gameMap, x, y)

		cellValueStrike := getMaxInt(verticalCellValueStrike, horizontalCellValueStrike, backSlashValueStrike, slashValueStrike)

		if getGameCellsForWinCount(*game) <= cellValueStrike {
			game.GameFinished = true
			return true
		}
	}
	return false
}

func getVerticallyValueStrikeForCell(gameMap [][]int, x int, y int) int {
	cellValue := gameMap[x][y]
	cellValueStrike := 0

	holdValue := false
	for i := 0; i < len(gameMap); i++ {
		if i == x {
			holdValue = true
		}
		if gameMap[i][y] == cellValue {
			cellValueStrike++
		} else {
			if !holdValue {
				//если мы еще не прошли через ячейку, то сбрасываем счетчик
				cellValueStrike = 0
			} else {
				//если уже прошли, то выходим
				break
			}
		}
	}
	return cellValueStrike
}

func getHorizontalValueStrikeForCell(gameMap [][]int, x int, y int) int {
	cellValue := gameMap[x][y]
	cellValueStrike := 0

	holdValue := false
	for i := 0; i < len(gameMap[0]); i++ {
		if i == y {
			holdValue = true
		}
		if gameMap[x][i] == cellValue {
			cellValueStrike++
		} else {
			if !holdValue {
				//если мы еще не прошли через ячейку, то сбрасываем счетчик
				cellValueStrike = 0
			} else {
				//если уже прошли, то выходим
				break
			}
		}
	}
	return cellValueStrike
}

/*
	\
*/
func getBackSlashValueStrikeForCell(gameMap [][]int, x int, y int) int {
	cellValue := gameMap[x][y]
	cellValueStrike := 0

	holdValue := false

	var startCell [2]int

	if x >= y {
		startCell[0] = x - y
		startCell[1] = 0
	} else {
		startCell[0] = 0
		startCell[1] = y - x
	}

	horizontalCellCoordinate := startCell[1]
	mapHorizontalLen := len(gameMap[0])
	for i := startCell[0]; i < len(gameMap); i++ {
		if i == x && horizontalCellCoordinate == y {
			holdValue = true
		}
		if gameMap[i][horizontalCellCoordinate] == cellValue {
			cellValueStrike++
		} else {
			if !holdValue {
				//если мы еще не прошли через ячейку, то сбрасываем счетчик
				cellValueStrike = 0
			} else {
				//если уже прошли, то выходим
				break
			}
		}

		horizontalCellCoordinate++

		if horizontalCellCoordinate >= mapHorizontalLen {
			break
		}
	}
	return cellValueStrike
}

/*
	/
*/
func getSlashValueStrikeForCell(gameMap [][]int, x int, y int) int {
	cellValue := gameMap[x][y]
	cellValueStrike := 0

	holdValue := false

	var startCell [2]int
	startCell[0] = 0
	startCell[1] = y + x
	mapHorizontalLen := len(gameMap[0])
	if startCell[1] >= mapHorizontalLen {
		startCell[0] = startCell[1] - mapHorizontalLen + 1
		startCell[1] = mapHorizontalLen - 1
	}

	horizontalCellCoordinate := startCell[1]

	for i := startCell[0]; i < len(gameMap); i++ {
		if i == x && horizontalCellCoordinate == y {
			holdValue = true
		}
		if gameMap[i][horizontalCellCoordinate] == cellValue {
			cellValueStrike++
		} else {
			if !holdValue {
				//если мы еще не прошли через ячейку, то сбрасываем счетчик
				cellValueStrike = 0
			} else {
				//если уже прошли, то выходим
				break
			}
		}

		horizontalCellCoordinate--

		if horizontalCellCoordinate < 0 {
			break
		}
	}
	return cellValueStrike
}

func getGameCellsForWinCount(game Game) int {
	gameMap := game.GetMap()
	verticalLen := len(gameMap)
	horizontalLen := len(gameMap[0])

	if verticalLen <= horizontalLen {
		return verticalLen
	} else {
		return horizontalLen
	}
}

func isNextTurnAvailable(game Game) bool {
	gameMap := game.GetMap()
	for _, gameMapRow := range gameMap {
		for _, gameMapCell := range gameMapRow {
			if gameMapCell == emptyCellValue {
				return true
			}
		}
	}
	return false
}

func getMaxInt(nums ...int) int {
	//TODO add nil check
	maxNum := nums[0]
	for _, num := range nums {
		if maxNum < num {
			maxNum = num
		}
	}
	return maxNum
}
